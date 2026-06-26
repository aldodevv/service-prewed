package main

import (
	"context"
	"log"
	netHttp "net/http"
	"os"
	"service-wedding/internal/delivery/http"
	"service-wedding/internal/repository/postgres"
	"service-wedding/internal/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// 1. Load env variables
	// Since we are running from Cwd service-wedding, it will load local .env
	_ = godotenv.Load()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set in environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 2. Initialize database connection pool
	pool, err := postgres.NewConnPool(dbURL)
	if err != nil {
		log.Fatalf("Failed to initialize database connection: %v", err)
	}
	defer pool.Close()

	// 3. Run auto-migrations
	runMigrations(pool)

	// 4. Seed admin user if none exists
	seedAdminUser(pool)

	// 5. Initialize clean arch layers
	userRepo := postgres.NewUserRepository(pool)
	themeRepo := postgres.NewThemeRepository(pool)
	contextRepo := postgres.NewContextRepository(pool)
	guestRepo := postgres.NewGuestRepository(pool)
	assetRepo := postgres.NewAssetRepository(pool)
	contactMsgRepo := postgres.NewContactRepository(pool)

	cldName := os.Getenv("CLOUDINARY_NAME")
	cldKey := os.Getenv("CLOUDINARY_KEY")
	cldSecret := os.Getenv("CLOUDINARY_SECRET")

	authUsecase := usecase.NewAuthUsecase(userRepo)
	themeUsecase := usecase.NewThemeUsecase(themeRepo)
	contextUsecase := usecase.NewContextUsecase(contextRepo, themeRepo)
	guestUsecase := usecase.NewGuestUsecase(guestRepo, contextRepo)
	assetUsecase := usecase.NewAssetUsecase(assetRepo, cldName, cldKey, cldSecret)
	contactUsecase := usecase.NewContactUsecase(contactMsgRepo)

	authH := http.NewAuthHandler(authUsecase)
	themeH := http.NewThemeHandler(themeUsecase)
	contextH := http.NewContextHandler(contextUsecase)
	guestH := http.NewGuestHandler(guestUsecase)
	assetH := http.NewAssetHandler(assetUsecase)
	publicH := http.NewPublicHandler(contextUsecase, guestUsecase, themeUsecase, assetUsecase)
	contactH := http.NewContactHandler(contactUsecase)

	// 6. Set up Gin engine
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	http.NewRouter(r, authH, themeH, contextH, guestH, assetH, publicH, contactH)

	// 7. Start server
	log.Printf("Server is starting and listening on port :%s", port)
	server := &netHttp.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != netHttp.ErrServerClosed {
		log.Fatalf("Could not listen on port %s: %v", port, err)
	}
}

func runMigrations(pool *pgxpool.Pool) {
	// Read migrations/001_init.sql
	migrationFile := "migrations/001_init.sql"
	sqlBytes, err := os.ReadFile(migrationFile)
	if err != nil {
		log.Printf("Warning: Could not read migration file: %v. Database schema might not be configured.", err)
		return
	}

	_, err = pool.Exec(context.Background(), string(sqlBytes))
	if err != nil {
		log.Fatalf("Failed to execute database migration: %v", err)
	}
	log.Println("Database schema auto-migrations executed successfully!")
}

func seedAdminUser(pool *pgxpool.Pool) {
	ctx := context.Background()

	// Check if the default admin user exists by email
	var count int
	err := pool.QueryRow(ctx, "SELECT COUNT(*) FROM users WHERE email = $1", "admin@wedify.com").Scan(&count)
	if err != nil {
		log.Printf("Warning: failed to query admin user count for seeding: %v", err)
		return
	}

	if count > 0 {
		return // Admin user already exists, no seeding needed
	}

	// Seed default admin
	name := "Administrator"
	email := "admin@wedify.com"
	pass := "adminpassword123"
	role := "admin"

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Warning: failed to hash password for seeding: %v", err)
		return
	}

	_, err = pool.Exec(ctx,
		"INSERT INTO users (name, email, password_hash, role, created_at, updated_at) VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)",
		name, email, string(hashedPass), role,
	)
	if err != nil {
		log.Printf("Warning: failed to seed admin user: %v", err)
		return
	}

	log.Println("==========================================================")
	log.Println("DEFAULT ADMIN ACCOUNT CREATED SUCCESSFULLY:")
	log.Println("Email:    admin@wedify.com")
	log.Println("Password: adminpassword123")
	log.Println("==========================================================")
}
