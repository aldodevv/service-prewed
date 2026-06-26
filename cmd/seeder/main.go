package main

import (
	"context"
	"log"
	"os"
	"service-wedding/internal/repository/postgres"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env from current directory or parent
	_ = godotenv.Load()
	_ = godotenv.Load(".env")
	_ = godotenv.Load("service-wedding/.env")
	_ = godotenv.Load("../service-wedding/.env")

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}
	dbURL = strings.ReplaceAll(dbURL, "-pooler", "")
	dbURL = strings.ReplaceAll(dbURL, "channel_binding=require", "")
	dbURL = strings.TrimRight(dbURL, "&?")

	log.Printf("Connecting to database (direct, no channel binding): %s\n", dbURL[:strings.Index(dbURL, "@")+1]+"...hidden...")

	pool, err := postgres.NewConnPool(dbURL)
	if err != nil {
		log.Fatalf("Failed to initialize database connection: %v", err)
	}
	defer pool.Close()

	log.Println("Starting database seeding for themes...")
	seedThemes(pool)
	log.Println("Seeded all themes successfully! Starting contexts seeding...")
	seedDummyContexts(pool)
	log.Println("Database seeding completed successfully!")
}

func seedRoyalGold(pool *pgxpool.Pool) {
	ctx := context.Background()

	// Always clear and re-create 'royal-gold' to update it with full templates, layout, and animations
	_, _ = pool.Exec(ctx, "DELETE FROM themes WHERE slug = $1", "royal-gold")

	name := "Royal Gold Premium"
	slug := "royal-gold"
	description := "Tema pernikahan premium dengan nuansa warna emas mewah, putih gading, dan layout komplit profesional."
	thumbnail := "https://images.unsplash.com/photo-1519741497674-611481863552?auto=format&fit=crop&w=400&q=80"

	themeDataJSON := `{
  "global": {
    "backgroundColor": "#faf6f0",
    "containerWidth": "100%",
    "fontFamily": "Outfit, sans-serif",
    "primaryColor": "#b89c5a"
  },
  "splash": {
    "title": "PERNIKAHAN DARI",
    "heading": "Romeo & Juliet",
    "fontFamily": "'Great Vibes', cursive",
    "bgImageUrl": "https://images.unsplash.com/photo-1519741497674-611481863552?auto=format&fit=crop&w=1200&q=80",
    "bgOverlayColor": "#09090b",
    "bgOverlayOpacity": "0.75",
    "cardBgColor": "rgba(24, 24, 27, 0.65)",
    "cardBorderColor": "#b89c5a",
    "cardBorderRadius": "24px",
    "cardTextColor": "#ffffff",
    "buttonText": "Buka Undangan",
    "buttonBgColor": "#b89c5a",
    "buttonTextColor": "#ffffff",
    "logoColor": "#b89c5a"
  },
  "sections": [
    {
      "id": "sec-cover",
      "name": "Cover & Opening",
      "paddingTop": "6rem",
      "paddingBottom": "6rem",
      "paddingLeft": "2rem",
      "paddingRight": "2rem",
      "marginTop": "0px",
      "marginBottom": "0px",
      "marginLeft": "0px",
      "marginRight": "0px",
      "backgroundColor": "#ffffff",
      "bgImageUrl": "https://images.unsplash.com/photo-1519741497674-611481863552?auto=format&fit=crop&w=1200&q=80",
      "bgVideoUrl": "https://res.cloudinary.com/dyuol7xfx/video/upload/v1782347274/wedding/scf8luzxedgpxsz49qg5.mp4",
      "bgImageSize": "cover",
      "bgImagePosition": "center",
      "bgImageRepeat": "no-repeat",
      "bgImageAttachment": "scroll",
      "bgOverlayColor": "#000000",
      "bgOverlayOpacity": "0.55",
      "borderRadius": "0px",
      "borderWidth": "0px",
      "borderColor": "transparent",
      "borderStyle": "none",
      "boxShadow": "none",
      "animation": "fade-in",
      "transition": "fade-up-parallax",
      "display": "block",
      "flexDirection": "column",
      "alignItems": "center",
      "justifyContent": "center",
      "gap": "1.5rem",
      "textColor": "#ffffff",
      "customStyles": "",
      "widgets": [
        {
          "id": "w-cover-icon",
          "type": "icon",
          "content": "ring",
          "style": {
            "margin": "0 auto 1.5rem",
            "color": "#b89c5a",
            "textAlign": "center",
            "display": "block"
          },
          "meta": { "size": "36", "strokeWidth": "2", "color": "#b89c5a" }
        },
        {
          "id": "w-cover-1",
          "type": "heading",
          "content": "THE WEDDING OF",
          "style": {
            "color": "#ffffff",
            "fontSize": "12px",
            "fontWeight": "700",
            "letterSpacing": "0.2em",
            "textAlign": "center",
            "textTransform": "uppercase"
          },
          "meta": { "level": "h4" }
        },
        {
          "id": "w-cover-2",
          "type": "heading",
          "content": "Romeo & Juliet",
          "style": {
            "color": "#b89c5a",
            "fontFamily": "Great Vibes, cursive",
            "fontSize": "56px",
            "fontWeight": "bold",
            "textAlign": "center",
            "margin": "1rem 0"
          },
          "meta": { "level": "h1" }
        },
        {
          "id": "w-cover-3",
          "type": "text",
          "content": "Sabtu, 12 Desember 2026",
          "style": {
            "color": "#e2e8f0",
            "fontSize": "15px",
            "textAlign": "center",
            "fontWeight": "500",
            "letterSpacing": "0.05em"
          }
        },
        {
          "id": "w-cover-4",
          "type": "countdown",
          "content": "",
          "style": {
            "margin": "2rem auto"
          },
          "meta": { "targetDate": "2026-12-12T08:00:00" }
        },
        {
          "id": "w-cover-5",
          "type": "audio",
          "content": "https://res.cloudinary.com/dyuol7xfx/video/upload/v1782361665/wedding/pzhplg7kmln1p1huxz6e.mp3",
          "style": {
            "margin": "1rem auto"
          }
        },
        {
          "id": "w-cover-flower-left",
          "type": "floating_decor",
          "content": "https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347235/wedding/tcjan0ozyqzxrchrm8yw.jpg",
          "style": {
            "position": "absolute",
            "top": "8%",
            "left": "-10px",
            "width": "80px",
            "height": "120px",
            "opacity": "0.85",
            "zIndex": "10"
          },
          "meta": { "animation": "sway", "speed": "6s", "delay": "0s" }
        },
        {
          "id": "w-cover-flower-right",
          "type": "floating_decor",
          "content": "https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347235/wedding/tcjan0ozyqzxrchrm8yw.jpg",
          "style": {
            "position": "absolute",
            "top": "15%",
            "right": "-10px",
            "width": "80px",
            "height": "120px",
            "opacity": "0.85",
            "zIndex": "10",
            "transform": "scaleX(-1)"
          },
          "meta": { "animation": "sway", "speed": "8s", "delay": "1s" }
        }
      ]
    },
    {
      "id": "sec-mempelai",
      "name": "Mempelai (Profil)",
      "paddingTop": "5rem",
      "paddingBottom": "5rem",
      "paddingLeft": "2rem",
      "paddingRight": "2rem",
      "marginTop": "0px",
      "marginBottom": "0px",
      "marginLeft": "0px",
      "marginRight": "0px",
      "backgroundColor": "#ffffff",
      "borderRadius": "0px",
      "borderWidth": "0px",
      "borderColor": "transparent",
      "borderStyle": "none",
      "boxShadow": "none",
      "animation": "fade-in",
      "transition": "fade-up",
      "display": "block",
      "flexDirection": "column",
      "alignItems": "center",
      "justifyContent": "center",
      "gap": "1.5rem",
      "textColor": "#333333",
      "widgets": [
        {
          "id": "w-memp-1",
          "type": "quote",
          "content": "Dan di antara tanda-tanda (kebesaran)-Nya ialah Dia menciptakan pasangan-pasangan untukmu dari jenismu sendiri, agar kamu cenderung dan merasa tenteram kepadanya, dan Dia menjadikan di antaramu rasa kasih dan sayang.",
          "style": {
            "color": "#555555",
            "fontSize": "13px",
            "fontStyle": "italic",
            "lineHeight": "1.7",
            "textAlign": "center"
          },
          "meta": { "author": "Ar-Rum: 21" }
        },
        {
          "id": "w-memp-2",
          "type": "divider",
          "content": "",
          "style": {
            "margin": "2.5rem auto"
          }
        },
        {
          "id": "w-memp-3",
          "type": "image",
          "content": "https://images.unsplash.com/photo-1507679799987-c73779587ccf?auto=format&fit=crop&w=400&q=80",
          "style": {
            "width": "200px",
            "height": "280px",
            "margin": "0 auto 1.5rem"
          },
          "meta": { "frame": "arch-classic" }
        },
        {
          "id": "w-memp-4",
          "type": "heading",
          "content": "Romeo Montague, S.Kom",
          "style": {
            "color": "#1a1a1a",
            "fontSize": "22px",
            "fontWeight": "700",
            "textAlign": "center"
          },
          "meta": { "level": "h3" }
        },
        {
          "id": "w-memp-5",
          "type": "text",
          "content": "Putra tercinta dari Bapak Montague & Ibu Montague\\nBandung, Jawa Barat",
          "style": {
            "color": "#666666",
            "fontSize": "13px",
            "textAlign": "center",
            "lineHeight": "1.6",
            "marginBottom": "1rem"
          }
        },
        {
          "id": "w-memp-groom-social",
          "type": "social_links",
          "content": "",
          "style": {
            "gap": "12px",
            "justifyContent": "center",
            "marginTop": "0.5rem",
            "marginBottom": "2.5rem"
          },
          "meta": {
            "instagram": "https://instagram.com/romeo_montague",
            "size": "20",
            "color": "#b89c5a"
          }
        },
        {
          "id": "w-memp-icon",
          "type": "icon",
          "content": "heart",
          "style": {
            "color": "#b89c5a",
            "textAlign": "center",
            "margin": "1rem 0"
          },
          "meta": { "size": "32", "strokeWidth": "2", "color": "#b89c5a" }
        },
        {
          "id": "w-memp-7",
          "type": "image",
          "content": "https://images.unsplash.com/photo-1519741497674-611481863552?auto=format&fit=crop&w=400&q=80",
          "style": {
            "width": "200px",
            "height": "280px",
            "margin": "1.5rem auto 1.5rem"
          },
          "meta": { "frame": "arch-classic" }
        },
        {
          "id": "w-memp-8",
          "type": "heading",
          "content": "Juliet Capulet, B.A",
          "style": {
            "color": "#1a1a1a",
            "fontSize": "22px",
            "fontWeight": "700",
            "textAlign": "center"
          },
          "meta": { "level": "h3" }
        },
        {
          "id": "w-memp-9",
          "type": "text",
          "content": "Putri tercinta dari Bapak Capulet & Ibu Capulet\\nJakarta Selatan, DKI Jakarta",
          "style": {
            "color": "#666666",
            "fontSize": "13px",
            "textAlign": "center",
            "lineHeight": "1.6",
            "marginBottom": "1rem"
          }
        },
        {
          "id": "w-memp-bride-social",
          "type": "social_links",
          "content": "",
          "style": {
            "gap": "12px",
            "justifyContent": "center",
            "marginTop": "0.5rem"
          },
          "meta": {
            "instagram": "https://instagram.com/juliet_capulet",
            "size": "20",
            "color": "#b89c5a"
          }
        },
        {
          "id": "w-memp-flower-left",
          "type": "floating_decor",
          "content": "https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347246/wedding/q2qgf7ktfml1ff9zupkb.jpg",
          "style": {
            "position": "absolute",
            "bottom": "15%",
            "left": "10px",
            "width": "60px",
            "height": "60px",
            "opacity": "0.35",
            "zIndex": "2"
          },
          "meta": { "animation": "drift", "speed": "8s", "delay": "0s" }
        },
        {
          "id": "w-memp-flower-right",
          "type": "floating_decor",
          "content": "https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347246/wedding/q2qgf7ktfml1ff9zupkb.jpg",
          "style": {
            "position": "absolute",
            "top": "35%",
            "right": "10px",
            "width": "50px",
            "height": "50px",
            "opacity": "0.3",
            "zIndex": "2",
            "transform": "scaleX(-1)"
          },
          "meta": { "animation": "drift", "speed": "10s", "delay": "2s" }
        }
      ]
    },
    {
      "id": "sec-event",
      "name": "Waktu & Tempat Acara",
      "paddingTop": "5rem",
      "paddingBottom": "5rem",
      "paddingLeft": "2rem",
      "paddingRight": "2rem",
      "marginTop": "0px",
      "marginBottom": "0px",
      "marginLeft": "0px",
      "marginRight": "0px",
      "backgroundColor": "#faf6f0",
      "borderRadius": "0px",
      "borderWidth": "0px",
      "borderColor": "transparent",
      "borderStyle": "none",
      "boxShadow": "none",
      "transition": "zoom-in",
      "widgets": [
        {
          "id": "w-evt-1",
          "type": "heading",
          "content": "Rangkaian Acara",
          "style": {
            "color": "#b89c5a",
            "fontSize": "26px",
            "fontWeight": "700",
            "textAlign": "center"
          },
          "meta": { "level": "h2" }
        },
        {
          "id": "w-evt-2",
          "type": "table",
          "content": "[[\"Acara\", \"Waktu & Lokasi\"], [\"Akad Nikah\", \"08:00 - 10:00 WIB\\nMasjid Al-Barkah, Bandung\"], [\"Resepsi Pernikahan\", \"11:00 WIB - Selesai\\nHall Grand Ballroom, Bandung\"]]",
          "style": {
            "width": "100%",
            "marginTop": "1.5rem",
            "marginBottom": "1.5rem",
            "fontSize": "13.5px",
            "color": "#333333"
          },
          "meta": { "hasHeader": true, "border": "1px solid #e5d8c0", "padding": "12px" }
        },
        {
          "id": "w-evt-3",
          "type": "list",
          "content": "Catatan Protokol Acara:\\n* Harap datang 15 menit sebelum acara dimulai.\\n* Wajib memindai kode QR undangan di meja registrasi.\\n* Dresscode: Pakaian Formal / Batik Modern.",
          "style": {
            "fontSize": "13px",
            "color": "#555555",
            "textAlign": "left",
            "marginTop": "1rem",
            "marginBottom": "2rem"
          },
          "meta": { "listType": "ul" }
        },
        {
          "id": "w-evt-icon",
          "type": "icon",
          "content": "map-pin",
          "style": {
            "color": "#b89c5a",
            "textAlign": "center",
            "margin": "2rem 0 0"
          },
          "meta": { "size": "32", "strokeWidth": "2", "color": "#b89c5a" }
        },
        {
          "id": "w-evt-4",
          "type": "heading",
          "content": "Peta Lokasi Acara",
          "style": {
            "color": "#1a1a1a",
            "fontSize": "18px",
            "fontWeight": "700",
            "textAlign": "center",
            "marginTop": "1rem"
          },
          "meta": { "level": "h3" }
        },
        {
          "id": "w-evt-5",
          "type": "map",
          "content": "https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3966.0528659100085!2d106.8911033!3d-6.256754!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2e69f30be6f19985%3A0xb21b10645c36395b!2sSasana%20Kriya!5e0!3m2!1sid!2sid!4v1655000000000!5m2!1sid!2sid",
          "style": {
            "margin": "1.5rem auto"
          }
        },
        {
          "id": "w-evt-6",
          "type": "button",
          "content": "Buka Google Maps Lokasi",
          "style": {
            "margin": "1.5rem auto",
            "color": "#ffffff",
            "fontWeight": "bold",
            "fontSize": "14px"
          },
          "meta": { "url": "https://maps.google.com/?q=Sasana+Kriya+TMII" }
        }
      ]
    },
    {
      "id": "sec-story",
      "name": "Kisah Cinta & Perjalanan",
      "paddingTop": "5rem",
      "paddingBottom": "5rem",
      "paddingLeft": "2rem",
      "paddingRight": "2rem",
      "marginTop": "0px",
      "marginBottom": "0px",
      "marginLeft": "0px",
      "marginRight": "0px",
      "backgroundColor": "#ffffff",
      "transition": "slide-left",
      "widgets": [
        {
          "id": "w-sty-1",
          "type": "heading",
          "content": "Kisah Cinta Kami",
          "style": {
            "color": "#b89c5a",
            "fontSize": "26px",
            "fontWeight": "700",
            "textAlign": "center"
          },
          "meta": { "level": "h2" }
        },
        {
          "id": "w-sty-2",
          "type": "accordion",
          "content": "Pertemuan pertama kami terjadi di awal tahun 2020 secara tidak sengaja di sebuah perpustakaan. Dari sekadar berbagi buku referensi, kami menyadari banyak ketertarikan yang sama. Waktu berlalu dan pertemanan itu berkembang menjadi rasa saling menghargai dan cinta.",
          "style": {
            "marginTop": "1rem",
            "marginBottom": "1rem",
            "color": "#555555"
          },
          "meta": { "summary": "Awal Pertemuan Pertama (2020)" }
        },
        {
          "id": "w-sty-3",
          "type": "accordion",
          "content": "Di tahun 2024, di hadapan kedua orang tua dan keluarga terdekat, kami mengikat janji pertunangan resmi. Sebuah komitmen tulus untuk melangkah bersama mempersiapkan hari pernikahan ini.",
          "style": {
            "marginTop": "1rem",
            "marginBottom": "1rem",
            "color": "#555555"
          },
          "meta": { "summary": "Pertunangan Resmi (2024)" }
        }
      ]
    },
    {
      "id": "sec-gallery",
      "name": "Galeri Foto & Video",
      "paddingTop": "5rem",
      "paddingBottom": "5rem",
      "paddingLeft": "2rem",
      "paddingRight": "2rem",
      "marginTop": "0px",
      "marginBottom": "0px",
      "marginLeft": "0px",
      "marginRight": "0px",
      "backgroundColor": "#faf6f0",
      "transition": "zoom-in-parallax",
      "widgets": [
        {
          "id": "w-gal-1",
          "type": "heading",
          "content": "Galeri Kenangan",
          "style": {
            "color": "#b89c5a",
            "fontSize": "26px",
            "fontWeight": "700",
            "textAlign": "center"
          },
          "meta": { "level": "h2" }
        },
        {
          "id": "w-gal-2",
          "type": "gallery",
          "content": "",
          "style": {
            "margin": "2rem auto"
          },
          "meta": {
            "cols": 2,
            "images": [
              "https://images.unsplash.com/photo-1519741497674-611481863552?auto=format&fit=crop&w=400&q=80",
              "https://images.unsplash.com/photo-1511795409834-ef04bbd61622?auto=format&fit=crop&w=400&q=80",
              "https://images.unsplash.com/photo-1519225495810-7512c696505a?auto=format&fit=crop&w=400&q=80",
              "https://images.unsplash.com/photo-1465495976277-4387d4b0b4c6?auto=format&fit=crop&w=400&q=80"
            ]
          }
        },
        {
          "id": "w-gal-3",
          "type": "video",
          "content": "https://res.cloudinary.com/dyuol7xfx/video/upload/v1782347274/wedding/scf8luzxedgpxsz49qg5.mp4",
          "style": {
            "margin": "1.5rem auto"
          }
        }
      ]
    },
    {
      "id": "sec-rsvp",
      "name": "Buku Tamu & RSVP",
      "paddingTop": "5rem",
      "paddingBottom": "5rem",
      "paddingLeft": "2rem",
      "paddingRight": "2rem",
      "marginTop": "0px",
      "marginBottom": "0px",
      "marginLeft": "0px",
      "marginRight": "0px",
      "backgroundColor": "#ffffff",
      "transition": "fade-up",
      "widgets": [
        {
          "id": "w-rsv-1",
          "type": "heading",
          "content": "Kehadiran & Doa Restu",
          "style": {
            "color": "#b89c5a",
            "fontSize": "26px",
            "fontWeight": "700",
            "textAlign": "center"
          },
          "meta": { "level": "h2" }
        },
        {
          "id": "w-rsv-2",
          "type": "text",
          "content": "Konfirmasikan kehadiran Anda dan berikan pesan hangat untuk kami berdua.",
          "style": {
            "color": "#666666",
            "fontSize": "13.5px",
            "textAlign": "center",
            "marginBottom": "2rem"
          }
        },
        {
          "id": "w-rsv-3",
          "type": "input",
          "content": "",
          "style": {
            "padding": "10px 14px",
            "borderWidth": "1px",
            "borderColor": "#e5d8c0",
            "borderStyle": "solid",
            "borderRadius": "8px",
            "width": "100%",
            "height": "42px",
            "fontSize": "14px",
            "backgroundColor": "#ffffff"
          },
          "meta": {
            "label": "Nama Lengkap Anda",
            "placeholder": "Ketik nama lengkap...",
            "name": "nama_tamu"
          }
        },
        {
          "id": "w-rsv-4",
          "type": "select",
          "content": "1 Tamu\n2 Tamu\n3 Tamu\n4 Tamu",
          "style": {
            "padding": "10px 14px",
            "borderWidth": "1px",
            "borderColor": "#e5d8c0",
            "borderStyle": "solid",
            "borderRadius": "8px",
            "width": "100%",
            "height": "42px",
            "fontSize": "14px",
            "backgroundColor": "#ffffff"
          },
          "meta": {
            "label": "Jumlah Tamu Yang Hadir",
            "name": "jumlah_tamu"
          }
        },
        {
          "id": "w-rsv-5",
          "type": "radio",
          "content": "Ya, Saya akan hadir\nMaaf, Saya berhalangan hadir",
          "style": {
            "marginTop": "0.5rem",
            "marginBottom": "0.5rem"
          },
          "meta": {
            "label": "Konfirmasi Kehadiran",
            "name": "kehadiran"
          }
        },
        {
          "id": "w-rsv-6",
          "type": "textarea",
          "content": "",
          "style": {
            "padding": "10px 14px",
            "borderWidth": "1px",
            "borderColor": "#e5d8c0",
            "borderStyle": "solid",
            "borderRadius": "8px",
            "width": "100%",
            "fontSize": "14px",
            "backgroundColor": "#ffffff"
          },
          "meta": {
            "label": "Doa Restu & Ucapan",
            "placeholder": "Ketik ucapan selamat...",
            "rows": 4,
            "name": "ucapan"
          }
        },
        {
          "id": "w-rsv-7",
          "type": "submit_btn",
          "content": "Kirim RSVP & Doa Restu",
          "style": {
            "width": "100%",
            "height": "46px",
            "borderRadius": "8px",
            "backgroundColor": "#b89c5a",
            "color": "#ffffff",
            "fontWeight": "bold",
            "fontSize": "14px",
            "border": "none",
            "marginTop": "1rem",
            "cursor": "pointer"
          }
        }
      ]
    },
    {
      "id": "sec-gift",
      "name": "Kado & Hadiah Digital",
      "paddingTop": "5rem",
      "paddingBottom": "5rem",
      "paddingLeft": "2rem",
      "paddingRight": "2rem",
      "marginTop": "0px",
      "marginBottom": "0px",
      "marginLeft": "0px",
      "marginRight": "0px",
      "backgroundColor": "#faf6f0",
      "transition": "zoom-in",
      "widgets": [
        {
          "id": "w-gft-1",
          "type": "heading",
          "content": "Kado Pernikahan",
          "style": {
            "color": "#b89c5a",
            "fontSize": "26px",
            "fontWeight": "700",
            "textAlign": "center"
          },
          "meta": { "level": "h2" }
        },
        {
          "id": "w-gft-2",
          "type": "text",
          "content": "Doa restu Anda adalah karunia terindah bagi kami. Namun bagi Anda yang ingin memberikan tanda kasih secara digital, Anda dapat mengirimkannya melalui rekening berikut:",
          "style": {
            "color": "#666666",
            "fontSize": "13.5px",
            "textAlign": "center",
            "lineHeight": "1.6",
            "marginBottom": "2rem"
          }
        },
        {
          "id": "w-gft-3",
          "type": "container",
          "content": "BCA (Bank Central Asia)\nNo. Rek: 809-1234-567\na.n. Romeo Montague",
          "style": {
            "padding": "20px",
            "borderRadius": "12px",
            "border": "1px solid #e5d8c0",
            "backgroundColor": "#ffffff",
            "textAlign": "center",
            "fontWeight": "600",
            "color": "#333333",
            "margin": "1rem auto",
            "maxWidth": "360px",
            "boxShadow": "0 4px 15px rgba(0,0,0,0.02)"
          }
        },
        {
          "id": "w-gft-4",
          "type": "container",
          "content": "Bank Mandiri\nNo. Rek: 131-00-123456-7\na.n. Juliet Capulet",
          "style": {
            "padding": "20px",
            "borderRadius": "12px",
            "border": "1px solid #e5d8c0",
            "backgroundColor": "#ffffff",
            "textAlign": "center",
            "fontWeight": "600",
            "color": "#333333",
            "margin": "1rem auto",
            "maxWidth": "360px",
            "boxShadow": "0 4px 15px rgba(0,0,0,0.02)"
          }
        },
        {
          "id": "w-gft-5",
          "type": "text_link",
          "content": "Kunjungi registries kado kami di Tokopedia",
          "style": {
            "color": "#b89c5a",
            "fontSize": "13.5px",
            "textDecoration": "underline",
            "display": "block",
            "textAlign": "center",
            "marginTop": "2rem"
          },
          "meta": { "url": "https://tokopedia.com" }
        }
      ]
    },
    {
      "id": "sec-footer",
      "name": "Ucapan Terima Kasih (Footer)",
      "paddingTop": "4rem",
      "paddingBottom": "4rem",
      "paddingLeft": "2rem",
      "paddingRight": "2rem",
      "marginTop": "0px",
      "marginBottom": "0px",
      "marginLeft": "0px",
      "marginRight": "0px",
      "backgroundColor": "#111111",
      "textColor": "#ffffff",
      "transition": "fade-in",
      "widgets": [
        {
          "id": "w-ftr-icon",
          "type": "icon",
          "content": "heart",
          "style": {
            "color": "#b89c5a",
            "textAlign": "center",
            "margin": "0 auto 1.5rem",
            "display": "block"
          },
          "meta": { "size": "36", "strokeWidth": "2", "color": "#b89c5a" }
        },
        {
          "id": "w-ftr-1",
          "type": "heading",
          "content": "Terima Kasih",
          "style": {
            "color": "#b89c5a",
            "fontFamily": "Great Vibes, cursive",
            "fontSize": "38px",
            "textAlign": "center"
          },
          "meta": { "level": "h2" }
        },
        {
          "id": "w-ftr-2",
          "type": "text",
          "content": "Merupakan suatu kehormatan dan kebahagiaan bagi kami apabila Bapak/Ibu/Saudara/i berkenan hadir dan memberikan doa restu kepada kami.\\n\\nKami yang berbahagia,\\nRomeo & Juliet",
          "style": {
            "color": "#999999",
            "fontSize": "13px",
            "textAlign": "center",
            "lineHeight": "1.7",
            "marginTop": "1.5rem"
          }
        },
        {
          "id": "w-ftr-socials",
          "type": "social_links",
          "content": "",
          "style": {
            "gap": "16px",
            "justifyContent": "center",
            "marginTop": "1.5rem"
          },
          "meta": {
            "instagram": "https://instagram.com/romeo_montague",
            "facebook": "https://facebook.com",
            "youtube": "https://youtube.com",
            "size": "24",
            "color": "#b89c5a"
          }
        }
      ]
    }
  ]
}`

	renderHTML := `
<style>
  @import url('https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;500;600;700;800&family=Great+Vibes&family=Playfair+Display:ital,wght@0,400..900;1,400..900&display=swap');
  body { margin: 0; padding: 0; font-family: 'Outfit', sans-serif; background-color: #faf6f0; }
  * { box-sizing: border-box; }
  .list-disc { list-style-type: disc; }
  .list-decimal { list-style-type: decimal; }

  /* Animation System */
  @keyframes floatSway {
    0% { transform: translateY(0px) rotate(0deg); }
    50% { transform: translateY(-15px) rotate(6deg); }
    100% { transform: translateY(0px) rotate(0deg); }
  }
  @keyframes floatSwaySlow {
    0% { transform: translateY(0px) rotate(0deg); }
    50% { transform: translateY(-25px) rotate(-10deg); }
    100% { transform: translateY(0px) rotate(0deg); }
  }
  @keyframes spinSlow {
    0% { transform: rotate(0deg); }
    100% { transform: rotate(360deg); }
  }
  @keyframes pulseSlow {
    0%, 100% { transform: scale(1); }
    50% { transform: scale(1.12); }
  }
  @keyframes driftLeftRight {
    0%, 100% { transform: translateX(0px) rotate(0deg); }
    50% { transform: translateX(20px) rotate(8deg); }
    100% { transform: translateX(0px) rotate(0deg); }
  }
  .animate-float-sway { animation: floatSway var(--float-speed, 6s) ease-in-out infinite var(--float-delay, 0s); }
  .animate-float-sway-slow { animation: floatSwaySlow var(--float-speed, 10s) ease-in-out infinite var(--float-delay, 0s); }
  .animate-spin-slow { animation: spinSlow var(--float-speed, 15s) linear infinite var(--float-delay, 0s); }
  .animate-pulse-slow { animation: pulseSlow var(--float-speed, 5s) ease-in-out infinite var(--float-delay, 0s); }
  .animate-drift-lr { animation: driftLeftRight var(--float-speed, 8s) ease-in-out infinite var(--float-delay, 0s); }

  @keyframes fadeUp {
    from { opacity: 0; transform: translateY(30px); }
    to { opacity: 1; transform: translateY(0); }
  }
  @keyframes fadeIn {
    from { opacity: 0; }
    to { opacity: 1; }
  }
  @keyframes slideLeft {
    from { opacity: 0; transform: translateX(-30px); }
    to { opacity: 1; transform: translateX(0); }
  }
  @keyframes slideRight {
    from { opacity: 0; transform: translateX(30px); }
    to { opacity: 1; transform: translateX(0); }
  }
  @keyframes zoomIn {
    from { opacity: 0; transform: scale(0.95); }
    to { opacity: 1; transform: scale(1); }
  }

  .animate-widget {
    opacity: 0;
    transition: opacity 0.8s cubic-bezier(0.16, 1, 0.3, 1), transform 0.8s cubic-bezier(0.16, 1, 0.3, 1);
  }

  .is-in-view[data-transition="fade-up"] .animate-widget,
  .is-in-view[data-transition="fade-up-parallax"] .animate-widget {
    animation: fadeUp 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
  }
  .is-in-view[data-transition="fade-in"] .animate-widget {
    animation: fadeIn 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
  }
  .is-in-view[data-transition="slide-left"] .animate-widget {
    animation: slideLeft 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
  }
  .is-in-view[data-transition="slide-right"] .animate-widget {
    animation: slideRight 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
  }
  .is-in-view[data-transition="zoom-in"] .animate-widget,
  .is-in-view[data-transition="zoom-in-parallax"] .animate-widget {
    animation: zoomIn 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
  }

  /* Parallax Background Layout */
  .parallax-section {
    position: relative;
    overflow: hidden;
  }
  .parallax-bg {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 125%;
    z-index: 0;
    pointer-events: none;
    background-size: cover;
    background-position: center;
    background-repeat: no-repeat;
    will-change: transform;
  }

  @media (prefers-reduced-motion: reduce) {
    .animate-widget {
      opacity: 1 !important;
      transform: none !important;
      animation: none !important;
      transition: none !important;
    }
    .parallax-bg {
      transform: none !important;
      height: 100% !important;
    }
  }
</style>
<div style="font-family: 'Outfit', sans-serif; background-color: #faf6f0; min-height: 100vh; width: 100%;">
  <div style="max-width: 100%; margin: 0 auto; min-height: 100vh; box-shadow: 0 10px 50px rgba(0,0,0,0.04); background: #ffffff;">

    <!-- Cover Section -->
    <section id="sec-cover" class="parallax-section" data-transition="fade-up-parallax" style="position: relative; padding: 6rem 2rem; text-align: center; min-height: 100vh; display: flex; flex-direction: column; align-items: center; justify-content: center; color: #ffffff; overflow: hidden; background-color: #ffffff;">
      <div class="parallax-bg" style="background-image: url('https://images.unsplash.com/photo-1519741497674-611481863552?auto=format&fit=crop&w=1200&q=80'); background-size: cover; background-position: center; transform: translateY(0px) scale(1.15);"></div>
      
      <!-- Video Loop Background -->
      <video autoplay loop muted playsinline style="position: absolute; top:0; left:0; width:100%; height:100%; object-fit:cover; z-index: 0; opacity: 0.25; pointer-events: none;">
        <source src="https://res.cloudinary.com/dyuol7xfx/video/upload/v1782347274/wedding/scf8luzxedgpxsz49qg5.mp4" type="video/mp4">
      </video>
      <div style="position: absolute; top:0; left:0; width:100%; height:100%; background-color: #000000; opacity: 0.55; z-index: 1;"></div>
      
      <!-- Floating Gold Foliage -->
      <div class="animate-float-sway" style="position: absolute; left: -10px; top: 8%; width: 80px; height: 120px; z-index: 10; pointer-events: none; opacity: 0.85; --float-speed: 6s; --float-delay: 0s;">
        <img src="https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347235/wedding/tcjan0ozyqzxrchrm8yw.jpg" style="width: 100%; height: 100%; object-fit: contain;" alt="Decor" />
      </div>
      <div class="animate-float-sway" style="position: absolute; right: -10px; top: 15%; width: 80px; height: 120px; z-index: 10; pointer-events: none; opacity: 0.85; transform: scaleX(-1); --float-speed: 8s; --float-delay: 1s;">
        <img src="https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347235/wedding/tcjan0ozyqzxrchrm8yw.jpg" style="width: 100%; height: 100%; object-fit: contain; transform: scaleX(-1);" alt="Decor" />
      </div>

      <div style="position: relative; z-index: 2; width: 100%;">
        <!-- Arch Ring Icon -->
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%; display: flex; justify-content: center; margin: 0 auto 1.5rem;">
          <svg viewBox="0 0 24 24" width="36" height="36" stroke="#b89c5a" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="14" r="7"/><path d="M12 2a4 4 0 0 1 4 4v1H8V6a4 4 0 0 1 4-4Z"/></svg>
        </div>

        <div class="animate-widget" style="animation-delay: 0.12s; transition-delay: 0.12s; width: 100%;">
          <h4 style="color: #ffffff; font-size: 12px; font-weight: 700; letter-spacing: 0.2em; text-align: center; text-transform: uppercase; margin-top: 0;">THE WEDDING OF</h4>
        </div>
        <div class="animate-widget" style="animation-delay: 0.24s; transition-delay: 0.24s; width: 100%;">
          <h1 style="color: #b89c5a; font-family: 'Great Vibes', cursive; font-size: 56px; font-weight: bold; text-align: center; margin: 1rem 0;">Romeo & Juliet</h1>
        </div>
        <div class="animate-widget" style="animation-delay: 0.36s; transition-delay: 0.36s; width: 100%;">
          <p style="color: #e2e8f0; font-size: 15px; text-align: center; font-weight: 500; letter-spacing: 0.05em;">Sabtu, 12 Desember 2026</p>
        </div>
        
        <div class="animate-widget" style="animation-delay: 0.48s; transition-delay: 0.48s; width: 100%;">
          <div style="display: flex; justify-content: center; gap: 1rem; margin: 2rem 0;">
            <div style="background: rgba(255,255,255,0.15); backdrop-filter: blur(10px); padding: 1.2rem; border-radius: 16px; min-width: 75px; border: 1px solid rgba(255,255,255,0.15);">
              <div style="font-size: 1.8rem; font-weight: 800; color: #b89c5a;">120</div>
              <div style="font-size: 0.65rem; text-transform: uppercase; color: #e2e8f0; font-weight: 600; margin-top: 2px;">Hari</div>
            </div>
            <div style="background: rgba(255,255,255,0.15); backdrop-filter: blur(10px); padding: 1.2rem; border-radius: 16px; min-width: 75px; border: 1px solid rgba(255,255,255,0.15);">
              <div style="font-size: 1.8rem; font-weight: 800; color: #b89c5a;">08</div>
              <div style="font-size: 0.65rem; text-transform: uppercase; color: #e2e8f0; font-weight: 600; margin-top: 2px;">Jam</div>
            </div>
            <div style="background: rgba(255,255,255,0.15); backdrop-filter: blur(10px); padding: 1.2rem; border-radius: 16px; min-width: 75px; border: 1px solid rgba(255,255,255,0.15);">
              <div style="font-size: 1.8rem; font-weight: 800; color: #b89c5a;">45</div>
              <div style="font-size: 0.65rem; text-transform: uppercase; color: #e2e8f0; font-weight: 600; margin-top: 2px;">Menit</div>
            </div>
          </div>
        </div>

        <div class="animate-widget" style="animation-delay: 0.6s; transition-delay: 0.6s; width: 100%;">
          <div style="background: rgba(255,255,255,0.15); border: 1px solid rgba(255,255,255,0.2); padding: 0.75rem 1.5rem; border-radius: 99px; display: inline-flex; align-items: center; gap: 1rem; margin: 1rem auto; backdrop-filter: blur(5px);">
            <div style="width: 28px; height: 28px; border-radius: 50%; background: #b89c5a; display: flex; align-items: center; justify-content: center; color: white; font-weight: bold; cursor: pointer;">▶</div>
            <span style="font-size: 13px; color: #ffffff; font-weight: 500;">Background Music</span>
          </div>
        </div>
      </div>
    </section>

    <!-- Mempelai Section -->
    <section id="sec-mempelai" data-transition="fade-up" style="position: relative; padding: 5rem 2rem; background-color: #ffffff; text-align: center; color: #333333; overflow: hidden;">
      <!-- Floating flowers in background -->
      <div class="animate-drift-lr" style="position: absolute; left: 10px; bottom: 15%; width: 60px; height: 60px; z-index: 2; opacity: 0.35; --float-speed: 8s; --float-delay: 0s;">
        <img src="https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347246/wedding/q2qgf7ktfml1ff9zupkb.jpg" style="width: 100%; height: 100%; object-fit: contain;" alt="Decor" />
      </div>
      <div class="animate-drift-lr" style="position: absolute; right: 10px; top: 35%; width: 50px; height: 50px; z-index: 2; opacity: 0.3; transform: scaleX(-1); --float-speed: 10s; --float-delay: 2s;">
        <img src="https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347246/wedding/q2qgf7ktfml1ff9zupkb.jpg" style="width: 100%; height: 100%; object-fit: contain; transform: scaleX(-1);" alt="Decor" />
      </div>

      <div style="position: relative; z-index: 2; width: 100%;">
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%;">
          <div style="padding: 1.5rem; border-left: 3px solid #b89c5a; margin: 2rem auto; max-width: 500px; text-align: left; background: #faf8f5; border-radius: 0 16px 16px 0;">
            <p style="margin: 0; font-style: italic; color: #555555; line-height: 1.7; font-size: 13px;">"Dan di antara tanda-tanda (kebesaran)-Nya ialah Dia menciptakan pasangan-pasangan untukmu dari jenismu sendiri, agar kamu cenderung dan merasa tenteram kepadanya, dan Dia menjadikan di antaramu rasa kasih dan sayang."</p>
            <p style="margin: 0.5rem 0 0 0; font-size: 12px; font-weight: bold; text-align: right; color: #b89c5a;">- Ar-Rum: 21</p>
          </div>
        </div>
        
        <div class="animate-widget" style="animation-delay: 0.12s; transition-delay: 0.12s; width: 100%;">
          <hr style="border: 0; border-top: 1px solid #eeeeee; width: 60%; margin: 2.5rem auto;" />
        </div>
        
        <!-- Groom framed portrait (arch-classic) -->
        <div class="animate-widget" style="animation-delay: 0.24s; transition-delay: 0.24s; width: 100%;">
          <div style="border-radius: 120px 120px 12px 12px; border: 2px solid #b89c5a; padding: 4px; display: inline-block; margin: 1rem auto; overflow: hidden; width: 200px; height: 280px; background: #ffffff; box-shadow: 0 10px 25px rgba(184,156,90,0.12);">
            <img src="https://images.unsplash.com/photo-1507679799987-c73779587ccf?auto=format&fit=crop&w=400&q=80" style="display: block; width: 100%; height: 100%; object-fit: cover; border-radius: 116px 116px 8px 8px;" alt="Groom" />
          </div>
        </div>

        <div class="animate-widget" style="animation-delay: 0.36s; transition-delay: 0.36s; width: 100%;">
          <h3 style="color: #1a1a1a; font-size: 22px; font-weight: 700; text-align: center; margin-top: 0; font-family: 'Playfair Display', serif;">Romeo Montague, S.Kom</h3>
        </div>
        <div class="animate-widget" style="animation-delay: 0.48s; transition-delay: 0.48s; width: 100%;">
          <p style="color: #666666; font-size: 13px; text-align: center; line-height: 1.6; margin-bottom: 0.5rem;">Putra tercinta dari Bapak Montague & Ibu Montague<br>Bandung, Jawa Barat</p>
        </div>
        
        <!-- Groom Social Links -->
        <div class="animate-widget" style="animation-delay: 0.6s; transition-delay: 0.6s; width: 100%;">
          <div style="display: flex; gap: 12px; justify-content: center; margin-bottom: 2.5rem;">
            <a href="https://instagram.com/romeo_montague" target="_blank" rel="noopener noreferrer" style="display: inline-flex; align-items: center; justify-content: center; width: 36px; height: 36px; border-radius: 50%; border: 1.5px solid #b89c5a; color: #b89c5a; text-decoration: none; transition: all 0.3s ease;">
              <svg viewBox="0 0 24 24" width="20" height="20" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><rect x="2" y="2" width="20" height="20" rx="5" ry="5"></rect><path d="M16 11.37A4 4 0 1 1 12.63 8 4 4 0 0 1 16 11.37z"></path><line x1="17.5" y1="6.5" x2="17.51" y2="6.5"></line></svg>
            </a>
          </div>
        </div>
        
        <!-- Heart Separator Icon -->
        <div class="animate-widget" style="animation-delay: 0.72s; transition-delay: 0.72s; width: 100%; display: flex; justify-content: center; margin: 1rem 0;">
          <svg viewBox="0 0 24 24" width="32" height="32" stroke="#b89c5a" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/></svg>
        </div>
        
        <!-- Bride framed portrait (arch-classic) -->
        <div class="animate-widget" style="animation-delay: 0.84s; transition-delay: 0.84s; width: 100%;">
          <div style="border-radius: 120px 120px 12px 12px; border: 2px solid #b89c5a; padding: 4px; display: inline-block; margin: 1.5rem auto 1.5rem; overflow: hidden; width: 200px; height: 280px; background: #ffffff; box-shadow: 0 10px 25px rgba(184,156,90,0.12);">
            <img src="https://images.unsplash.com/photo-1519741497674-611481863552?auto=format&fit=crop&w=400&q=80" style="display: block; width: 100%; height: 100%; object-fit: cover; border-radius: 116px 116px 8px 8px;" alt="Bride" />
          </div>
        </div>

        <div class="animate-widget" style="animation-delay: 0.96s; transition-delay: 0.96s; width: 100%;">
          <h3 style="color: #1a1a1a; font-size: 22px; font-weight: 700; text-align: center; margin-top: 0; font-family: 'Playfair Display', serif;">Juliet Capulet, B.A</h3>
        </div>
        <div class="animate-widget" style="animation-delay: 1.08s; transition-delay: 1.08s; width: 100%;">
          <p style="color: #666666; font-size: 13px; text-align: center; line-height: 1.6; margin-bottom: 0.5rem;">Putri tercinta dari Bapak Capulet & Ibu Capulet<br>Jakarta Selatan, DKI Jakarta</p>
        </div>
        
        <!-- Bride Social Links -->
        <div class="animate-widget" style="animation-delay: 1.2s; transition-delay: 1.2s; width: 100%;">
          <div style="display: flex; gap: 12px; justify-content: center;">
            <a href="https://instagram.com/juliet_capulet" target="_blank" rel="noopener noreferrer" style="display: inline-flex; align-items: center; justify-content: center; width: 36px; height: 36px; border-radius: 50%; border: 1.5px solid #b89c5a; color: #b89c5a; text-decoration: none; transition: all 0.3s ease;">
              <svg viewBox="0 0 24 24" width="20" height="20" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><rect x="2" y="2" width="20" height="20" rx="5" ry="5"></rect><path d="M16 11.37A4 4 0 1 1 12.63 8 4 4 0 0 1 16 11.37z"></path><line x1="17.5" y1="6.5" x2="17.51" y2="6.5"></line></svg>
            </a>
          </div>
        </div>
      </div>
    </section>

    <!-- Rangkaian Acara -->
    <section id="sec-event" data-transition="zoom-in" style="padding: 5rem 2rem; background-color: #faf6f0; text-align: center; color: #333333; position: relative;">
      <div style="position: relative; z-index: 2; width: 100%;">
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%;">
          <h2 style="color: #b89c5a; font-size: 26px; font-weight: 700; text-align: center; margin-top: 0; font-family: 'Playfair Display', serif; letter-spacing: 0.05em;">Rangkaian Acara</h2>
        </div>
        <div class="animate-widget" style="animation-delay: 0.12s; transition-delay: 0.12s; width: 100%;">
          <div style="width: 60px; height: 2px; background: #b89c5a; margin: 1rem auto 2.5rem;"></div>
        </div>

        <div style="display: flex; flex-direction: column; gap: 2rem; max-width: 450px; margin: 0 auto;">
          <!-- Akad Nikah Card -->
          <div class="animate-widget" style="animation-delay: 0.24s; transition-delay: 0.24s; width: 100%;">
            <div style="background: #ffffff; border: 2px double #b89c5a; border-radius: 12px; padding: 2rem; box-shadow: 0 8px 30px rgba(184, 156, 90, 0.05); text-align: center;">
              <h3 style="color: #b89c5a; font-family: 'Great Vibes', cursive; font-size: 32px; margin-top: 0; margin-bottom: 0.5rem; font-weight: normal;">Akad Nikah</h3>
              <p style="font-size: 15px; font-weight: 600; margin: 0.5rem 0; color: #1a1a1a;">Sabtu, 12 Desember 2026</p>
              <p style="font-size: 14px; margin: 0.25rem 0; color: #b89c5a; font-weight: 700;">08:00 - 10:00 WIB</p>
              <div style="width: 30px; height: 1px; background: #e5d8c0; margin: 1rem auto;"></div>
              <p style="font-size: 13.5px; line-height: 1.6; color: #555555; margin: 0;"><b>Masjid Al-Barkah</b><br>Jl. Asia Afrika No. 12, Bandung</p>
            </div>
          </div>

          <!-- Resepsi Card -->
          <div class="animate-widget" style="animation-delay: 0.36s; transition-delay: 0.36s; width: 100%;">
            <div style="background: #ffffff; border: 2px double #b89c5a; border-radius: 12px; padding: 2rem; box-shadow: 0 8px 30px rgba(184, 156, 90, 0.05); text-align: center;">
              <h3 style="color: #b89c5a; font-family: 'Great Vibes', cursive; font-size: 32px; margin-top: 0; margin-bottom: 0.5rem; font-weight: normal;">Resepsi Pernikahan</h3>
              <p style="font-size: 15px; font-weight: 600; margin: 0.5rem 0; color: #1a1a1a;">Sabtu, 12 Desember 2026</p>
              <p style="font-size: 14px; margin: 0.25rem 0; color: #b89c5a; font-weight: 700;">11:00 WIB - Selesai</p>
              <div style="width: 30px; height: 1px; background: #e5d8c0; margin: 1rem auto;"></div>
              <p style="font-size: 13.5px; line-height: 1.6; color: #555555; margin: 0;"><b>Hall Grand Ballroom</b><br>Hotel Luxury Regency, Bandung</p>
            </div>
          </div>
        </div>

        <div style="margin-top: 3rem;">
          <!-- Map Pin Icon -->
          <div class="animate-widget" style="animation-delay: 0.48s; transition-delay: 0.48s; width: 100%; display: flex; justify-content: center; margin: 2rem auto 0;">
            <svg viewBox="0 0 24 24" width="32" height="32" stroke="#b89c5a" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><path d="M20 10c0 6-8 12-8 12s-8-6-8-12a8 8 0 0 1 16 0Z"/><circle cx="12" cy="10" r="3"/></svg>
          </div>
          <div class="animate-widget" style="animation-delay: 0.6s; transition-delay: 0.6s; width: 100%;">
            <h3 style="color: #1a1a1a; font-size: 18px; font-weight: 700; text-align: center; font-family: 'Playfair Display', serif; margin-top: 1rem;">Peta Lokasi Acara</h3>
          </div>
          <div class="animate-widget" style="animation-delay: 0.72s; transition-delay: 0.72s; width: 100%;">
            <div style="margin: 1.5rem auto; max-width: 480px; border: 2px double #b89c5a; border-radius: 16px; padding: 6px; background: #ffffff;">
              <iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3966.0528659100085!2d106.8911033!3d-6.256754!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2e69f30be6f19985%3A0xb21b10645c36395b!2sSasana%20Kriya!5e0!3m2!1sid!2sid!4v1655000000000!5m2!1sid!2sid" width="100%" height="280" style="border:0; border-radius:10px;" allowfullscreen="" loading="lazy"></iframe>
            </div>
          </div>
          <div class="animate-widget" style="animation-delay: 0.84s; transition-delay: 0.84s; width: 100%;">
            <a href="https://maps.google.com/?q=Sasana+Kriya+TMII" target="_blank" style="display: inline-flex; align-items: center; justify-content: center; padding: 0.75rem 2.25rem; border-radius: 4px; background-color: #b89c5a; color: #ffffff; text-decoration: none; font-weight: 600; font-size: 13.5px; letter-spacing: 0.05em; border: 1px solid #b89c5a; box-shadow: 0 4px 15px rgba(184,156,90,0.15); transition: all 0.3s;">
              BUKA GOOGLE MAPS
            </a>
          </div>
        </div>
      </div>
    </section>

    <!-- Kisah Cinta -->
    <section id="sec-story" data-transition="slide-left" style="padding: 5rem 2rem; background-color: #ffffff; text-align: center; position: relative;">
      <div style="position: relative; z-index: 2; width: 100%;">
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%;">
          <h2 style="color: #b89c5a; font-size: 26px; font-weight: 700; text-align: center; margin-top: 0; font-family: 'Playfair Display', serif; letter-spacing: 0.05em;">Kisah Cinta Kami</h2>
        </div>
        <div class="animate-widget" style="animation-delay: 0.12s; transition-delay: 0.12s; width: 100%;">
          <div style="width: 60px; height: 2px; background: #b89c5a; margin: 1rem auto 3rem;"></div>
        </div>

        <div style="position: relative; max-width: 500px; margin: 0 auto; padding: 1rem 0 1rem 2rem; text-align: left; border-left: 2px solid #e5d8c0;">
          <!-- Timeline Item 1 -->
          <div style="position: relative; margin-bottom: 2.5rem;">
            <div style="position: absolute; left: -2.4rem; top: 4px; width: 12px; height: 12px; border-radius: 50%; background: #ffffff; border: 3px solid #b89c5a; box-shadow: 0 0 0 4px #faf6f0;"></div>
            <div class="animate-widget" style="animation-delay: 0.24s; transition-delay: 0.24s; width: 100%;">
              <span style="font-size: 12px; font-weight: 800; color: #b89c5a; letter-spacing: 0.1em;">AWAL PERTEMUAN (2020)</span>
              <h3 style="font-size: 18px; font-family: 'Playfair Display', serif; color: #1a1a1a; margin: 0.25rem 0 0.5rem 0; font-weight: 700;">Menemukan Satu Sama Lain</h3>
              <p style="font-size: 13.5px; color: #555555; line-height: 1.7; margin: 0;">Pertemuan pertama kami terjadi di awal tahun 2020 secara tidak sengaja di sebuah perpustakaan. Dari sekadar berbagi buku referensi, kami menyadari banyak ketertarikan yang sama. Waktu berlalu dan pertemanan itu berkembang menjadi rasa saling menghargai dan cinta.</p>
            </div>
          </div>

          <!-- Timeline Item 2 -->
          <div style="position: relative; margin-bottom: 0;">
            <div style="position: absolute; left: -2.4rem; top: 4px; width: 12px; height: 12px; border-radius: 50%; background: #ffffff; border: 3px solid #b89c5a; box-shadow: 0 0 0 4px #faf6f0;"></div>
            <div class="animate-widget" style="animation-delay: 0.36s; transition-delay: 0.36s; width: 100%;">
              <span style="font-size: 12px; font-weight: 800; color: #b89c5a; letter-spacing: 0.1em;">PERTUNANGAN RESMI (2024)</span>
              <h3 style="font-size: 18px; font-family: 'Playfair Display', serif; color: #1a1a1a; margin: 0.25rem 0 0.5rem 0; font-weight: 700;">Mengikat Janji Bersama</h3>
              <p style="font-size: 13.5px; color: #555555; line-height: 1.7; margin: 0;">Di tahun 2024, di hadapan kedua orang tua dan keluarga terdekat, kami mengikat janji pertunangan resmi. Sebuah komitmen tulus untuk melangkah bersama mempersiapkan hari pernikahan ini.</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Galeri -->
    <section id="sec-gallery" data-transition="zoom-in-parallax" style="padding: 5rem 2rem; background-color: #faf6f0; text-align: center;">
      <div style="position: relative; z-index: 2; width: 100%;">
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%;">
          <h2 style="color: #b89c5a; font-size: 26px; font-weight: 700; text-align: center; margin-top: 0; font-family: 'Playfair Display', serif; letter-spacing: 0.05em;">Galeri Kenangan</h2>
        </div>
        <div class="animate-widget" style="animation-delay: 0.12s; transition-delay: 0.12s; width: 100%;">
          <div style="width: 60px; height: 2px; background: #b89c5a; margin: 1rem auto 2.5rem;"></div>
        </div>

        <div class="animate-widget" style="animation-delay: 0.24s; transition-delay: 0.24s; width: 100%;">
          <div style="display: grid; grid-template-columns: repeat(2, 1fr); gap: 1.25rem; margin: 2rem 0; max-width: 500px; margin-left: auto; margin-right: auto;">
            <div style="border: 1px solid #b89c5a; padding: 6px; background: #ffffff; box-shadow: 0 4px 15px rgba(0,0,0,0.03); border-radius: 4px;">
              <img src="https://images.unsplash.com/photo-1519741497674-611481863552?auto=format&fit=crop&w=400&q=80" style="width: 100%; height: 160px; object-fit: cover; border-radius: 2px; display: block;" alt="Gallery 1" />
            </div>
            <div style="border: 1px solid #b89c5a; padding: 6px; background: #ffffff; box-shadow: 0 4px 15px rgba(0,0,0,0.03); border-radius: 4px;">
              <img src="https://images.unsplash.com/photo-1511795409834-ef04bbd61622?auto=format&fit=crop&w=400&q=80" style="width: 100%; height: 160px; object-fit: cover; border-radius: 2px; display: block;" alt="Gallery 2" />
            </div>
            <div style="border: 1px solid #b89c5a; padding: 6px; background: #ffffff; box-shadow: 0 4px 15px rgba(0,0,0,0.03); border-radius: 4px;">
              <img src="https://images.unsplash.com/photo-1519225495810-7512c696505a?auto=format&fit=crop&w=400&q=80" style="width: 100%; height: 160px; object-fit: cover; border-radius: 2px; display: block;" alt="Gallery 3" />
            </div>
            <div style="border: 1px solid #b89c5a; padding: 6px; background: #ffffff; box-shadow: 0 4px 15px rgba(0,0,0,0.03); border-radius: 4px;">
              <img src="https://images.unsplash.com/photo-1465495976277-4387d4b0b4c6?auto=format&fit=crop&w=400&q=80" style="width: 100%; height: 160px; object-fit: cover; border-radius: 2px; display: block;" alt="Gallery 4" />
            </div>
          </div>
        </div>

        <div class="animate-widget" style="animation-delay: 0.36s; transition-delay: 0.36s; width: 100%;">
          <div style="max-width: 500px; margin: 1.5rem auto; border: 1px solid #b89c5a; padding: 6px; background: #ffffff; box-shadow: 0 4px 15px rgba(0,0,0,0.03); border-radius: 4px;">
            <video controls style="width:100%; border-radius:2px; display:block;">
              <source src="https://res.cloudinary.com/dyuol7xfx/video/upload/v1782347274/wedding/scf8luzxedgpxsz49qg5.mp4" type="video/mp4">
            </video>
          </div>
        </div>
      </div>
    </section>

    <!-- RSVP Section -->
    <section id="sec-rsvp" data-transition="fade-up" style="padding: 5rem 2rem; background-color: #ffffff; text-align: center;">
      <div style="position: relative; z-index: 2; width: 100%;">
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%;">
          <h2 style="color: #b89c5a; font-size: 26px; font-weight: 700; text-align: center; margin-top: 0; font-family: 'Playfair Display', serif; letter-spacing: 0.05em;">Konfirmasi Kehadiran</h2>
        </div>
        <div class="animate-widget" style="animation-delay: 0.12s; transition-delay: 0.12s; width: 100%;">
          <div style="width: 60px; height: 2px; background: #b89c5a; margin: 1rem auto 1.5rem;"></div>
        </div>
        <div class="animate-widget" style="animation-delay: 0.24s; transition-delay: 0.24s; width: 100%;">
          <p style="color: #666666; font-size: 13.5px; text-align: center; margin-bottom: 2.5rem; max-width: 420px; margin-left: auto; margin-right: auto; line-height: 1.6;">Konfirmasikan kehadiran Anda dan berikan pesan hangat untuk kami berdua.</p>
        </div>

        <div class="animate-widget" style="animation-delay: 0.36s; transition-delay: 0.36s; width: 100%;">
          <div style="background: #ffffff; padding: 3rem 2rem; border-radius: 40px 40px 12px 12px; max-width: 460px; margin: 0 auto; text-align: left; box-shadow: 0 10px 40px rgba(184, 156, 90, 0.08); border: 1px solid #e5d8c0;">
            <form onsubmit="event.preventDefault(); alert('Terima kasih atas konfirmasi Anda!')">
              <div style="text-align: left; width: 100%; margin-bottom: 2rem;">
                <label style="display: block; font-size: 11px; font-weight: 700; letter-spacing: 0.1em; text-transform: uppercase; margin-bottom: 8px; color: #b89c5a;">NAMA LENGKAP</label>
                <input type="text" placeholder="Tuliskan nama Anda..." required style="border: none; border-bottom: 1px solid #b89c5a; border-radius: 0; outline: none; background: transparent; padding: 10px 0; width: 100%; font-size: 14px; color: #1a1a1a;" />
              </div>

              <div style="text-align: left; width: 100%; margin-bottom: 2rem;">
                <label style="display: block; font-size: 11px; font-weight: 700; letter-spacing: 0.1em; text-transform: uppercase; margin-bottom: 8px; color: #b89c5a;">JUMLAH TAMU</label>
                <select style="border: none; border-bottom: 1px solid #b89c5a; border-radius: 0; outline: none; background: transparent; padding: 10px 0; width: 100%; font-size: 14px; color: #1a1a1a;">
                  <option>1 Orang</option>
                  <option>2 Orang</option>
                  <option>3 Orang</option>
                  <option>4 Orang</option>
                </select>
              </div>

              <div style="text-align: left; width: 100%; margin-bottom: 2rem;">
                <label style="display: block; font-size: 11px; font-weight: 700; letter-spacing: 0.1em; text-transform: uppercase; margin-bottom: 12px; color: #b89c5a;">KONFIRMASI</label>
                <div style="display: flex; align-items: center; margin-bottom: 8px;">
                  <input type="radio" name="attendance" value="hadir" id="rg-att-1" checked style="accent-color: #b89c5a;" />
                  <label for="rg-att-1" style="font-size: 13px; color: #555555; margin-left: 8px; cursor: pointer;">Ya, Saya akan hadir</label>
                </div>
                <div style="display: flex; align-items: center; margin-bottom: 8px;">
                  <input type="radio" name="attendance" value="tidak_hadir" id="rg-att-2" style="accent-color: #b89c5a;" />
                  <label for="rg-att-2" style="font-size: 13px; color: #555555; margin-left: 8px; cursor: pointer;">Maaf, Saya berhalangan hadir</label>
                </div>
              </div>

              <div style="text-align: left; width: 100%; margin-bottom: 2rem;">
                <label style="display: block; font-size: 11px; font-weight: 700; letter-spacing: 0.1em; text-transform: uppercase; margin-bottom: 8px; color: #b89c5a;">DOA RESTU & UCAPAN</label>
                <textarea placeholder="Tuliskan ucapan selamat Anda..." rows="3" required style="border: none; border-bottom: 1px solid #b89c5a; border-radius: 0; outline: none; background: transparent; padding: 10px 0; width: 100%; font-size: 14px; color: #1a1a1a; resize: none;"></textarea>
              </div>

              <button type="submit" style="width: 100%; height: 48px; border-radius: 4px; background-color: #b89c5a; color: #ffffff; font-weight: bold; font-size: 13px; letter-spacing: 0.1em; border: none; cursor: pointer; text-transform: uppercase; transition: all 0.3s; box-shadow: 0 4px 15px rgba(184, 156, 90, 0.2);">Kirim Konfirmasi</button>
            </form>
          </div>
        </div>
      </div>
    </section>

    <!-- Gift Section -->
    <section id="sec-gift" data-transition="zoom-in" style="padding: 5rem 2rem; background-color: #faf6f0; text-align: center;">
      <div style="position: relative; z-index: 2; width: 100%;">
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%;">
          <h2 style="color: #b89c5a; font-size: 26px; font-weight: 700; text-align: center; margin-top: 0; font-family: 'Playfair Display', serif; letter-spacing: 0.05em;">Kado Pernikahan</h2>
        </div>
        <div class="animate-widget" style="animation-delay: 0.12s; transition-delay: 0.12s; width: 100%;">
          <div style="width: 60px; height: 2px; background: #b89c5a; margin: 1rem auto 1.5rem;"></div>
        </div>
        <div class="animate-widget" style="animation-delay: 0.24s; transition-delay: 0.24s; width: 100%;">
          <p style="color: #666666; font-size: 13.5px; text-align: center; line-height: 1.6; margin-bottom: 2.5rem; max-width: 440px; margin-left: auto; margin-right: auto;">Doa restu Anda adalah karunia terindah bagi kami. Namun bagi Anda yang ingin memberikan tanda kasih secara digital, Anda dapat mengirimkannya melalui rekening berikut:</p>
        </div>

        <div style="display: flex; flex-direction: column; gap: 1.5rem; max-width: 380px; margin: 0 auto;">
          <!-- BCA -->
          <div class="animate-widget" style="animation-delay: 0.36s; transition-delay: 0.36s; width: 100%;">
            <div style="padding: 24px; border-radius: 8px; border: 2px double #b89c5a; background-color: #ffffff; text-align: center; color: #333333; box-shadow: 0 4px 15px rgba(0,0,0,0.02); position: relative;">
              <p style="margin: 0; font-size: 12px; font-weight: bold; color: #b89c5a; letter-spacing: 0.05em;">BANK CENTRAL ASIA (BCA)</p>
              <p style="margin: 0.5rem 0; font-size: 20px; font-weight: 700; color: #1a1a1a; letter-spacing: 0.05em;">809-1234-567</p>
              <p style="margin: 0 0 1rem 0; font-size: 13px; color: #666666;">a.n. Romeo Montague</p>
              <button onclick="navigator.clipboard.writeText('809-1234-567'); const btn = this; const orig = btn.innerText; btn.innerText = 'Tersalin!'; setTimeout(() => btn.innerText = orig, 2000);" style="padding: 6px 16px; border-radius: 4px; border: 1px solid #b89c5a; background: transparent; color: #b89c5a; font-size: 12px; font-weight: 700; cursor: pointer; transition: all 0.2s;">Salin Rekening</button>
            </div>
          </div>

          <!-- Mandiri -->
          <div class="animate-widget" style="animation-delay: 0.48s; transition-delay: 0.48s; width: 100%;">
            <div style="padding: 24px; border-radius: 8px; border: 2px double #b89c5a; background-color: #ffffff; text-align: center; color: #333333; box-shadow: 0 4px 15px rgba(0,0,0,0.02); position: relative;">
              <p style="margin: 0; font-size: 12px; font-weight: bold; color: #b89c5a; letter-spacing: 0.05em;">BANK MANDIRI</p>
              <p style="margin: 0.5rem 0; font-size: 20px; font-weight: 700; color: #1a1a1a; letter-spacing: 0.05em;">131-00-123456-7</p>
              <p style="margin: 0 0 1rem 0; font-size: 13px; color: #666666;">a.n. Juliet Capulet</p>
              <button onclick="navigator.clipboard.writeText('131-00-123456-7'); const btn = this; const orig = btn.innerText; btn.innerText = 'Tersalin!'; setTimeout(() => btn.innerText = orig, 2000);" style="padding: 6px 16px; border-radius: 4px; border: 1px solid #b89c5a; background: transparent; color: #b89c5a; font-size: 12px; font-weight: 700; cursor: pointer; transition: all 0.2s;">Salin Rekening</button>
            </div>
          </div>
        </div>

        <div class="animate-widget" style="animation-delay: 0.6s; transition-delay: 0.6s; width: 100%;">
          <a href="https://tokopedia.com" target="_blank" style="color: #b89c5a; font-size: 13.5px; font-weight: 600; text-decoration: underline; display: block; text-align: center; margin-top: 2.5rem;">Kunjungi registries kado kami di Tokopedia</a>
        </div>
      </div>
    </section>

    <!-- Footer Section -->
    <section id="sec-footer" style="padding: 5rem 2rem; background-color: #1a150e; color: #ffffff; text-align: center; border-top: 3px solid #b89c5a; position: relative;">
      <div style="position: relative; z-index: 2; width: 100%;">
        <!-- Heart Icon -->
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%; display: flex; justify-content: center; margin: 0 auto 1.5rem;">
          <svg viewBox="0 0 24 24" width="36" height="36" stroke="#b89c5a" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/></svg>
        </div>
        
        <div class="animate-widget" style="animation-delay: 0.12s; transition-delay: 0.12s; width: 100%;">
          <h2 style="color: #b89c5a; font-family: 'Great Vibes', cursive; font-size: 38px; text-align: center; margin: 0; font-weight: normal;">Terima Kasih</h2>
        </div>
        <div class="animate-widget" style="animation-delay: 0.24s; transition-delay: 0.24s; width: 100%;">
          <div style="width: 40px; height: 1px; background: #b89c5a; margin: 1rem auto 1.5rem;"></div>
        </div>
        <div class="animate-widget" style="animation-delay: 0.36s; transition-delay: 0.36s; width: 100%;">
          <p style="color: #cbc3b7; font-size: 13.5px; text-align: center; line-height: 1.8; max-width: 440px; margin: 0 auto;">Merupakan suatu kehormatan dan kebahagiaan bagi kami apabila Bapak/Ibu/Saudara/i berkenan hadir dan memberikan doa restu kepada kami.<br><br>Kami yang berbahagia,<br><span style="color:#b89c5a; font-weight:bold;">Romeo & Juliet</span></p>
        </div>

        <!-- Social Media Link Row -->
        <div class="animate-widget" style="animation-delay: 0.48s; transition-delay: 0.48s; width: 100%;">
          <div style="display: flex; gap: 16px; justify-content: center; margin-top: 1.5rem;">
            <a href="https://instagram.com/romeo_montague" target="_blank" rel="noopener noreferrer" style="display: inline-flex; align-items: center; justify-content: center; width: 40px; height: 40px; border-radius: 50%; border: 1.5px solid #b89c5a; color: #b89c5a; text-decoration: none; transition: all 0.3s ease;">
              <svg viewBox="0 0 24 24" width="24" height="24" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><rect x="2" y="2" width="20" height="20" rx="5" ry="5"></rect><path d="M16 11.37A4 4 0 1 1 12.63 8 4 4 0 0 1 16 11.37z"></path><line x1="17.5" y1="6.5" x2="17.51" y2="6.5"></line></svg>
            </a>
            <a href="https://facebook.com" target="_blank" rel="noopener noreferrer" style="display: inline-flex; align-items: center; justify-content: center; width: 40px; height: 40px; border-radius: 50%; border: 1.5px solid #b89c5a; color: #b89c5a; text-decoration: none; transition: all 0.3s ease;">
              <svg viewBox="0 0 24 24" width="24" height="24" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><path d="M18 2h-3a5 5 0 0 0-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 0 1 1-1h3z"></path></svg>
            </a>
            <a href="https://youtube.com" target="_blank" rel="noopener noreferrer" style="display: inline-flex; align-items: center; justify-content: center; width: 40px; height: 40px; border-radius: 50%; border: 1.5px solid #b89c5a; color: #b89c5a; text-decoration: none; transition: all 0.3s ease;">
              <svg viewBox="0 0 24 24" width="24" height="24" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><path d="M22.54 6.42a2.78 2.78 0 0 0-1.95-1.96C18.88 4 12 4 12 4s-6.88 0-8.59.46A2.78 2.78 0 0 0 1.46 6.42 29 29 0 0 0 1 12a29 29 0 0 0 .46 5.58 2.78 2.78 0 0 0 1.95 1.96C5.12 20 12 20 12 20s6.88 0 8.59-.46a2.78 2.78 0 0 0 1.95-1.96A29 29 0 0 0 23 12a29 29 0 0 0-.46-5.58z"></path></svg>
            </a>
          </div>
        </div>
      </div>
    </section>

  </div>
</div>
`

	_, err := pool.Exec(ctx,
		`INSERT INTO themes (name, slug, description, thumbnail, theme_data, render_html, created_at, updated_at) 
		 VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`,
		name, slug, description, thumbnail, themeDataJSON, renderHTML,
	)
	if err != nil {
		log.Printf("Warning: failed to seed professional theme: %v", err)
		return
	}

	log.Println("Seeded professional dummy theme 'Royal Gold Premium' successfully!")
}
func seedThemes(pool *pgxpool.Pool) {
	seedRoyalGold(pool)
	seedModernSinis(pool)
	seedVintageRomance(pool)
}

func seedModernSinis(pool *pgxpool.Pool) {
	ctx := context.Background()
	_, _ = pool.Exec(ctx, "DELETE FROM themes WHERE slug = $1", "modern-sinis")

	name := "Modern Sinis"
	slug := "modern-sinis"
	description := "Tema pernikahan modern, clean, dan editorial dengan nuansa warna electric blue, putih-biru, dan layout bento-grid."
	thumbnail := "https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347308/wedding/ql04hsqybmcksxzrr3zu.jpg"

	themeDataJSON := `{
  "global": {
    "backgroundColor": "#F8F9FD",
    "containerWidth": "100%",
    "fontFamily": "Outfit, sans-serif",
    "primaryColor": "#2A07FF"
  },
  "splash": {
    "title": "WEDDING INVITATION",
    "heading": "Bagas & Arja",
    "fontFamily": "'BylinerScript_PERSONAL_USE_ONLY'",
    "bgImageUrl": "https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347308/wedding/ql04hsqybmcksxzrr3zu.jpg",
    "bgOverlayColor": "#0d1117",
    "bgOverlayOpacity": "0.7",
    "cardBgColor": "rgba(255, 255, 255, 0.9)",
    "cardBorderColor": "rgba(42, 7, 255, 0.15)",
    "cardBorderRadius": "24px",
    "cardTextColor": "#1e293b",
    "buttonText": "Open Invitation",
    "buttonBgColor": "#2A07FF",
    "buttonTextColor": "#ffffff",
    "logoColor": "#2A07FF"
  },
  "sections": [
    {
      "id": "sec-modern-cover",
      "name": "Cover & Opening",
      "paddingTop": "6rem",
      "paddingBottom": "6rem",
      "paddingLeft": "2rem",
      "paddingRight": "2rem",
      "marginTop": "0px",
      "marginBottom": "0px",
      "marginLeft": "0px",
      "marginRight": "0px",
      "backgroundColor": "#ffffff",
      "bgImageUrl": "https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347308/wedding/ql04hsqybmcksxzrr3zu.jpg",
      "bgImageSize": "cover",
      "bgImagePosition": "center",
      "bgImageRepeat": "no-repeat",
      "bgOverlayColor": "#0d1117",
      "bgOverlayOpacity": "0.6",
      "borderRadius": "0px",
      "borderWidth": "0px",
      "borderColor": "transparent",
      "borderStyle": "none",
      "boxShadow": "none",
      "transition": "zoom-in-parallax",
      "textColor": "#ffffff",
      "widgets": [
        {
          "id": "w-m-cover-icon",
          "type": "icon",
          "content": "ring",
          "style": {
            "margin": "0 auto 1.5rem",
            "color": "#2A07FF",
            "textAlign": "center",
            "display": "block"
          },
          "meta": { "size": "36", "strokeWidth": "2.5", "color": "#2A07FF" }
        },
        {
          "id": "w-m-cover-1",
          "type": "heading",
          "content": "THE WEDDING OF",
          "style": {
            "color": "#ffffff",
            "fontSize": "11px",
            "fontWeight": "800",
            "letterSpacing": "0.3em",
            "textAlign": "center",
            "textTransform": "uppercase"
          },
          "meta": { "level": "h4" }
        },
        {
          "id": "w-m-cover-2",
          "type": "heading",
          "content": "Bagas & Arja",
          "style": {
            "color": "#2A07FF",
            "fontFamily": "BylinerScript_PERSONAL_USE_ONLY",
            "fontSize": "58px",
            "fontWeight": "bold",
            "textAlign": "center",
            "margin": "1rem 0"
          },
          "meta": { "level": "h1" }
        },
        {
          "id": "w-m-cover-3",
          "type": "text",
          "content": "25.07.2026",
          "style": {
            "color": "#e2e8f0",
            "fontSize": "14px",
            "fontWeight": "600",
            "letterSpacing": "0.15em",
            "textAlign": "center"
          }
        },
        {
          "id": "w-m-cover-music",
          "type": "audio",
          "content": "https://res.cloudinary.com/dyuol7xfx/video/upload/v1782361688/wedding/x5bpdsj75b7hmxor8usu.mp3",
          "style": {
            "margin": "1.5rem auto"
          }
        },
        {
          "id": "w-m-cover-flower-left",
          "type": "floating_decor",
          "content": "https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347246/wedding/q2qgf7ktfml1ff9zupkb.jpg",
          "style": {
            "position": "absolute",
            "top": "12%",
            "left": "5px",
            "width": "60px",
            "height": "60px",
            "opacity": "0.75",
            "zIndex": "10",
            "customStyles": "filter: hue-rotate(180deg) saturate(2);"
          },
          "meta": { "animation": "drift", "speed": "7s", "delay": "0s" }
        },
        {
          "id": "w-m-cover-flower-right",
          "type": "floating_decor",
          "content": "https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347246/wedding/q2qgf7ktfml1ff9zupkb.jpg",
          "style": {
            "position": "absolute",
            "top": "20%",
            "right": "5px",
            "width": "50px",
            "height": "50px",
            "opacity": "0.7",
            "zIndex": "10",
            "transform": "scaleX(-1)",
            "customStyles": "filter: hue-rotate(180deg) saturate(2);"
          },
          "meta": { "animation": "drift", "speed": "9s", "delay": "1s" }
        }
      ]
    },
    {
      "id": "sec-modern-profile",
      "name": "Mempelai",
      "paddingTop": "5rem",
      "paddingBottom": "5rem",
      "paddingLeft": "2rem",
      "paddingRight": "2rem",
      "backgroundColor": "#ffffff",
      "transition": "slide-left",
      "textColor": "#1e293b",
      "widgets": [
        {
          "id": "w-m-prof-1",
          "type": "heading",
          "content": "MEET THE COUPLE",
          "style": {
            "color": "#2A07FF",
            "fontSize": "11px",
            "fontWeight": "800",
            "letterSpacing": "0.25em",
            "textAlign": "center",
            "marginBottom": "2rem"
          },
          "meta": { "level": "h4" }
        },
        {
          "id": "w-m-prof-2",
          "type": "image",
          "content": "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?auto=format&fit=crop&w=400&q=80",
          "style": {
            "width": "140px",
            "height": "180px",
            "margin": "0 auto 1.5rem auto"
          },
          "meta": { "frame": "modern-border" }
        },
        {
          "id": "w-m-prof-3",
          "type": "heading",
          "content": "Bagas Adinata",
          "style": {
            "fontSize": "22px",
            "fontWeight": "800",
            "textAlign": "center"
          },
          "meta": { "level": "h3" }
        },
        {
          "id": "w-m-prof-4",
          "type": "text",
          "content": "Putra pertama dari Bapak Heri Adinata & Ibu Rini Adinata",
          "style": {
            "fontSize": "13px",
            "color": "#64748b",
            "textAlign": "center",
            "marginBottom": "1rem"
          }
        },
        {
          "id": "w-m-prof-groom-social",
          "type": "social_links",
          "content": "",
          "style": {
            "gap": "12px",
            "justifyContent": "center",
            "marginTop": "0.5rem",
            "marginBottom": "2.5rem"
          },
          "meta": {
            "instagram": "https://instagram.com/bagas_adinata",
            "size": "20",
            "color": "#2A07FF"
          }
        },
        {
          "id": "w-m-prof-icon",
          "type": "icon",
          "content": "heart",
          "style": {
            "color": "#2A07FF",
            "textAlign": "center",
            "margin": "1.5rem 0"
          },
          "meta": { "size": "32", "strokeWidth": "2.5", "color": "#2A07FF" }
        },
        {
          "id": "w-m-prof-5",
          "type": "image",
          "content": "https://images.unsplash.com/photo-1494790108377-be9c29b29330?auto=format&fit=crop&w=400&q=80",
          "style": {
            "width": "140px",
            "height": "180px",
            "margin": "0 auto 1.5rem auto"
          },
          "meta": { "frame": "modern-border" }
        },
        {
          "id": "w-m-prof-6",
          "type": "heading",
          "content": "Arjanti Kirana",
          "style": {
            "fontSize": "22px",
            "fontWeight": "800",
            "textAlign": "center"
          },
          "meta": { "level": "h3" }
        },
        {
          "id": "w-m-prof-7",
          "type": "text",
          "content": "Putri kedua dari Bapak Bambang Kirana & Ibu Sita Kirana",
          "style": {
            "fontSize": "13px",
            "color": "#64748b",
            "textAlign": "center",
            "marginBottom": "1rem"
          }
        },
        {
          "id": "w-m-prof-bride-social",
          "type": "social_links",
          "content": "",
          "style": {
            "gap": "12px",
            "justifyContent": "center",
            "marginTop": "0.5rem"
          },
          "meta": {
            "instagram": "https://instagram.com/arja_kirana",
            "size": "20",
            "color": "#2A07FF"
          }
        }
      ]
    },
    {
      "id": "sec-modern-events",
      "name": "Acara Pernikahan",
      "paddingTop": "5rem",
      "paddingBottom": "5rem",
      "paddingLeft": "2rem",
      "paddingRight": "2rem",
      "backgroundColor": "#F8F9FD",
      "bgVideoUrl": "https://www.youtube.com/watch?v=F3x9T8S_BPE",
      "bgOverlayColor": "#F8F9FD",
      "bgOverlayOpacity": "0.85",
      "transition": "slide-right",
      "textColor": "#1e293b",
      "widgets": [
        {
          "id": "w-m-evt-icon",
          "type": "icon",
          "content": "calendar",
          "style": {
            "color": "#2A07FF",
            "textAlign": "center",
            "margin": "0 auto 1.5rem"
          },
          "meta": { "size": "32", "strokeWidth": "2.5", "color": "#2A07FF" }
        },
        {
          "id": "w-m-evt-1",
          "type": "heading",
          "content": "SAVE THE DATE",
          "style": {
            "color": "#2A07FF",
            "fontSize": "11px",
            "fontWeight": "800",
            "letterSpacing": "0.25em",
            "textAlign": "center",
            "marginBottom": "2rem"
          },
          "meta": { "level": "h4" }
        },
        {
          "id": "w-m-evt-2",
          "type": "container",
          "content": "Akad Nikah\nSabtu, 25 Juli 2026\n09:00 WIB - Selesai\nHotel Indonesia Kempinski, Jakarta",
          "style": {
            "backgroundColor": "#ffffff",
            "padding": "2.25rem",
            "borderRadius": "16px",
            "boxShadow": "6px 6px 0px #2A07FF",
            "border": "2.5px solid #000000",
            "margin": "0 auto 1.5rem auto",
            "maxWidth": "360px",
            "textAlign": "center",
            "fontWeight": "600"
          }
        },
        {
          "id": "w-m-evt-3",
          "type": "container",
          "content": "Resepsi Pernikahan\nSabtu, 25 Juli 2026\n18:30 WIB - Selesai\nGrand Ballroom Kempinski, Jakarta",
          "style": {
            "backgroundColor": "#ffffff",
            "padding": "2.25rem",
            "borderRadius": "16px",
            "boxShadow": "6px 6px 0px #2A07FF",
            "border": "2.5px solid #000000",
            "margin": "0 auto 1.5rem auto",
            "maxWidth": "360px",
            "textAlign": "center",
            "fontWeight": "600"
          }
        }
      ]
    },
    {
      "id": "sec-modern-gallery",
      "name": "Galeri Kenangan",
      "paddingTop": "5rem",
      "paddingBottom": "5rem",
      "paddingLeft": "2rem",
      "paddingRight": "2rem",
      "backgroundColor": "#ffffff",
      "transition": "zoom-in-parallax",
      "textColor": "#1e293b",
      "widgets": [
        {
          "id": "w-m-gal-icon",
          "type": "icon",
          "content": "star",
          "style": {
            "color": "#2A07FF",
            "textAlign": "center",
            "margin": "0 auto 1.5rem"
          },
          "meta": { "size": "32", "strokeWidth": "2.5", "color": "#2A07FF" }
        },
        {
          "id": "w-m-gal-1",
          "type": "heading",
          "content": "OUR GALLERY",
          "style": {
            "color": "#2A07FF",
            "fontSize": "11px",
            "fontWeight": "800",
            "letterSpacing": "0.25em",
            "textAlign": "center",
            "marginBottom": "2rem"
          },
          "meta": { "level": "h4" }
        },
        {
          "id": "w-m-gal-2",
          "type": "gallery",
          "content": "",
          "style": {
            "margin": "1rem auto"
          },
          "meta": {
            "cols": 2,
            "images": [
              "https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347308/wedding/ql04hsqybmcksxzrr3zu.jpg",
              "https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347235/wedding/tcjan0ozyqzxrchrm8yw.jpg",
              "https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347253/wedding/mxgeijfajqrr0loowcjo.jpg",
              "https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347246/wedding/q2qgf7ktfml1ff9zupkb.jpg"
            ]
          }
        }
      ]
    },
    {
      "id": "sec-modern-rsvp",
      "name": "RSVP",
      "paddingTop": "5rem",
      "paddingBottom": "5rem",
      "paddingLeft": "2rem",
      "paddingRight": "2rem",
      "backgroundColor": "#F8F9FD",
      "transition": "fade-up",
      "textColor": "#1e293b",
      "widgets": [
        {
          "id": "w-m-rsv-1",
          "type": "heading",
          "content": "CONFIRM KEHADIRAN",
          "style": {
            "color": "#2A07FF",
            "fontSize": "11px",
            "fontWeight": "800",
            "letterSpacing": "0.25em",
            "textAlign": "center",
            "marginBottom": "2rem"
          },
          "meta": { "level": "h4" }
        },
        {
          "id": "w-m-rsv-2",
          "type": "rsvp",
          "content": "",
          "style": {
            "maxWidth": "480px",
            "margin": "0 auto"
          }
        }
      ]
    },
    {
      "id": "sec-modern-gift",
      "name": "Kado",
      "paddingTop": "5rem",
      "paddingBottom": "5rem",
      "paddingLeft": "2rem",
      "paddingRight": "2rem",
      "backgroundColor": "#ffffff",
      "transition": "fade-in",
      "textColor": "#1e293b",
      "widgets": [
        {
          "id": "w-m-gft-icon",
          "type": "icon",
          "content": "gift",
          "style": {
            "color": "#2A07FF",
            "textAlign": "center",
            "margin": "0 auto 1.5rem"
          },
          "meta": { "size": "32", "strokeWidth": "2.5", "color": "#2A07FF" }
        },
        {
          "id": "w-m-gft-1",
          "type": "heading",
          "content": "WEDDING GIFT",
          "style": {
            "color": "#2A07FF",
            "fontSize": "11px",
            "fontWeight": "800",
            "letterSpacing": "0.25em",
            "textAlign": "center",
            "marginBottom": "2rem"
          },
          "meta": { "level": "h4" }
        },
        {
          "id": "w-m-gft-2",
          "type": "container",
          "content": "Bank Mandiri\nNo. Rek: 123-45678-90\na.n Bagas Adinata",
          "style": {
            "backgroundColor": "#F8F9FD",
            "border": "1px solid rgba(42, 7, 255, 0.08)",
            "padding": "1.5rem",
            "borderRadius": "12px",
            "textAlign": "center",
            "maxWidth": "320px",
            "margin": "1rem auto"
          }
        }
      ]
    },
    {
      "id": "sec-modern-footer",
      "name": "Footer",
      "paddingTop": "4rem",
      "paddingBottom": "4rem",
      "paddingLeft": "2rem",
      "paddingRight": "2rem",
      "backgroundColor": "#0d1117",
      "transition": "fade-up",
      "textColor": "#ffffff",
      "widgets": [
        {
          "id": "w-m-ftr-icon",
          "type": "icon",
          "content": "heart",
          "style": {
            "color": "#2A07FF",
            "textAlign": "center",
            "margin": "0 auto 1.5rem",
            "display": "block"
          },
          "meta": { "size": "36", "strokeWidth": "2.5", "color": "#2A07FF" }
        },
        {
          "id": "w-m-ftr-1",
          "type": "heading",
          "content": "Thank You",
          "style": {
            "color": "#2A07FF",
            "fontFamily": "BylinerScript_PERSONAL_USE_ONLY",
            "fontSize": "44px",
            "textAlign": "center"
          },
          "meta": { "level": "h2" }
        },
        {
          "id": "w-m-ftr-2",
          "type": "text",
          "content": "Bagas & Arja",
          "style": {
            "color": "#94a3b8",
            "fontSize": "14px",
            "fontWeight": "600",
            "letterSpacing": "0.1em",
            "textAlign": "center",
            "marginTop": "0.5rem"
          }
        },
        {
          "id": "w-m-ftr-socials",
          "type": "social_links",
          "content": "",
          "style": {
            "gap": "16px",
            "justifyContent": "center",
            "marginTop": "1.5rem"
          },
          "meta": {
            "instagram": "https://instagram.com/bagas_arja",
            "tiktok": "https://tiktok.com/@bagas_arja",
            "whatsapp": "https://wa.me/62812345678",
            "size": "24",
            "color": "#2A07FF"
          }
        }
      ]
    }
  ]
}`

	renderHTML := `
<style>
  @import url('https://fonts.googleapis.com/css2?family=Outfit:wght@300;400;500;600;700;800&display=swap');
  @font-face {
    font-family: 'BylinerScript_PERSONAL_USE_ONLY';
    src: url('https://res.cloudinary.com/dyuol7xfx/raw/upload/v1782361588/wedding/x5bpdsj75b7hmxor8usu');
    font-weight: normal;
    font-style: normal;
    font-display: swap;
  }
  body { margin: 0; padding: 0; font-family: 'Outfit', sans-serif; background-color: #F8F9FD; }
  * { box-sizing: border-box; }
  .list-disc { list-style-type: disc; }
  .list-decimal { list-style-type: decimal; }

  /* Animation System */
  @keyframes floatSway {
    0% { transform: translateY(0px) rotate(0deg); }
    50% { transform: translateY(-15px) rotate(6deg); }
    100% { transform: translateY(0px) rotate(0deg); }
  }
  @keyframes floatSwaySlow {
    0% { transform: translateY(0px) rotate(0deg); }
    50% { transform: translateY(-25px) rotate(-10deg); }
    100% { transform: translateY(0px) rotate(0deg); }
  }
  @keyframes spinSlow {
    0% { transform: rotate(0deg); }
    100% { transform: rotate(360deg); }
  }
  @keyframes pulseSlow {
    0%, 100% { transform: scale(1); }
    50% { transform: scale(1.12); }
  }
  @keyframes driftLeftRight {
    0%, 100% { transform: translateX(0px) rotate(0deg); }
    50% { transform: translateX(20px) rotate(8deg); }
    100% { transform: translateX(0px) rotate(0deg); }
  }
  .animate-float-sway { animation: floatSway var(--float-speed, 6s) ease-in-out infinite var(--float-delay, 0s); }
  .animate-float-sway-slow { animation: floatSwaySlow var(--float-speed, 10s) ease-in-out infinite var(--float-delay, 0s); }
  .animate-spin-slow { animation: spinSlow var(--float-speed, 15s) linear infinite var(--float-delay, 0s); }
  .animate-pulse-slow { animation: pulseSlow var(--float-speed, 5s) ease-in-out infinite var(--float-delay, 0s); }
  .animate-drift-lr { animation: driftLeftRight var(--float-speed, 8s) ease-in-out infinite var(--float-delay, 0s); }

  @keyframes fadeUp {
    from { opacity: 0; transform: translateY(30px); }
    to { opacity: 1; transform: translateY(0); }
  }
  @keyframes fadeIn {
    from { opacity: 0; }
    to { opacity: 1; }
  }
  @keyframes slideLeft {
    from { opacity: 0; transform: translateX(-30px); }
    to { opacity: 1; transform: translateX(0); }
  }
  @keyframes slideRight {
    from { opacity: 0; transform: translateX(30px); }
    to { opacity: 1; transform: translateX(0); }
  }
  @keyframes zoomIn {
    from { opacity: 0; transform: scale(0.95); }
    to { opacity: 1; transform: scale(1); }
  }

  .animate-widget {
    opacity: 0;
    transition: opacity 0.8s cubic-bezier(0.16, 1, 0.3, 1), transform 0.8s cubic-bezier(0.16, 1, 0.3, 1);
  }

  .is-in-view[data-transition="fade-up"] .animate-widget,
  .is-in-view[data-transition="fade-up-parallax"] .animate-widget {
    animation: fadeUp 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
  }
  .is-in-view[data-transition="fade-in"] .animate-widget {
    animation: fadeIn 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
  }
  .is-in-view[data-transition="slide-left"] .animate-widget {
    animation: slideLeft 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
  }
  .is-in-view[data-transition="slide-right"] .animate-widget {
    animation: slideRight 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
  }
  .is-in-view[data-transition="zoom-in"] .animate-widget,
  .is-in-view[data-transition="zoom-in-parallax"] .animate-widget {
    animation: zoomIn 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
  }

  /* Parallax Background Layout */
  .parallax-section {
    position: relative;
    overflow: hidden;
  }
  .parallax-bg {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 125%; /* Taller for translation bounds */
    z-index: 0;
    pointer-events: none;
    background-size: cover;
    background-position: center;
    background-repeat: no-repeat;
    will-change: transform;
  }

  @media (prefers-reduced-motion: reduce) {
    .animate-widget {
      opacity: 1 !important;
      transform: none !important;
      animation: none !important;
      transition: none !important;
    }
    .parallax-bg {
      transform: none !important;
      height: 100% !important;
    }
  }
</style>
<div style="font-family: 'Outfit', sans-serif; background-color: #F8F9FD; min-height: 100vh; width: 100%;">
  <div style="max-width: 100%; margin: 0 auto; min-height: 100vh; box-shadow: 0 10px 50px rgba(0,0,0,0.04); background: #ffffff;">

    <!-- Cover Section -->
    <section id="sec-modern-cover" class="parallax-section" data-transition="zoom-in-parallax" style="position: relative; padding: 6rem 2rem; text-align: center; min-height: 100vh; color: #ffffff; display: flex; flex-direction: column; align-items: center; justify-content: center; overflow: hidden; background-color: #ffffff;">
      <div class="parallax-bg" style="background-image: url('https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347308/wedding/ql04hsqybmcksxzrr3zu.jpg'); background-size: cover; background-position: center; transform: translateY(0px) scale(1.15);"></div>
      <div style="position: absolute; top: 0; left: 0; width: 100%; height: 100%; background-color: #0d1117; opacity: 0.6; z-index: 1; pointer-events: none;"></div>
      
      <!-- Floating Blue Leaf Ornaments -->
      <div class="animate-drift-lr" style="position: absolute; left: 5px; top: 12%; width: 60px; height: 60px; z-index: 10; pointer-events: none; opacity: 0.75; --float-speed: 7s; --float-delay: 0s; filter: hue-rotate(180deg) saturate(2);">
        <img src="https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347246/wedding/q2qgf7ktfml1ff9zupkb.jpg" style="width: 100%; height: 100%; object-fit: contain;" alt="Decor" />
      </div>
      <div class="animate-drift-lr" style="position: absolute; right: 5px; top: 20%; width: 50px; height: 50px; z-index: 10; pointer-events: none; opacity: 0.7; transform: scaleX(-1); --float-speed: 9s; --float-delay: 1s; filter: hue-rotate(180deg) saturate(2);">
        <img src="https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347246/wedding/q2qgf7ktfml1ff9zupkb.jpg" style="width: 100%; height: 100%; object-fit: contain; transform: scaleX(-1);" alt="Decor" />
      </div>

      <div style="position: relative; z-index: 2; width: 100%;">
        <!-- Ring Icon -->
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%; display: flex; justify-content: center; margin: 0 auto 1.5rem;">
          <svg viewBox="0 0 24 24" width="36" height="36" stroke="#2A07FF" stroke-width="2.5" fill="none" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="14" r="7"/><path d="M12 2a4 4 0 0 1 4 4v1H8V6a4 4 0 0 1 4-4Z"/></svg>
        </div>

        <div class="animate-widget" style="animation-delay: 0.12s; transition-delay: 0.12s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <h4 style="color: #ffffff; font-size: 11px; font-weight: 800; letter-spacing: 0.3em; text-align: center; text-transform: uppercase; margin-top: 0;">THE WEDDING OF</h4>
        </div>
        <div class="animate-widget" style="animation-delay: 0.24s; transition-delay: 0.24s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <h1 style="color: #2A07FF; font-family: 'BylinerScript_PERSONAL_USE_ONLY'; font-size: 58px; font-weight: bold; text-align: center; margin: 1rem 0;">Bagas & Arja</h1>
        </div>
        <div class="animate-widget" style="animation-delay: 0.36s; transition-delay: 0.36s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <p style="color: #e2e8f0; font-size: 14px; font-weight: 600; letter-spacing: 0.15em; text-align: center;">25.07.2026</p>
        </div>

        <div class="animate-widget" style="animation-delay: 0.48s; transition-delay: 0.48s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <div style="background: rgba(255,255,255,0.15); border: 1px solid rgba(255,255,255,0.2); padding: 0.75rem 1.5rem; border-radius: 99px; display: inline-flex; align-items: center; gap: 1rem; margin: 1.5rem auto; backdrop-filter: blur(5px);">
            <div style="width: 28px; height: 28px; border-radius: 50%; background: #2A07FF; display: flex; align-items: center; justify-content: center; color: white; font-weight: bold; cursor: pointer;">▶</div>
            <span style="font-size: 13px; color: #ffffff; font-weight: 500;">Background Music</span>
          </div>
        </div>
      </div>
    </section>

    <!-- Mempelai Section -->
    <section id="sec-modern-profile" data-transition="slide-left" style="padding: 5rem 2rem; background-color: #ffffff; text-align: center; color: #1e293b; position: relative;">
      <div style="position: relative; z-index: 2; width: 100%;">
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <h4 style="color: #2A07FF; font-size: 11px; font-weight: 800; letter-spacing: 0.25em; text-align: center; margin-bottom: 2rem;">MEET THE COUPLE</h4>
        </div>
        
        <!-- Groom Box (modern-border frame) -->
        <div class="animate-widget" style="animation-delay: 0.12s; transition-delay: 0.12s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <div style="border: 4px solid #1e293b; box-shadow: 8px 8px 0px #2A07FF; display: inline-block; margin: 0 auto 1.5rem auto; max-width: 100%; width: 140px; height: 180px; overflow: hidden;">
            <img src="https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?auto=format&fit=crop&w=400&q=80" style="display: block; width: 100%; height: 100%; object-fit: cover;" alt="Bagas" />
          </div>
        </div>
        <div class="animate-widget" style="animation-delay: 0.24s; transition-delay: 0.24s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <h3 style="font-size: 22px; font-weight: 800; text-align: center; margin-top: 0;">Bagas Adinata</h3>
        </div>
        <div class="animate-widget" style="animation-delay: 0.36s; transition-delay: 0.36s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <p style="font-size: 13px; color: #64748b; text-align: center; margin-bottom: 1rem;">Putra pertama dari Bapak Heri Adinata & Ibu Rini Adinata</p>
        </div>
        
        <!-- Groom Social Row -->
        <div class="animate-widget" style="animation-delay: 0.48s; transition-delay: 0.48s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <div style="display: flex; gap: 12px; justify-content: center; margin-bottom: 2.5rem;">
            <a href="https://instagram.com/bagas_adinata" target="_blank" rel="noopener noreferrer" style="display: inline-flex; align-items: center; justify-content: center; width: 36px; height: 36px; border-radius: 50%; border: 1.5px solid #2A07FF; color: #2A07FF; text-decoration: none; transition: all 0.3s ease;">
              <svg viewBox="0 0 24 24" width="20" height="20" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><rect x="2" y="2" width="20" height="20" rx="5" ry="5"></rect><path d="M16 11.37A4 4 0 1 1 12.63 8 4 4 0 0 1 16 11.37z"></path><line x1="17.5" y1="6.5" x2="17.51" y2="6.5"></line></svg>
            </a>
          </div>
        </div>

        <!-- Heart Icon -->
        <div class="animate-widget" style="animation-delay: 0.6s; transition-delay: 0.6s; width: 100%; display: flex; justify-content: center; margin: 1.5rem 0;">
          <svg viewBox="0 0 24 24" width="32" height="32" stroke="#2A07FF" stroke-width="2.5" fill="none" stroke-linecap="round" stroke-linejoin="round"><path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/></svg>
        </div>

        <!-- Bride Box (modern-border frame) -->
        <div class="animate-widget" style="animation-delay: 0.72s; transition-delay: 0.72s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <div style="border: 4px solid #1e293b; box-shadow: 8px 8px 0px #2A07FF; display: inline-block; margin: 0 auto 1.5rem auto; max-width: 100%; width: 140px; height: 180px; overflow: hidden;">
            <img src="https://images.unsplash.com/photo-1494790108377-be9c29b29330?auto=format&fit=crop&w=400&q=80" style="display: block; width: 100%; height: 100%; object-fit: cover;" alt="Arja" />
          </div>
        </div>
        <div class="animate-widget" style="animation-delay: 0.84s; transition-delay: 0.84s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <h3 style="font-size: 22px; font-weight: 800; text-align: center; margin-top: 0;">Arjanti Kirana</h3>
        </div>
        <div class="animate-widget" style="animation-delay: 0.96s; transition-delay: 0.96s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <p style="font-size: 13px; color: #64748b; text-align: center; marginBottom: 1rem;">Putri kedua dari Bapak Bambang Kirana & Ibu Sita Kirana</p>
        </div>
        
        <!-- Bride Social Row -->
        <div class="animate-widget" style="animation-delay: 1.08s; transition-delay: 1.08s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <div style="display: flex; gap: 12px; justify-content: center;">
            <a href="https://instagram.com/arja_kirana" target="_blank" rel="noopener noreferrer" style="display: inline-flex; align-items: center; justify-content: center; width: 36px; height: 36px; border-radius: 50%; border: 1.5px solid #2A07FF; color: #2A07FF; text-decoration: none; transition: all 0.3s ease;">
              <svg viewBox="0 0 24 24" width="20" height="20" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><rect x="2" y="2" width="20" height="20" rx="5" ry="5"></rect><path d="M16 11.37A4 4 0 1 1 12.63 8 4 4 0 0 1 16 11.37z"></path><line x1="17.5" y1="6.5" x2="17.51" y2="6.5"></line></svg>
            </a>
          </div>
        </div>
      </div>
    </section>

    <!-- Events Section -->
    <section id="sec-modern-events" data-transition="slide-right" style="padding: 5rem 2rem; background-color: #F8F9FD; text-align: center; color: #1e293b; position: relative; overflow: hidden;">
      <!-- Embedded YouTube Video Loop Background -->
      <iframe src="https://www.youtube.com/embed/F3x9T8S_BPE?autoplay=1&mute=1&loop=1&playlist=F3x9T8S_BPE&controls=0&showinfo=0&rel=0&iv_load_policy=3&playsinline=1&enablejsapi=1" style="position: absolute; top:0; left:0; width:100%; height:100%; border:none; object-fit: cover; opacity: 0.15; pointer-events:none; z-index:0;" allow="autoplay; encrypted-media"></iframe>
      <div style="position: absolute; top: 0; left: 0; width: 100%; height: 100%; background-color: #F8F9FD; opacity: 0.85; z-index: 1; pointer-events: none;"></div>

      <div style="position: relative; z-index: 2; width: 100%;">
        <!-- Calendar Icon -->
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%; display: flex; justify-content: center; margin: 0 auto 1.5rem;">
          <svg viewBox="0 0 24 24" width="32" height="32" stroke="#2A07FF" stroke-width="2.5" fill="none" stroke-linecap="round" stroke-linejoin="round"><rect width="18" height="18" x="3" y="4" rx="2" ry="2"/><line x1="16" x2="16" y1="2" y2="6"/><line x1="8" x2="8" y1="2" y2="6"/><line x1="3" x2="21" y1="10" y2="10"/></svg>
        </div>
        <div class="animate-widget" style="animation-delay: 0.12s; transition-delay: 0.12s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <h4 style="color: #2A07FF; font-size: 11px; font-weight: 800; letter-spacing: 0.25em; text-align: center; margin-bottom: 2rem; text-transform: uppercase;">SAVE THE DATE</h4>
        </div>
        
        <div style="display: flex; flex-direction: column; gap: 2rem; max-width: 400px; margin: 0 auto;">
          <!-- Akad Bento Card -->
          <div class="animate-widget" style="animation-delay: 0.24s; transition-delay: 0.24s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
            <div style="background-color: #ffffff; padding: 2.25rem; border-radius: 16px; border: 2.5px solid #000000; box-shadow: 6px 6px 0px #2A07FF; text-align: left;">
              <span style="background: #2A07FF; color: #ffffff; font-size: 10px; font-weight: 800; padding: 4px 10px; border-radius: 99px; letter-spacing: 0.05em; display: inline-block; margin-bottom: 1rem;">AKAD NIKAH</span>
              <h3 style="font-size: 24px; font-weight: 800; color: #000000; margin: 0 0 0.5rem 0; letter-spacing: -0.02em;">Hotel Indonesia Kempinski</h3>
              <p style="margin: 0.25rem 0; font-size: 14px; font-weight: 700; color: #2A07FF;">Sabtu, 25 Juli 2026 — 09:00 WIB</p>
              <p style="margin: 1rem 0 0 0; font-size: 13px; color: #64748b; line-height: 1.6;">Lantai 2, Room Heritage, Jl. M.H. Thamrin No.1, Menteng, Jakarta Pusat</p>
            </div>
          </div>

          <!-- Resepsi Bento Card -->
          <div class="animate-widget" style="animation-delay: 0.36s; transition-delay: 0.36s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
            <div style="background-color: #ffffff; padding: 2.25rem; border-radius: 16px; border: 2.5px solid #000000; box-shadow: 6px 6px 0px #2A07FF; text-align: left;">
              <span style="background: #000000; color: #ffffff; font-size: 10px; font-weight: 800; padding: 4px 10px; border-radius: 99px; letter-spacing: 0.05em; display: inline-block; margin-bottom: 1rem;">RESEPSI</span>
              <h3 style="font-size: 24px; font-weight: 800; color: #000000; margin: 0 0 0.5rem 0; letter-spacing: -0.02em;">Grand Ballroom Kempinski</h3>
              <p style="margin: 0.25rem 0; font-size: 14px; font-weight: 700; color: #2A07FF;">Sabtu, 25 Juli 2026 — 18:30 WIB</p>
              <p style="margin: 1rem 0 0 0; font-size: 13px; color: #64748b; line-height: 1.6;">Lantai Dasar, Lobby Utama, Jl. M.H. Thamrin No.1, Menteng, Jakarta Pusat</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Gallery Section -->
    <section id="sec-modern-gallery" class="parallax-section" data-transition="zoom-in-parallax" style="padding: 5rem 2rem; background-color: #ffffff; text-align: center; color: #1e293b; position: relative;">
      <div style="position: relative; z-index: 2; width: 100%;">
        <!-- Star Icon -->
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%; display: flex; justify-content: center; margin: 0 auto 1.5rem;">
          <svg viewBox="0 0 24 24" width="32" height="32" stroke="#2A07FF" stroke-width="2.5" fill="none" stroke-linecap="round" stroke-linejoin="round"><polygon points="12 2 15.09 8.26 22 9.27 17 14.14 18.18 21.02 12 17.77 5.82 21.02 7 14.14 2 9.27 8.91 8.26 12 2"/></svg>
        </div>
        <div class="animate-widget" style="animation-delay: 0.12s; transition-delay: 0.12s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <h4 style="color: #2A07FF; font-size: 11px; font-weight: 800; letter-spacing: 0.25em; text-align: center; margin-bottom: 2rem; text-transform: uppercase;">OUR GALLERY</h4>
        </div>
        
        <!-- Asymmetric Bento Grid for Gallery -->
        <div class="animate-widget" style="animation-delay: 0.24s; transition-delay: 0.24s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <div style="display: grid; grid-template-columns: 1.2fr 0.8fr; gap: 1rem; margin: 2rem auto; max-width: 500px; width: 100%;">
            <!-- Grid 1: Wide Banner -->
            <div style="grid-column: span 2; border: 2px solid #000000; border-radius: 12px; overflow: hidden; height: 180px; box-shadow: 4px 4px 0px rgba(0,0,0,0.15);">
              <img src="https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347308/wedding/ql04hsqybmcksxzrr3zu.jpg" style="width: 100%; height: 100%; object-fit: cover; transition: transform 0.5s;" onmouseover="this.style.transform='scale(1.05)'" onmouseout="this.style.transform='scale(1)'" alt="Gallery 1" />
            </div>
            <!-- Grid 2: Tall Left -->
            <div style="border: 2px solid #000000; border-radius: 12px; overflow: hidden; height: 220px; box-shadow: 4px 4px 0px rgba(0,0,0,0.15);">
              <img src="https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347235/wedding/tcjan0ozyqzxrchrm8yw.jpg" style="width: 100%; height: 100%; object-fit: cover; transition: transform 0.5s;" onmouseover="this.style.transform='scale(1.05)'" onmouseout="this.style.transform='scale(1)'" alt="Gallery 2" />
            </div>
            <!-- Grid 3: Stacked Right column -->
            <div style="display: flex; flex-direction: column; gap: 1rem; height: 220px;">
              <div style="border: 2px solid #000000; border-radius: 12px; overflow: hidden; height: 102px; box-shadow: 4px 4px 0px rgba(0,0,0,0.15); width: 100%;">
                <img src="https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347253/wedding/mxgeijfajqrr0loowcjo.jpg" style="width: 100%; height: 100%; object-fit: cover; transition: transform 0.5s;" onmouseover="this.style.transform='scale(1.05)'" onmouseout="this.style.transform='scale(1)'" alt="Gallery 3" />
              </div>
              <div style="border: 2px solid #000000; border-radius: 12px; overflow: hidden; height: 102px; box-shadow: 4px 4px 0px rgba(0,0,0,0.15); width: 100%;">
                <img src="https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347246/wedding/q2qgf7ktfml1ff9zupkb.jpg" style="width: 100%; height: 100%; object-fit: cover; transition: transform 0.5s;" onmouseover="this.style.transform='scale(1.05)'" onmouseout="this.style.transform='scale(1)'" alt="Gallery 4" />
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- RSVP Section -->
    <section id="sec-modern-rsvp" data-transition="fade-up" style="padding: 5rem 2rem; background-color: #F8F9FD; text-align: center; color: #1e293b; position: relative;">
      <div style="position: relative; z-index: 2; width: 100%;">
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <h4 style="color: #2A07FF; font-size: 11px; font-weight: 800; letter-spacing: 0.25em; text-align: center; margin-bottom: 2rem; text-transform: uppercase;">CONFIRM KEHADIRAN</h4>
        </div>
        
        <div class="animate-widget" style="animation-delay: 0.12s; transition-delay: 0.12s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <div style="background: #ffffff; padding: 2.5rem; border-radius: 20px; max-width: 480px; margin: 0 auto; text-align: left; box-shadow: 8px 8px 0px #000000; border: 3px solid #000000;">
            <form onsubmit="event.preventDefault(); alert('Terima kasih atas konfirmasi Anda!')">
              <div style="margin-bottom: 1.5rem;">
                <label style="display: block; font-size: 0.75rem; font-weight: 800; margin-bottom: 0.5rem; color: #000000; text-transform: uppercase; letter-spacing: 0.05em;">Nama Lengkap</label>
                <input type="text" placeholder="Masukkan nama..." required style="width: 100%; height: 46px; padding: 0 1rem; border-radius: 8px; border: 2.5px solid #000000; outline: none; background: #ffffff; font-weight: 700; font-size: 14px; box-shadow: 3px 3px 0px rgba(0,0,0,0.1);" />
              </div>
              
              <div style="margin-bottom: 1.5rem;">
                <label style="display: block; font-size: 0.75rem; font-weight: 800; margin-bottom: 0.5rem; color: #000000; text-transform: uppercase; letter-spacing: 0.05em;">Pesan / Ucapan Hangat</label>
                <textarea rows="3" placeholder="Tuliskan ucapan selamat..." required style="width: 100%; padding: 0.75rem 1rem; border-radius: 8px; border: 2.5px solid #000000; outline: none; background: #ffffff; resize: none; font-weight: 700; font-size: 14px; box-shadow: 3px 3px 0px rgba(0,0,0,0.1);"></textarea>
              </div>
              
              <button type="submit" style="width: 100%; height: 48px; border-radius: 8px; background: #2A07FF; color: #ffffff; font-weight: 800; border: 2.5px solid #000000; cursor: pointer; text-transform: uppercase; font-size: 12px; letter-spacing: 0.05em; box-shadow: 4px 4px 0px #000000; transition: all 0.1s;" onmouseover="this.style.transform='translate(-2px, -2px)'; this.style.boxShadow='6px 6px 0px #000000';" onmouseout="this.style.transform='none'; this.style.boxShadow='4px 4px 0px #000000';">Kirim Konfirmasi</button>
            </form>
          </div>
        </div>
      </div>
    </section>

    <!-- Gift Section -->
    <section id="sec-modern-gift" data-transition="fade-in" style="padding: 5rem 2rem; background-color: #ffffff; text-align: center; color: #1e293b; position: relative;">
      <div style="position: relative; z-index: 2; width: 100%;">
        <!-- Gift Icon -->
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%; display: flex; justify-content: center; margin: 0 auto 1.5rem;">
          <svg viewBox="0 0 24 24" width="32" height="32" stroke="#2A07FF" stroke-width="2.5" fill="none" stroke-linecap="round" stroke-linejoin="round"><rect width="18" height="14" x="3" y="8" rx="2"/><path d="M12 5a3 3 0 1 0-3 3h6a3 3 0 1 0-3-3Z"/><path d="M12 2v20"/><path d="M3 12h18"/></svg>
        </div>
        <div class="animate-widget" style="animation-delay: 0.12s; transition-delay: 0.12s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <h4 style="color: #2A07FF; font-size: 11px; font-weight: 800; letter-spacing: 0.25em; text-align: center; margin-bottom: 2rem; text-transform: uppercase;">WEDDING GIFT</h4>
        </div>
        
        <div class="animate-widget" style="animation-delay: 0.24s; transition-delay: 0.24s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <!-- BCA / Credit Card Style -->
          <div style="background: linear-gradient(135deg, #0d1117 0%, #1e293b 100%); border: 2.5px solid #000000; border-radius: 16px; padding: 1.75rem; text-align: left; max-width: 340px; margin: 1.5rem auto; color: #ffffff; box-shadow: 6px 6px 0px #2A07FF; position: relative; overflow: hidden;">
            <div style="position: absolute; right: -20px; bottom: -20px; width: 120px; height: 120px; background: rgba(42, 7, 255, 0.15); border-radius: 50%; filter: blur(20px); pointer-events: none;"></div>
            <p style="margin: 0; font-size: 10px; font-weight: 800; letter-spacing: 0.1em; color: #38bdf8;">DIGITAL GIFT CARD</p>
            <p style="margin: 1.5rem 0 0.5rem 0; font-size: 22px; font-weight: 800; letter-spacing: 0.05em; font-family: monospace;">123-45678-90</p>
            <p style="margin: 0; font-size: 11px; color: #94a3b8; text-transform: uppercase;">BANK MANDIRI — a.n Bagas Adinata</p>
            <button onclick="navigator.clipboard.writeText('123-45678-90'); const btn = this; const orig = btn.innerText; btn.innerText = 'COPIED!'; setTimeout(() => btn.innerText = orig, 2000);" style="position: absolute; right: 1.5rem; top: 1.5rem; background: #2A07FF; color: white; border: 1.5px solid #ffffff; padding: 4px 10px; font-size: 9px; font-weight: 800; border-radius: 6px; cursor: pointer; letter-spacing: 0.05em;">COPY</button>
          </div>
        </div>
      </div>
    </section>

    <!-- Footer Section -->
    <section id="sec-modern-footer" data-transition="fade-up" style="padding: 6rem 2rem; background-color: #000000; text-align: center; color: #ffffff; position: relative;">
      <div style="position: relative; z-index: 2; width: 100%;">
        <!-- Heart Icon -->
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%; display: flex; justify-content: center; margin: 0 auto 1.5rem;">
          <svg viewBox="0 0 24 24" width="36" height="36" stroke="#2A07FF" stroke-width="2.5" fill="none" stroke-linecap="round" stroke-linejoin="round"><path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/></svg>
        </div>
        <div class="animate-widget" style="animation-delay: 0.12s; transition-delay: 0.12s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <h2 style="color: #2A07FF; font-family: 'BylinerScript_PERSONAL_USE_ONLY'; font-size: 56px; text-align: center; margin-top: 0; margin-bottom: 0;">Thank You</h2>
        </div>
        <div class="animate-widget" style="animation-delay: 0.24s; transition-delay: 0.24s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <p style="color: #ffffff; font-size: 16px; font-weight: 900; letter-spacing: 0.3em; text-align: center; margin-top: 1.5rem; text-transform: uppercase;">BAGAS & ARJA</p>
        </div>
        
        <!-- Social Media Link Row -->
        <div class="animate-widget" style="animation-delay: 0.36s; transition-delay: 0.36s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <div style="display: flex; gap: 16px; justify-content: center; margin-top: 1.5rem;">
            <a href="https://instagram.com/bagas_arja" target="_blank" rel="noopener noreferrer" style="display: inline-flex; align-items: center; justify-content: center; width: 40px; height: 40px; border-radius: 50%; border: 1.5px solid #2A07FF; color: #2A07FF; text-decoration: none; transition: all 0.3s ease;">
              <svg viewBox="0 0 24 24" width="24" height="24" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><rect x="2" y="2" width="20" height="20" rx="5" ry="5"></rect><path d="M16 11.37A4 4 0 1 1 12.63 8 4 4 0 0 1 16 11.37z"></path><line x1="17.5" y1="6.5" x2="17.51" y2="6.5"></line></svg>
            </a>
            <a href="https://tiktok.com/@bagas_arja" target="_blank" rel="noopener noreferrer" style="display: inline-flex; align-items: center; justify-content: center; width: 40px; height: 40px; border-radius: 50%; border: 1.5px solid #2A07FF; color: #2A07FF; text-decoration: none; transition: all 0.3s ease;">
              <svg viewBox="0 0 24 24" width="24" height="24" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><path d="M9 12a4 4 0 1 0 4 4V4a5 5 0 0 0 5 5"></path></svg>
            </a>
            <a href="https://wa.me/62812345678" target="_blank" rel="noopener noreferrer" style="display: inline-flex; align-items: center; justify-content: center; width: 40px; height: 40px; border-radius: 50%; border: 1.5px solid #2A07FF; color: #2A07FF; text-decoration: none; transition: all 0.3s ease;">
              <svg viewBox="0 0 24 24" width="24" height="24" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><path d="M21 11.5a8.38 8.38 0 0 1-.9 3.8 8.5 8.5 0 0 1-7.6 4.7 8.38 8.38 0 0 1-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 0 1-.9-3.8 8.5 8.5 0 0 1 4.7-7.6 8.38 8.38 0 0 1 3.8-.9h.5a8.48 8.48 0 0 1 8 8v.5z"></path></svg>
            </a>
          </div>
        </div>
      </div>
    </section>

  </div>
</div>
`

	_, err := pool.Exec(ctx,
		`INSERT INTO themes (name, slug, description, thumbnail, theme_data, render_html, created_at, updated_at) 
		 VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`,
		name, slug, description, thumbnail, themeDataJSON, renderHTML,
	)
	if err != nil {
		log.Printf("Warning: failed to seed modern theme: %v", err)
		return
	}
	log.Println("Seeded professional dummy theme 'Modern Sinis' successfully!")
}

func seedVintageRomance(pool *pgxpool.Pool) {
	ctx := context.Background()
	_, _ = pool.Exec(ctx, "DELETE FROM themes WHERE slug = $1", "vintage-romance")

	name := "Vintage Romance"
	slug := "vintage-romance"
	description := "Tema pernikahan klasik romantis dengan nuansa warna sepia hangat, tekstur kertas tua, dan typography serif elegan."
	thumbnail := "https://images.unsplash.com/photo-1465495976277-4387d4b0b4c6?auto=format&fit=crop&w=400&q=80"

	themeDataJSON := `{
  "global": {
    "backgroundColor": "#fdfbf7",
    "containerWidth": "100%",
    "fontFamily": "'Playfair Display', serif",
    "primaryColor": "#8c6239"
  },
  "splash": {
    "title": "THE MARRIAGE OF",
    "heading": "Carter & Evelyn",
    "fontFamily": "'LatenzaScript_PERSONAL_USE_ONLY'",
    "bgImageUrl": "https://images.unsplash.com/photo-1465495976277-4387d4b0b4c6?auto=format&fit=crop&w=1200&q=80",
    "bgOverlayColor": "#292524",
    "bgOverlayOpacity": "0.8",
    "cardBgColor": "rgba(253, 251, 247, 0.9)",
    "cardBorderColor": "rgba(140, 98, 57, 0.2)",
    "cardBorderRadius": "16px",
    "cardTextColor": "#44403c",
    "buttonText": "Enter Invitation",
    "buttonBgColor": "#8c6239",
    "buttonTextColor": "#ffffff",
    "logoColor": "#8c6239"
  },
  "sections": [
    {
      "id": "sec-vintage-cover",
      "name": "Cover & Opening",
      "paddingTop": "6rem",
      "paddingBottom": "6rem",
      "paddingLeft": "2rem",
      "paddingRight": "2rem",
      "backgroundColor": "#fdfbf7",
      "bgImageUrl": "https://images.unsplash.com/photo-1465495976277-4387d4b0b4c6?auto=format&fit=crop&w=1200&q=80",
      "bgVideoUrl": "https://www.youtube.com/watch?v=F3x9T8S_BPE",
      "bgImageSize": "cover",
      "bgImagePosition": "center",
      "bgOverlayColor": "#292524",
      "bgOverlayOpacity": "0.6",
      "transition": "fade-up-parallax",
      "textColor": "#ffffff",
      "widgets": [
        {
          "id": "w-v-cover-audio",
          "type": "audio",
          "content": "https://res.cloudinary.com/dyuol7xfx/video/upload/v1782361603/wedding/pa9z7ymm20vla9zcnrgp.mp3",
          "style": {
            "margin": "1rem auto"
          }
        },
        {
          "id": "w-v-cover-decor-left",
          "type": "floating_decor",
          "content": "https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347246/wedding/q2qgf7ktfml1ff9zupkb.jpg",
          "style": {
            "position": "absolute",
            "top": "8%",
            "left": "-10px",
            "width": "80px",
            "height": "120px",
            "opacity": "0.85",
            "zIndex": "10"
          },
          "meta": { "animation": "sway", "speed": "6s", "delay": "0s" }
        },
        {
          "id": "w-v-cover-decor-right",
          "type": "floating_decor",
          "content": "https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347246/wedding/q2qgf7ktfml1ff9zupkb.jpg",
          "style": {
            "position": "absolute",
            "top": "15%",
            "right": "-10px",
            "width": "80px",
            "height": "120px",
            "opacity": "0.85",
            "zIndex": "10",
            "transform": "scaleX(-1)"
          },
          "meta": { "animation": "sway", "speed": "8s", "delay": "1s" }
        },
        {
          "id": "w-v-cover-icon",
          "type": "icon",
          "content": "heart",
          "style": {
            "margin": "0 auto 1.5rem",
            "color": "#8c6239",
            "textAlign": "center",
            "display": "block"
          },
          "meta": { "size": "36", "strokeWidth": "2", "color": "#8c6239" }
        },
        {
          "id": "w-v-cover-1",
          "type": "heading",
          "content": "THE MARRIAGE OF",
          "style": {
            "fontSize": "11px",
            "fontWeight": "600",
            "letterSpacing": "0.25em",
            "textAlign": "center",
            "color": "#ffffff"
          },
          "meta": { "level": "h4" }
        },
        {
          "id": "w-v-cover-2",
          "type": "heading",
          "content": "Carter & Evelyn",
          "style": {
            "fontFamily": "LatenzaScript_PERSONAL_USE_ONLY",
            "fontSize": "56px",
            "color": "#eab308",
            "textAlign": "center",
            "margin": "1rem 0"
          },
          "meta": { "level": "h1" }
        },
        {
          "id": "w-v-cover-3",
          "type": "text",
          "content": "12.09.2026",
          "style": {
            "fontSize": "15px",
            "fontFamily": "Playfair Display, serif",
            "textAlign": "center",
            "color": "#f5f5f4"
          }
        }
      ]
    },
    {
      "id": "sec-vintage-profile",
      "name": "Mempelai",
      "paddingTop": "5rem",
      "paddingBottom": "5rem",
      "paddingLeft": "2rem",
      "paddingRight": "2rem",
      "backgroundColor": "#fdfbf7",
      "transition": "fade-in",
      "textColor": "#44403c",
      "widgets": [
        {
          "id": "w-v-prof-decor-left",
          "type": "floating_decor",
          "content": "https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347246/wedding/q2qgf7ktfml1ff9zupkb.jpg",
          "style": {
            "position": "absolute",
            "top": "8%",
            "left": "-10px",
            "width": "80px",
            "height": "120px",
            "opacity": "0.85",
            "zIndex": "10"
          },
          "meta": { "animation": "drift", "speed": "7s", "delay": "0s" }
        },
        {
          "id": "w-v-prof-decor-right",
          "type": "floating_decor",
          "content": "https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347246/wedding/q2qgf7ktfml1ff9zupkb.jpg",
          "style": {
            "position": "absolute",
            "bottom": "8%",
            "right": "-10px",
            "width": "80px",
            "height": "120px",
            "opacity": "0.85",
            "zIndex": "10",
            "transform": "scaleX(-1)"
          },
          "meta": { "animation": "drift", "speed": "9s", "delay": "1s" }
        },
        {
          "id": "w-v-prof-1",
          "type": "heading",
          "content": "THE COUPLE",
          "style": {
            "fontSize": "12px",
            "fontWeight": "700",
            "letterSpacing": "0.2em",
            "textAlign": "center",
            "color": "#8c6239",
            "marginBottom": "2rem"
          },
          "meta": { "level": "h4" }
        },
        {
          "id": "w-v-prof-2",
          "type": "image",
          "content": "https://images.unsplash.com/photo-1500648767791-00dcc994a43e?auto=format&fit=crop&w=400&q=80",
          "style": {
            "width": "180px",
            "height": "240px",
            "margin": "0 auto 1rem auto",
            "objectFit": "cover"
          },
          "meta": {
            "frame": "polaroid",
            "caption": "The Groom"
          }
        },
        {
          "id": "w-v-prof-3",
          "type": "heading",
          "content": "Carter Henderson",
          "style": {
            "fontFamily": "Playfair Display, serif",
            "fontSize": "22px",
            "textAlign": "center"
          },
          "meta": { "level": "h3" }
        },
        {
          "id": "w-v-prof-4",
          "type": "text",
          "content": "Son of Thomas & Beatrice Henderson",
          "style": {
            "fontSize": "13px",
            "fontStyle": "italic",
            "textAlign": "center",
            "color": "#78716c",
            "marginBottom": "1rem"
          }
        },
        {
          "id": "w-v-groom-social",
          "type": "social_links",
          "content": "",
          "style": {
            "gap": "12px",
            "justifyContent": "center",
            "marginTop": "0.5rem",
            "marginBottom": "2.5rem"
          },
          "meta": {
            "instagram": "https://instagram.com/carter_henderson",
            "facebook": "https://facebook.com/carter_henderson",
            "size": "20",
            "color": "#8c6239"
          }
        },
        {
          "id": "w-v-prof-divider",
          "type": "icon",
          "content": "heart",
          "style": {
            "color": "#8c6239",
            "textAlign": "center",
            "margin": "1.5rem auto",
            "display": "block"
          },
          "meta": { "size": "28", "strokeWidth": "2", "color": "#8c6239" }
        },
        {
          "id": "w-v-prof-5",
          "type": "image",
          "content": "https://images.unsplash.com/photo-1438761681033-6461ffad8d80?auto=format&fit=crop&w=400&q=80",
          "style": {
            "width": "180px",
            "height": "240px",
            "margin": "1.5rem auto 1rem auto",
            "objectFit": "cover"
          },
          "meta": {
            "frame": "polaroid",
            "caption": "The Bride"
          }
        },
        {
          "id": "w-v-prof-6",
          "type": "heading",
          "content": "Evelyn Sterling",
          "style": {
            "fontFamily": "Playfair Display, serif",
            "fontSize": "22px",
            "textAlign": "center"
          },
          "meta": { "level": "h3" }
        },
        {
          "id": "w-v-prof-7",
          "type": "text",
          "content": "Daughter of Arthur & Penelope Sterling",
          "style": {
            "fontSize": "13px",
            "fontStyle": "italic",
            "textAlign": "center",
            "color": "#78716c",
            "marginBottom": "1rem"
          }
        },
        {
          "id": "w-v-bride-social",
          "type": "social_links",
          "content": "",
          "style": {
            "gap": "12px",
            "justifyContent": "center",
            "marginTop": "0.5rem",
            "marginBottom": "2rem"
          },
          "meta": {
            "instagram": "https://instagram.com/evelyn_sterling",
            "tiktok": "https://tiktok.com/@evelyn_sterling",
            "size": "20",
            "color": "#8c6239"
          }
        }
      ]
    },
    {
      "id": "sec-vintage-events",
      "name": "Wedding Schedule",
      "paddingTop": "5rem",
      "paddingBottom": "5rem",
      "paddingLeft": "2rem",
      "paddingRight": "2rem",
      "backgroundColor": "#f5f2eb",
      "transition": "zoom-in",
      "textColor": "#44403c",
      "widgets": [
        {
          "id": "w-v-evt-icon",
          "type": "icon",
          "content": "calendar",
          "style": {
            "color": "#8c6239",
            "textAlign": "center",
            "margin": "0 auto 1.5rem auto",
            "display": "block"
          },
          "meta": { "size": "32", "strokeWidth": "2", "color": "#8c6239" }
        },
        {
          "id": "w-v-evt-1",
          "type": "heading",
          "content": "THE VENUE",
          "style": {
            "fontSize": "12px",
            "fontWeight": "700",
            "letterSpacing": "0.2em",
            "textAlign": "center",
            "color": "#8c6239",
            "marginBottom": "2rem"
          },
          "meta": { "level": "h4" }
        },
        {
          "id": "w-v-evt-2",
          "type": "container",
          "content": "Wedding Ceremony\nSaturday, September 12, 2026\n3:00 PM - 4:30 PM\nSt. Jude Chapel, Boston",
          "style": {
            "backgroundColor": "#ffffff",
            "padding": "2rem",
            "borderRadius": "4px",
            "border": "1px solid #dcd7cc",
            "margin": "0 auto 1.5rem auto",
            "maxWidth": "360px",
            "textAlign": "center",
            "fontFamily": "Playfair Display, serif"
          }
        },
        {
          "id": "w-v-evt-3",
          "type": "container",
          "content": "Grand Reception\nSaturday, September 12, 2026\n6:00 PM - 10:00 PM\nHeritage Manor, Boston",
          "style": {
            "backgroundColor": "#ffffff",
            "padding": "2rem",
            "borderRadius": "4px",
            "border": "1px solid #dcd7cc",
            "margin": "0 auto 1.5rem auto",
            "maxWidth": "360px",
            "textAlign": "center",
            "fontFamily": "Playfair Display, serif"
          }
        }
      ]
    },
    {
      "id": "sec-vintage-gallery",
      "name": "Gallery",
      "paddingTop": "5rem",
      "paddingBottom": "5rem",
      "paddingLeft": "2rem",
      "paddingRight": "2rem",
      "backgroundColor": "#fdfbf7",
      "transition": "zoom-in-parallax",
      "textColor": "#44403c",
      "widgets": [
        {
          "id": "w-v-gal-1",
          "type": "heading",
          "content": "MEMORIES",
          "style": {
            "fontSize": "12px",
            "fontWeight": "700",
            "letterSpacing": "0.2em",
            "textAlign": "center",
            "color": "#8c6239",
            "marginBottom": "2rem"
          },
          "meta": { "level": "h4" }
        },
        {
          "id": "w-v-gal-2",
          "type": "gallery",
          "content": "",
          "style": {
            "margin": "1rem auto"
          },
          "meta": {
            "cols": 2,
            "images": [
              "https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347253/wedding/mxgeijfajqrr0loowcjo.jpg",
              "https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347246/wedding/q2qgf7ktfml1ff9zupkb.jpg",
              "https://images.unsplash.com/photo-1511795409834-ef04bbd61622?auto=format&fit=crop&w=400&q=80",
              "https://images.unsplash.com/photo-1519225495810-7512c696505a?auto=format&fit=crop&w=400&q=80"
            ]
          }
        }
      ]
    },
    {
      "id": "sec-vintage-rsvp",
      "name": "RSVP",
      "paddingTop": "5rem",
      "paddingBottom": "5rem",
      "paddingLeft": "2rem",
      "paddingRight": "2rem",
      "backgroundColor": "#f5f2eb",
      "transition": "fade-up",
      "textColor": "#44403c",
      "widgets": [
        {
          "id": "w-v-rsv-icon",
          "type": "icon",
          "content": "mail",
          "style": {
            "color": "#8c6239",
            "textAlign": "center",
            "margin": "0 auto 1.5rem auto",
            "display": "block"
          },
          "meta": { "size": "32", "strokeWidth": "2", "color": "#8c6239" }
        },
        {
          "id": "w-v-rsv-1",
          "type": "heading",
          "content": "RSVP",
          "style": {
            "fontSize": "12px",
            "fontWeight": "700",
            "letterSpacing": "0.2em",
            "textAlign": "center",
            "color": "#8c6239",
            "marginBottom": "2rem"
          },
          "meta": { "level": "h4" }
        },
        {
          "id": "w-v-rsv-2",
          "type": "rsvp",
          "content": "",
          "style": {
            "maxWidth": "480px",
            "margin": "0 auto"
          }
        }
      ]
    },
    {
      "id": "sec-vintage-gift",
      "name": "Gift Registry",
      "paddingTop": "5rem",
      "paddingBottom": "5rem",
      "paddingLeft": "2rem",
      "paddingRight": "2rem",
      "backgroundColor": "#fdfbf7",
      "transition": "fade-up",
      "textColor": "#44403c",
      "widgets": [
        {
          "id": "w-v-gft-icon",
          "type": "icon",
          "content": "gift",
          "style": {
            "color": "#8c6239",
            "textAlign": "center",
            "margin": "0 auto 1.5rem auto",
            "display": "block"
          },
          "meta": { "size": "32", "strokeWidth": "2", "color": "#8c6239" }
        },
        {
          "id": "w-v-gft-1",
          "type": "heading",
          "content": "GIFT REGISTRY",
          "style": {
            "fontSize": "12px",
            "fontWeight": "700",
            "letterSpacing": "0.2em",
            "textAlign": "center",
            "color": "#8c6239",
            "marginBottom": "2rem"
          },
          "meta": { "level": "h4" }
        },
        {
          "id": "w-v-gft-2",
          "type": "container",
          "content": "Heritage Bank\nAccount: 456-789-012\nBeneficiary: Evelyn Sterling",
          "style": {
            "backgroundColor": "#ffffff",
            "border": "1px solid #dcd7cc",
            "padding": "1.5rem",
            "borderRadius": "4px",
            "textAlign": "center",
            "maxWidth": "320px",
            "margin": "1rem auto",
            "fontFamily": "Playfair Display, serif"
          }
        }
      ]
    },
    {
      "id": "sec-vintage-footer",
      "name": "Footer",
      "paddingTop": "4rem",
      "paddingBottom": "4rem",
      "paddingLeft": "2rem",
      "paddingRight": "2rem",
      "backgroundColor": "#292524",
      "transition": "fade-in",
      "textColor": "#ffffff",
      "widgets": [
        {
          "id": "w-v-ftr-1",
          "type": "heading",
          "content": "Thank You",
          "style": {
            "color": "#eab308",
            "fontFamily": "LatenzaScript_PERSONAL_USE_ONLY",
            "fontSize": "44px",
            "textAlign": "center"
          },
          "meta": { "level": "h2" }
        },
        {
          "id": "w-v-ftr-2",
          "type": "text",
          "content": "Carter & Evelyn",
          "style": {
            "color": "#d6d3d1",
            "fontSize": "14px",
            "fontFamily": "Playfair Display, serif",
            "textAlign": "center",
            "marginTop": "0.5rem"
          }
        }
      ]
    }
  ]
}`

	renderHTML := `
<style>
  @import url('https://fonts.googleapis.com/css2?family=Playfair+Display:ital,wght@0,400..900;1,400..900&display=swap');
  @font-face {
    font-family: 'LatenzaScript_PERSONAL_USE_ONLY';
    src: url('https://res.cloudinary.com/dyuol7xfx/raw/upload/v1782361603/wedding/pa9z7ymm20vla9zcnrgp');
    font-weight: normal;
    font-style: normal;
    font-display: swap;
  }
  body { margin: 0; padding: 0; font-family: 'Playfair Display', serif; background-color: #fdfbf7; }
  * { box-sizing: border-box; }
  .list-disc { list-style-type: disc; }
  .list-decimal { list-style-type: decimal; }

  /* Animation System */
  @keyframes floatSway {
    0% { transform: translateY(0px) rotate(0deg); }
    50% { transform: translateY(-15px) rotate(6deg); }
    100% { transform: translateY(0px) rotate(0deg); }
  }
  @keyframes floatSwaySlow {
    0% { transform: translateY(0px) rotate(0deg); }
    50% { transform: translateY(-25px) rotate(-10deg); }
    100% { transform: translateY(0px) rotate(0deg); }
  }
  @keyframes spinSlow {
    0% { transform: rotate(0deg); }
    100% { transform: rotate(360deg); }
  }
  @keyframes pulseSlow {
    0%, 100% { transform: scale(1); }
    50% { transform: scale(1.12); }
  }
  @keyframes driftLeftRight {
    0%, 100% { transform: translateX(0px) rotate(0deg); }
    50% { transform: translateX(20px) rotate(8deg); }
    100% { transform: translateX(0px) rotate(0deg); }
  }
  .animate-float-sway { animation: floatSway var(--float-speed, 6s) ease-in-out infinite var(--float-delay, 0s); }
  .animate-float-sway-slow { animation: floatSwaySlow var(--float-speed, 10s) ease-in-out infinite var(--float-delay, 0s); }
  .animate-spin-slow { animation: spinSlow var(--float-speed, 15s) linear infinite var(--float-delay, 0s); }
  .animate-pulse-slow { animation: pulseSlow var(--float-speed, 5s) ease-in-out infinite var(--float-delay, 0s); }
  .animate-drift-lr { animation: driftLeftRight var(--float-speed, 8s) ease-in-out infinite var(--float-delay, 0s); }

  @keyframes fadeUp {
    from { opacity: 0; transform: translateY(30px); }
    to { opacity: 1; transform: translateY(0); }
  }
  @keyframes fadeIn {
    from { opacity: 0; }
    to { opacity: 1; }
  }
  @keyframes slideLeft {
    from { opacity: 0; transform: translateX(-30px); }
    to { opacity: 1; transform: translateX(0); }
  }
  @keyframes slideRight {
    from { opacity: 0; transform: translateX(30px); }
    to { opacity: 1; transform: translateX(0); }
  }
  @keyframes zoomIn {
    from { opacity: 0; transform: scale(0.95); }
    to { opacity: 1; transform: scale(1); }
  }

  .animate-widget {
    opacity: 0;
    transition: opacity 0.8s cubic-bezier(0.16, 1, 0.3, 1), transform 0.8s cubic-bezier(0.16, 1, 0.3, 1);
  }

  .is-in-view[data-transition="fade-up"] .animate-widget,
  .is-in-view[data-transition="fade-up-parallax"] .animate-widget {
    animation: fadeUp 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
  }
  .is-in-view[data-transition="fade-in"] .animate-widget {
    animation: fadeIn 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
  }
  .is-in-view[data-transition="slide-left"] .animate-widget {
    animation: slideLeft 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
  }
  .is-in-view[data-transition="slide-right"] .animate-widget {
    animation: slideRight 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
  }
  .is-in-view[data-transition="zoom-in"] .animate-widget,
  .is-in-view[data-transition="zoom-in-parallax"] .animate-widget {
    animation: zoomIn 0.8s cubic-bezier(0.16, 1, 0.3, 1) forwards;
  }

  /* Parallax Background Layout */
  .parallax-section {
    position: relative;
    overflow: hidden;
  }
  .parallax-bg {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 125%; /* Taller for translation bounds */
    z-index: 0;
    pointer-events: none;
    background-size: cover;
    background-position: center;
    background-repeat: no-repeat;
    will-change: transform;
  }

  @media (prefers-reduced-motion: reduce) {
    .animate-widget {
      opacity: 1 !important;
      transform: none !important;
      animation: none !important;
      transition: none !important;
    }
    .parallax-bg {
      transform: none !important;
      height: 100% !important;
    }
  }
</style>
<div style="font-family: 'Playfair Display', serif; background-color: #fdfbf7; min-height: 100vh; width: 100%;">
  <div style="max-width: 100%; margin: 0 auto; min-height: 100vh; box-shadow: 0 10px 50px rgba(0,0,0,0.04); background: #ffffff;">

    <!-- Cover Section -->
    <section id="sec-vintage-cover" class="parallax-section" data-transition="fade-up-parallax" style="position: relative; padding: 6rem 2rem; text-align: center; min-height: 100vh; color: #ffffff; display: flex; flex-direction: column; align-items: center; justify-content: center; overflow: hidden; background-color: #fdfbf7;">
      <div class="parallax-bg" style="background-image: url('https://images.unsplash.com/photo-1465495976277-4387d4b0b4c6?auto=format&fit=crop&w=1200&q=80'); background-size: cover; background-position: center; transform: translateY(0px) scale(1.15);"></div>
      
      <!-- Embedded YouTube Video Loop Background -->
      <iframe src="https://www.youtube.com/embed/F3x9T8S_BPE?autoplay=1&mute=1&loop=1&playlist=F3x9T8S_BPE&controls=0&showinfo=0&rel=0&iv_load_policy=3&playsinline=1&enablejsapi=1" style="position: absolute; top:0; left:0; width:100%; height:100%; border:none; object-fit: cover; opacity: 0.22; pointer-events:none; z-index:0; filter: sepia(0.6) contrast(0.95) brightness(0.65);" allow="autoplay; encrypted-media"></iframe>

      <div style="position: absolute; top: 0; left: 0; width: 100%; height: 100%; background-color: #292524; opacity: 0.55; z-index: 1; pointer-events: none;"></div>
      
      <!-- Floating Leaves on Cover -->
      <div class="animate-float-sway" style="position: absolute; top: 8%; left: -10px; width: 80px; height: 120px; opacity: 0.85; z-index: 10; pointer-events: none; --float-speed: 6s; --float-delay: 0s;">
        <img src="https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347246/wedding/q2qgf7ktfml1ff9zupkb.jpg" style="width: 100%; height: 100%; object-fit: contain; filter: sepia(0.2) drop-shadow(0 4px 8px rgba(0,0,0,0.15));" />
      </div>
      <div class="animate-float-sway-slow" style="position: absolute; top: 18%; right: -15px; width: 80px; height: 120px; opacity: 0.85; z-index: 10; pointer-events: none; transform: scaleX(-1); --float-speed: 8s; --float-delay: 1s;">
        <img src="https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347246/wedding/q2qgf7ktfml1ff9zupkb.jpg" style="width: 100%; height: 100%; object-fit: contain; filter: sepia(0.2) drop-shadow(0 4px 8px rgba(0,0,0,0.15));" />
      </div>

      <div style="position: relative; z-index: 2; width: 100%;">
        <!-- Audio / Music Player Widget -->
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <div style="background: rgba(253, 251, 247, 0.88); border: 1px solid rgba(140, 98, 57, 0.35); padding: 0.65rem 1.2rem; border-radius: 99px; display: inline-flex; align-items: center; gap: 0.65rem; margin: 0 auto 1.5rem auto; box-shadow: 0 4px 15px rgba(0,0,0,0.06); backdrop-filter: blur(4px);">
            <div style="width: 26px; height: 26px; border-radius: 50%; background: #8c6239; display: flex; align-items: center; justify-content: center; color: white; font-size: 10px; font-weight: bold; cursor: pointer;">▶</div>
            <span style="font-size: 11px; color: #8c6239; font-weight: 700; font-family: monospace; letter-spacing: 0.05em; text-transform: uppercase;">Vintage Melody</span>
          </div>
        </div>

        <!-- Heart Icon -->
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%; display: flex; justify-content: center; margin: 0 auto 1.5rem auto;">
          <svg viewBox="0 0 24 24" width="36" height="36" stroke="#8c6239" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/></svg>
        </div>

        <div class="animate-widget" style="animation-delay: 0.12s; transition-delay: 0.12s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <h4 style="color: #ffffff; font-size: 11px; font-weight: 600; letter-spacing: 0.25em; text-align: center; text-transform: uppercase; margin-top: 0;">THE MARRIAGE OF</h4>
        </div>
        <div class="animate-widget" style="animation-delay: 0.24s; transition-delay: 0.24s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <h1 style="color: #eab308; font-family: 'LatenzaScript_PERSONAL_USE_ONLY'; font-size: 56px; text-align: center; margin: 1rem 0;">Carter & Evelyn</h1>
        </div>
        <div class="animate-widget" style="animation-delay: 0.36s; transition-delay: 0.36s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <p style="color: #f5f5f4; font-size: 15px; text-align: center;">12.09.2026</p>
        </div>
      </div>
    </section>

    <!-- Mempelai Section -->
    <section id="sec-vintage-profile" data-transition="fade-in" style="padding: 5rem 2rem; background-color: #fdfbf7; text-align: center; color: #44403c; position: relative; overflow: hidden;">
      <!-- Floating Leaves on Profile -->
      <div class="animate-drift-lr" style="position: absolute; top: 10%; left: -15px; width: 80px; height: 120px; opacity: 0.7; z-index: 10; pointer-events: none; --float-speed: 7s; --float-delay: 0.5s;">
        <img src="https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347246/wedding/q2qgf7ktfml1ff9zupkb.jpg" style="width: 100%; height: 100%; object-fit: contain; filter: sepia(0.2);" />
      </div>
      <div class="animate-drift-lr" style="position: absolute; bottom: 15%; right: -15px; width: 80px; height: 120px; opacity: 0.7; z-index: 10; pointer-events: none; transform: scaleX(-1); --float-speed: 9s; --float-delay: 1.2s;">
        <img src="https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347246/wedding/q2qgf7ktfml1ff9zupkb.jpg" style="width: 100%; height: 100%; object-fit: contain; filter: sepia(0.2);" />
      </div>

      <div style="position: relative; z-index: 2; width: 100%;">
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <h4 style="color: #8c6239; font-size: 12px; font-weight: 700; letter-spacing: 0.2em; text-align: center; margin-bottom: 2rem;">THE COUPLE</h4>
        </div>
        
        <!-- Groom framed portrait (Polaroid Style) -->
        <div class="animate-widget" style="animation-delay: 0.12s; transition-delay: 0.12s; width: 100%;">
          <div style="background: #ffffff; padding: 12px 12px 28px 12px; border: 1.5px solid #dcd7cc; box-shadow: 0 10px 25px rgba(0,0,0,0.06); transform: rotate(-2.5deg); display: inline-block; margin: 1.5rem auto 1.5rem; max-width: 190px; width: 100%; text-align: center;">
            <img src="https://images.unsplash.com/photo-1500648767791-00dcc994a43e?auto=format&fit=crop&w=400&q=80" style="display: block; width: 100%; height: 210px; object-fit: cover; filter: sepia(0.15);" alt="Groom" />
            <p style="font-family: 'LatenzaScript_PERSONAL_USE_ONLY', cursive; font-size: 22px; margin: 12px 0 0 0; color: #8c6239;">The Groom</p>
          </div>
        </div>

        <div class="animate-widget" style="animation-delay: 0.24s; transition-delay: 0.24s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <h3 style="font-family: 'Playfair Display', serif; font-size: 22px; text-align: center; margin-top: 0;">Carter Henderson</h3>
        </div>
        <div class="animate-widget" style="animation-delay: 0.36s; transition-delay: 0.36s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <p style="font-size: 13px; font-style: italic; color: #78716c; text-align: center; margin-bottom: 0.75rem;">Son of Thomas & Beatrice Henderson</p>
        </div>

        <!-- Groom Social Links -->
        <div class="animate-widget" style="animation-delay: 0.42s; transition-delay: 0.42s; width: 100%;">
          <div style="display: flex; gap: 12px; justify-content: center; margin-bottom: 2.5rem; margin-top: 0.5rem;">
            <a href="https://instagram.com/carter_henderson" target="_blank" rel="noopener noreferrer" style="display: inline-flex; align-items: center; justify-content: center; width: 36px; height: 36px; border-radius: 50%; border: 1.5px solid #8c6239; color: #8c6239; text-decoration: none; transition: all 0.3s ease;">
              <svg viewBox="0 0 24 24" width="20" height="20" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><rect x="2" y="2" width="20" height="20" rx="5" ry="5"></rect><path d="M16 11.37A4 4 0 1 1 12.63 8 4 4 0 0 1 16 11.37z"></path><line x1="17.5" y1="6.5" x2="17.51" y2="6.5"></line></svg>
            </a>
            <a href="https://facebook.com/carter_henderson" target="_blank" rel="noopener noreferrer" style="display: inline-flex; align-items: center; justify-content: center; width: 36px; height: 36px; border-radius: 50%; border: 1.5px solid #8c6239; color: #8c6239; text-decoration: none; transition: all 0.3s ease;">
              <svg viewBox="0 0 24 24" width="20" height="20" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><path d="M18 2h-3a5 5 0 0 0-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 0 1 1-1h3z"></path></svg>
            </a>
          </div>
        </div>

        <!-- Heart Separator Icon -->
        <div class="animate-widget" style="animation-delay: 0.48s; transition-delay: 0.48s; width: 100%; display: flex; justify-content: center; margin: 1.5rem 0;">
          <svg viewBox="0 0 24 24" width="28" height="28" stroke="#8c6239" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/></svg>
        </div>
        
        <!-- Bride framed portrait (Polaroid Style) -->
        <div class="animate-widget" style="animation-delay: 0.54s; transition-delay: 0.54s; width: 100%;">
          <div style="background: #ffffff; padding: 12px 12px 28px 12px; border: 1.5px solid #dcd7cc; box-shadow: 0 10px 25px rgba(0,0,0,0.06); transform: rotate(2.5deg); display: inline-block; margin: 1.5rem auto 1.5rem; max-width: 190px; width: 100%; text-align: center;">
            <img src="https://images.unsplash.com/photo-1438761681033-6461ffad8d80?auto=format&fit=crop&w=400&q=80" style="display: block; width: 100%; height: 210px; object-fit: cover; filter: sepia(0.15);" alt="Bride" />
            <p style="font-family: 'LatenzaScript_PERSONAL_USE_ONLY', cursive; font-size: 22px; margin: 12px 0 0 0; color: #8c6239;">The Bride</p>
          </div>
        </div>

        <div class="animate-widget" style="animation-delay: 0.6s; transition-delay: 0.6s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <h3 style="font-family: 'Playfair Display', serif; font-size: 22px; text-align: center; margin-top: 0;">Evelyn Sterling</h3>
        </div>
        <div class="animate-widget" style="animation-delay: 0.72s; transition-delay: 0.72s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <p style="font-size: 13px; font-style: italic; color: #78716c; text-align: center; margin-bottom: 0.75rem;">Daughter of Arthur & Penelope Sterling</p>
        </div>

        <!-- Bride Social Links -->
        <div class="animate-widget" style="animation-delay: 0.78s; transition-delay: 0.78s; width: 100%;">
          <div style="display: flex; gap: 12px; justify-content: center; margin-bottom: 2rem; margin-top: 0.5rem;">
            <a href="https://instagram.com/evelyn_sterling" target="_blank" rel="noopener noreferrer" style="display: inline-flex; align-items: center; justify-content: center; width: 36px; height: 36px; border-radius: 50%; border: 1.5px solid #8c6239; color: #8c6239; text-decoration: none; transition: all 0.3s ease;">
              <svg viewBox="0 0 24 24" width="20" height="20" stroke="currentColor" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><rect x="2" y="2" width="20" height="20" rx="5" ry="5"></rect><path d="M16 11.37A4 4 0 1 1 12.63 8 4 4 0 0 1 16 11.37z"></path><line x1="17.5" y1="6.5" x2="17.51" y2="6.5"></line></svg>
            </a>
            <a href="https://tiktok.com/@evelyn_sterling" target="_blank" rel="noopener noreferrer" style="display: inline-flex; align-items: center; justify-content: center; width: 36px; height: 36px; border-radius: 50%; border: 1.5px solid #8c6239; color: #8c6239; text-decoration: none; transition: all 0.3s ease;">
              <svg viewBox="0 0 24 24" width="18" height="18" fill="currentColor"><path d="M12.525.02c1.31-.02 2.61-.01 3.91-.02.08 1.53.63 3.02 1.62 4.2-1.2.19-2.39.69-3.38 1.48-.05-1.89-.01-3.78-.03-5.67-.85-.02-1.72-.01-2.58-.02V13.51c.04 1.54-.51 3.05-1.57 4.19-1.34 1.41-3.41 2.05-5.32 1.66-2.22-.44-3.95-2.37-4.14-4.63-.33-3.21 2.37-6.07 5.58-5.91.06 1.76.01 3.53.03 5.29-.86-.09-1.76.24-2.3.96-.64.81-.59 2.05.15 2.76.92.83 2.47.74 3.25-.26.47-.64.53-1.46.52-2.23V.02z"/></svg>
            </a>
          </div>
        </div>
      </div>
    </section>

    <!-- Schedule Section -->
    <section id="sec-vintage-events" data-transition="zoom-in" style="padding: 5rem 2rem; background-color: #f5f2eb; text-align: center; color: #44403c; position: relative;">
      <div style="position: relative; z-index: 2; width: 100%;">
        <!-- Calendar Icon -->
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%; display: flex; justify-content: center; margin: 0 auto 1.5rem auto;">
          <svg viewBox="0 0 24 24" width="32" height="32" stroke="#8c6239" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="4" width="18" height="18" rx="2" ry="2"></rect><line x1="16" y1="2" x2="16" y2="6"></line><line x1="8" y1="2" x2="8" y2="6"></line><line x1="3" y1="10" x2="21" y2="10"></line></svg>
        </div>

        <div class="animate-widget" style="animation-delay: 0.05s; transition-delay: 0.05s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <h4 style="color: #8c6239; font-size: 12px; font-weight: 700; letter-spacing: 0.25em; text-align: center; margin-bottom: 2rem; text-transform: uppercase;">THE VENUE</h4>
        </div>

        <div style="display: flex; flex-direction: column; gap: 2rem; max-width: 420px; margin: 0 auto;">
          <!-- Ceremony Card (Diary page style) -->
          <div class="animate-widget" style="animation-delay: 0.12s; transition-delay: 0.12s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
            <div style="background-color: #faf6eb; padding: 2rem; border-radius: 4px; border: 1.5px solid #dcd7cc; box-shadow: 0 10px 25px rgba(0,0,0,0.03); text-align: left; position: relative;">
              <div style="position: absolute; left: 1rem; top: 0; bottom: 0; width: 1px; border-left: 1px dashed rgba(140, 98, 57, 0.25);"></div>
              <div style="padding-left: 1.5rem;">
                <h3 style="font-family: 'Playfair Display', serif; color: #8c6239; margin: 0 0 0.5rem 0; font-size: 22px; font-style: italic; font-weight: 700;">Wedding Ceremony</h3>
                <p style="margin: 0.5rem 0; font-size: 14px; font-weight: 600; color: #44403c;">Saturday, September 12, 2026</p>
                <p style="margin: 0.25rem 0; font-size: 13.5px; color: #8c6239; font-weight: 700;">3:00 PM - 4:30 PM</p>
                <div style="width: 40px; height: 1px; background: #dcd7cc; margin: 1rem 0;"></div>
                <p style="font-size: 13px; color: #78716c; line-height: 1.6; margin: 0;"><b>St. Jude Chapel</b><br>104 Chapel Hill Road, Boston, MA</p>
              </div>
            </div>
          </div>

          <!-- Reception Card -->
          <div class="animate-widget" style="animation-delay: 0.24s; transition-delay: 0.24s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
            <div style="background-color: #faf6eb; padding: 2rem; border-radius: 4px; border: 1.5px solid #dcd7cc; box-shadow: 0 10px 25px rgba(0,0,0,0.03); text-align: left; position: relative;">
              <div style="position: absolute; left: 1rem; top: 0; bottom: 0; width: 1px; border-left: 1px dashed rgba(140, 98, 57, 0.25);"></div>
              <div style="padding-left: 1.5rem;">
                <h3 style="font-family: 'Playfair Display', serif; color: #8c6239; margin: 0 0 0.5rem 0; font-size: 22px; font-style: italic; font-weight: 700;">Grand Reception</h3>
                <p style="margin: 0.5rem 0; font-size: 14px; font-weight: 600; color: #44403c;">Saturday, September 12, 2026</p>
                <p style="margin: 0.25rem 0; font-size: 13.5px; color: #8c6239; font-weight: 700;">6:00 PM - 10:00 PM</p>
                <div style="width: 40px; height: 1px; background: #dcd7cc; margin: 1rem 0;"></div>
                <p style="font-size: 13px; color: #78716c; line-height: 1.6; margin: 0;"><b>Heritage Manor</b><br>405 Historic Way, Boston, MA</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Gallery Section -->
    <section id="sec-vintage-gallery" class="parallax-section" data-transition="zoom-in-parallax" style="padding: 5rem 2rem; background-color: #fdfbf7; text-align: center; color: #44403c; position: relative;">
      <div style="position: relative; z-index: 2; width: 100%;">
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <h4 style="color: #8c6239; font-size: 12px; font-weight: 700; letter-spacing: 0.25em; text-align: center; margin-bottom: 2rem; text-transform: uppercase;">MEMORIES</h4>
        </div>
        
        <!-- Polaroid Scrapbook style -->
        <div class="animate-widget" style="animation-delay: 0.12s; transition-delay: 0.12s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <div style="display: grid; grid-template-columns: repeat(2, 1fr); gap: 1.5rem; margin: 2rem auto; max-width: 480px; width: 100%;">
            <!-- Polaroid 1 -->
            <div style="background: #ffffff; padding: 12px 12px 30px 12px; border: 1px solid #dcd7cc; box-shadow: 0 8px 20px rgba(0,0,0,0.06); transform: rotate(-2.5deg); transition: transform 0.3s; width: 100%;" onmouseover="this.style.transform='rotate(0deg) scale(1.03)'" onmouseout="this.style.transform='rotate(-2.5deg)'">
              <img src="https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347253/wedding/mxgeijfajqrr0loowcjo.jpg" style="width: 100%; height: 140px; object-fit: cover; filter: sepia(0.2);" alt="Gallery 1" />
              <p style="font-family: 'LatenzaScript_PERSONAL_USE_ONLY'; font-size: 20px; margin: 12px 0 0 0; color: #8c6239; text-align: center;">Our Sweet Days</p>
            </div>
            <!-- Polaroid 2 -->
            <div style="background: #ffffff; padding: 12px 12px 30px 12px; border: 1px solid #dcd7cc; box-shadow: 0 8px 20px rgba(0,0,0,0.06); transform: rotate(2.5deg); transition: transform 0.3s; width: 100%;" onmouseover="this.style.transform='rotate(0deg) scale(1.03)'" onmouseout="this.style.transform='rotate(2.5deg)'">
              <img src="https://res.cloudinary.com/dyuol7xfx/image/upload/v1782347246/wedding/q2qgf7ktfml1ff9zupkb.jpg" style="width: 100%; height: 140px; object-fit: cover; filter: sepia(0.2);" alt="Gallery 2" />
              <p style="font-family: 'LatenzaScript_PERSONAL_USE_ONLY'; font-size: 20px; margin: 12px 0 0 0; color: #8c6239; text-align: center;">Autumn Walk</p>
            </div>
            <!-- Polaroid 3 -->
            <div style="background: #ffffff; padding: 12px 12px 30px 12px; border: 1px solid #dcd7cc; box-shadow: 0 8px 20px rgba(0,0,0,0.06); transform: rotate(2deg); transition: transform 0.3s; width: 100%;" onmouseover="this.style.transform='rotate(0deg) scale(1.03)'" onmouseout="this.style.transform='rotate(2deg)'">
              <img src="https://images.unsplash.com/photo-1511795409834-ef04bbd61622?auto=format&fit=crop&w=400&q=80" style="width: 100%; height: 140px; object-fit: cover; filter: sepia(0.2);" alt="Gallery 3" />
              <p style="font-family: 'LatenzaScript_PERSONAL_USE_ONLY'; font-size: 20px; margin: 12px 0 0 0; color: #8c6239; text-align: center;">Precious Moments</p>
            </div>
            <!-- Polaroid 4 -->
            <div style="background: #ffffff; padding: 12px 12px 30px 12px; border: 1px solid #dcd7cc; box-shadow: 0 8px 20px rgba(0,0,0,0.06); transform: rotate(-2deg); transition: transform 0.3s; width: 100%;" onmouseover="this.style.transform='rotate(0deg) scale(1.03)'" onmouseout="this.style.transform='rotate(-2deg)'">
              <img src="https://images.unsplash.com/photo-1519225495810-7512c696505a?auto=format&fit=crop&w=400&q=80" style="width: 100%; height: 140px; object-fit: cover; filter: sepia(0.2);" alt="Gallery 4" />
              <p style="font-family: 'LatenzaScript_PERSONAL_USE_ONLY'; font-size: 20px; margin: 12px 0 0 0; color: #8c6239; text-align: center;">Forever Began</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- RSVP Section -->
    <section id="sec-vintage-rsvp" data-transition="fade-up" style="padding: 5rem 2rem; background-color: #f5f2eb; text-align: center; color: #44403c; position: relative;">
      <div style="position: relative; z-index: 2; width: 100%;">
        <!-- Mail Icon -->
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%; display: flex; justify-content: center; margin: 0 auto 1.5rem auto;">
          <svg viewBox="0 0 24 24" width="32" height="32" stroke="#8c6239" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4h16c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H4c-1.1 0-2-.9-2-2V6c0-1.1.9-2 2-2z"></path><polyline points="22,6 12,13 2,6"></polyline></svg>
        </div>

        <div class="animate-widget" style="animation-delay: 0.05s; transition-delay: 0.05s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <h4 style="color: #8c6239; font-size: 12px; font-weight: 700; letter-spacing: 0.25em; text-align: center; margin-bottom: 2rem; text-transform: uppercase;">RSVP</h4>
        </div>
        
        <div class="animate-widget" style="animation-delay: 0.12s; transition-delay: 0.12s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <!-- Postcard Style RSVP Box -->
          <div style="background: #faf6eb; padding: 2.5rem 2rem; border-radius: 4px; max-width: 480px; margin: 0 auto; text-align: left; box-shadow: 0 10px 30px rgba(140, 98, 57, 0.05); border: 1.5px dashed #8c6239;">
            <form onsubmit="event.preventDefault(); alert('Thank you for RSVPing!')">
              <div style="margin-bottom: 2rem;">
                <label style="display: block; font-size: 0.7rem; font-weight: bold; margin-bottom: 0.25rem; color: #8c6239; text-transform: uppercase; font-family: monospace; letter-spacing: 0.1em;">Kindly Respond By Writing Your Name</label>
                <input type="text" placeholder="Your name..." required style="width: 100%; border: none; border-bottom: 1.5px dashed #8c6239; outline: none; background: transparent; font-family: 'LatenzaScript_PERSONAL_USE_ONLY'; font-size: 24px; color: #8c6239; padding: 6px 0;" />
              </div>
              
              <div style="margin-bottom: 2rem;">
                <label style="display: block; font-size: 0.7rem; font-weight: bold; margin-bottom: 0.25rem; color: #8c6239; text-transform: uppercase; font-family: monospace; letter-spacing: 0.1em;">Send Us Warm Wishes</label>
                <textarea rows="3" placeholder="Write your letter..." required style="width: 100%; border: none; border-bottom: 1.5px dashed #8c6239; outline: none; background: transparent; font-family: 'LatenzaScript_PERSONAL_USE_ONLY'; font-size: 22px; color: #8c6239; padding: 6px 0; resize: none;"></textarea>
              </div>
              
              <button type="submit" style="width: 100%; height: 46px; border-radius: 4px; background: #8c6239; color: #ffffff; font-weight: bold; border: 0; cursor: pointer; text-transform: uppercase; font-size: 12px; letter-spacing: 0.1em; font-family: monospace; box-shadow: 0 4px 12px rgba(140,98,57,0.15);">Send Postcard</button>
            </form>
          </div>
        </div>
      </div>
    </section>

    <!-- Gift Section -->
    <section id="sec-vintage-gift" data-transition="fade-up" style="padding: 5rem 2rem; background-color: #fdfbf7; text-align: center; color: #44403c; position: relative;">
      <div style="position: relative; z-index: 2; width: 100%;">
        <!-- Gift Icon -->
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%; display: flex; justify-content: center; margin: 0 auto 1.5rem auto;">
          <svg viewBox="0 0 24 24" width="32" height="32" stroke="#8c6239" stroke-width="2" fill="none" stroke-linecap="round" stroke-linejoin="round"><polyline points="20 12 20 22 4 22 4 12"></polyline><rect x="2" y="7" width="20" height="5"></rect><line x1="12" y1="22" x2="12" y2="7"></line><path d="M12 7H7.5a2.5 2.5 0 0 1 0-5C11 2 12 7 12 7z"></path><path d="M12 7h4.5a2.5 2.5 0 0 0 0-5C13 2 12 7 12 7z"></path></svg>
        </div>

        <div class="animate-widget" style="animation-delay: 0.05s; transition-delay: 0.05s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <h4 style="color: #8c6239; font-size: 12px; font-weight: 700; letter-spacing: 0.25em; text-align: center; margin-bottom: 2rem; text-transform: uppercase;">GIFT REGISTRY</h4>
        </div>
        
        <div class="animate-widget" style="animation-delay: 0.12s; transition-delay: 0.12s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <!-- Stamp/Parcel Style Box -->
          <div style="background-color: #faf6eb; border: 1.5px solid #dcd7cc; padding: 1.75rem; border-radius: 4px; text-align: center; max-width: 320px; margin: 1.5rem auto; font-family: 'Playfair Display', serif; position: relative; box-shadow: 0 6px 20px rgba(0,0,0,0.02);">
            <!-- Postal Stamp simulation -->
            <div style="position: absolute; right: 12px; top: 12px; width: 35px; height: 45px; border: 2px dashed #8c6239; background: #fff; font-size: 8px; color: #8c6239; display: flex; flex-direction: column; align-items: center; justify-content: center; font-family: monospace; font-weight: bold; transform: rotate(5deg);">
              <span>POST</span>
              <span>12c</span>
            </div>
            
            <p style="margin: 0; font-weight: bold; font-size: 13px; color: #8c6239; text-transform: uppercase; font-family: monospace; letter-spacing: 0.05em; text-align: left;">HERITAGE BANK</p>
            <p style="margin: 1rem 0 0.5rem 0; font-size: 20px; color: #8c6239; font-weight: 800; font-family: monospace; text-align: left; letter-spacing: 0.05em;">456-789-012</p>
            <p style="margin: 0 0 1rem 0; font-size: 12px; color: #78716c; text-align: left; font-style: italic;">a.n Evelyn Sterling</p>
            <button onclick="navigator.clipboard.writeText('456-789-012'); const btn = this; const orig = btn.innerText; btn.innerText = 'COPIED!'; setTimeout(() => btn.innerText = orig, 2000);" style="width: 100%; padding: 6px; border-radius: 4px; border: 1px solid #8c6239; background: transparent; color: #8c6239; font-size: 11px; font-weight: bold; font-family: monospace; cursor: pointer; text-transform: uppercase; transition: all 0.2s;">Copy Number</button>
          </div>
        </div>
      </div>
    </section>

    <!-- Footer Section -->
    <section id="sec-vintage-footer" data-transition="fade-in" style="padding: 5rem 2rem; background-color: #2b2520; text-align: center; color: #ffffff; position: relative; border-top: 1px solid #3d352e;">
      <div style="position: relative; z-index: 2; width: 100%;">
        <!-- Leaf design flourish -->
        <div style="margin-bottom: 1.5rem; opacity: 0.4;">
          <svg width="60" height="20" viewBox="0 0 60 20" fill="none" xmlns="http://www.w3.org/2000/svg" style="display: block; margin: 0 auto;">
            <path d="M5 10C15 10 20 5 30 10C40 15 45 10 55 10" stroke="#8c6239" stroke-width="1.5" stroke-linecap="round"/>
            <circle cx="30" cy="10" r="3" fill="#8c6239"/>
          </svg>
        </div>
        <div class="animate-widget" style="animation-delay: 0s; transition-delay: 0s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <h2 style="color: #8c6239; font-family: 'LatenzaScript_PERSONAL_USE_ONLY'; font-size: 44px; text-align: center; margin-top: 0; margin-bottom: 0;">Thank You</h2>
        </div>
        <div class="animate-widget" style="animation-delay: 0.12s; transition-delay: 0.12s; width: 100%; display: flex; flex-direction: column; align-items: inherit;">
          <p style="color: #d6d3d1; font-size: 14px; font-family: 'Playfair Display', serif; text-align: center; margin-top: 1.5rem; letter-spacing: 0.05em; font-style: italic;">Carter & Evelyn</p>
        </div>
      </div>
    </section>

  </div>
</div>
</div>
`

	_, err := pool.Exec(ctx,
		`INSERT INTO themes (name, slug, description, thumbnail, theme_data, render_html, created_at, updated_at) 
		 VALUES ($1, $2, $3, $4, $5, $6, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`,
		name, slug, description, thumbnail, themeDataJSON, renderHTML,
	)
	if err != nil {
		log.Printf("Warning: failed to seed vintage theme: %v", err)
		return
	}
	log.Println("Seeded professional dummy theme 'Vintage Romance' successfully!")
}


func seedDummyContexts(pool *pgxpool.Pool) {
	ctx := context.Background()

	demos := []struct {
		contextName string
		contextSlug string
		themeSlug   string
	}{
		{
			contextName: "Elegant Demo",
			contextSlug: "elegant-demo",
			themeSlug:   "royal-gold",
		},
		{
			contextName: "Modern Demo",
			contextSlug: "modern-demo",
			themeSlug:   "modern-sinis",
		},
		{
			contextName: "Vintage Demo",
			contextSlug: "vintage-demo",
			themeSlug:   "vintage-romance",
		},
	}

	for _, demo := range demos {
		// Always clear existing demo context to refresh its render_html and data from the seeded theme
		_, _ = pool.Exec(ctx, "DELETE FROM guests WHERE context_id IN (SELECT id FROM contexts WHERE slug = $1)", demo.contextSlug)
		_, _ = pool.Exec(ctx, "DELETE FROM contexts WHERE slug = $1", demo.contextSlug)

		// Context does not exist, look up theme_id and details
		var contextID int
		var themeID int
		var themeDataJSON string
		var renderHTML string
		err := pool.QueryRow(ctx, "SELECT id, theme_data, render_html FROM themes WHERE slug = $1", demo.themeSlug).Scan(&themeID, &themeDataJSON, &renderHTML)
		if err != nil {
			log.Printf("Warning: failed to find theme %s for seeding context %s: %v", demo.themeSlug, demo.contextSlug, err)
			continue
		}

		// Insert context
		err = pool.QueryRow(ctx,
			`INSERT INTO contexts (name, slug, theme_id, render_html, content_json, created_at, updated_at)
			 VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id`,
			demo.contextName, demo.contextSlug, themeID, renderHTML, themeDataJSON,
		).Scan(&contextID)
		if err != nil {
			log.Printf("Warning: failed to seed context %s: %v", demo.contextSlug, err)
			continue
		}
		log.Printf("Seeded context %s successfully!", demo.contextSlug)

		// Insert guest
		_, err = pool.Exec(ctx,
			"INSERT INTO guests (context_id, name, slug, created_at, updated_at) VALUES ($1, 'Guest Demo', 'guest-demo', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)",
			contextID,
		)
		if err != nil {
			log.Printf("Warning: failed to seed guest-demo for context %s: %v", demo.contextSlug, err)
		} else {
			log.Printf("Seeded guest-demo for context %s successfully!", demo.contextSlug)
		}
	}
}

