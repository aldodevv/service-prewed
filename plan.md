# service-wedding — Refactoring & Upgrade Plan

## ROLE

Kamu adalah **Principal Software Engineer & Software Architect**.

Refactor dan upgrade **service-wedding** (Go backend) agar production-ready, zero hardcode, zero mock, dan memiliki fitur lengkap sesuai kebutuhan Wedding Invitation Platform.

---

## CURRENT STATE ANALYSIS

### Architecture (Sudah Baik)
- Clean Architecture: `domain` → `repository` → `usecase` → `delivery/http`
- PostgreSQL + JSONB untuk flexible schema
- Gin HTTP Framework
- JWT Authentication + Refresh Token

### Problems to Fix

1. **Seeder file terlalu besar (3468 lines)** — HTML strings, CSS, dan JSON theme data semua di-hardcode di `cmd/seeder/main.go`
2. **Public handler inefficient** — `GetPublicWedding` fetches all guests lalu loop untuk cari slug (O(n) instead of O(1) query)
3. **No RSVP system** — Database tidak punya tabel untuk konfirmasi kehadiran tamu
4. **No pagination** — Semua `GetAll` endpoints return full dataset tanpa limit/offset
5. **No proper error handling structure** — Error responses tidak konsisten
6. **Missing API validation** — Request body tidak divalidasi secara strict
7. **No rate limiting** — Public endpoints tanpa proteksi abuse
8. **Guest slug lookup** — Repository punya `GetBySlug` tapi handler tidak memakainya
9. **Asset fetching di public** — `GetPublicWedding` memuat SEMUA assets, bukan hanya yang relevan untuk context tersebut

---

## REFACTORING TASKS

### 1. Database Schema Migration (`migrations/002_upgrade.sql`)

```sql
-- RSVP table for guest confirmations
CREATE TABLE IF NOT EXISTS rsvps (
    id SERIAL PRIMARY KEY,
    guest_id INTEGER NOT NULL REFERENCES guests(id) ON DELETE CASCADE,
    attendance VARCHAR(20) NOT NULL DEFAULT 'pending', -- 'attending', 'not_attending', 'pending'
    guest_count INTEGER DEFAULT 1,
    message TEXT DEFAULT '',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_guest_rsvp UNIQUE (guest_id)
);

-- Context-Asset junction table (asset dipakai per context, bukan global semua)
CREATE TABLE IF NOT EXISTS context_assets (
    id SERIAL PRIMARY KEY,
    context_id INTEGER NOT NULL REFERENCES contexts(id) ON DELETE CASCADE,
    asset_id INTEGER NOT NULL REFERENCES assets(id) ON DELETE CASCADE,
    role VARCHAR(50) DEFAULT 'general', -- 'background', 'music', 'photo', 'font', 'video'
    sort_order INTEGER DEFAULT 0,
    CONSTRAINT unique_context_asset UNIQUE (context_id, asset_id)
);

-- Add indexes for performance
CREATE INDEX IF NOT EXISTS idx_guests_context_slug ON guests(context_id, slug);
CREATE INDEX IF NOT EXISTS idx_contexts_slug ON contexts(slug);
CREATE INDEX IF NOT EXISTS idx_themes_slug ON themes(slug);
CREATE INDEX IF NOT EXISTS idx_rsvps_guest_id ON rsvps(guest_id);
CREATE INDEX IF NOT EXISTS idx_context_assets_context ON context_assets(context_id);
```

### 2. Domain Layer Updates

#### [MODIFY] `internal/domain/guest.go`
- Tambah field RSVP status di struct Guest (optional relation)

#### [NEW] `internal/domain/rsvp.go`
```go
type RSVP struct {
    ID         int64     `json:"id"`
    GuestID    int64     `json:"guest_id"`
    Attendance string    `json:"attendance"` // attending, not_attending, pending
    GuestCount int       `json:"guest_count"`
    Message    string    `json:"message"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
}

type RSVPRepository interface {
    GetByGuestID(ctx context.Context, guestID int64) (*RSVP, error)
    GetAllByContextID(ctx context.Context, contextID int64) ([]RSVP, error)
    Upsert(ctx context.Context, rsvp *RSVP) error
}
```

#### [NEW] `internal/domain/context_asset.go`
```go
type ContextAsset struct {
    ID        int64  `json:"id"`
    ContextID int64  `json:"context_id"`
    AssetID   int64  `json:"asset_id"`
    Role      string `json:"role"` // background, music, photo, font, video
    SortOrder int    `json:"sort_order"`
}

type ContextAssetRepository interface {
    GetByContextID(ctx context.Context, contextID int64) ([]Asset, error)
    Link(ctx context.Context, contextID int64, assetID int64, role string) error
    Unlink(ctx context.Context, contextID int64, assetID int64) error
}
```

#### [MODIFY] `internal/domain/theme.go`
- Ubah `ThemeData map[string]interface{}` menjadi typed struct:
```go
type ThemeData struct {
    GlobalStyle  GlobalStyle  `json:"globalStyle"`
    Splash       SplashConfig `json:"splash"`
    Sections     []Section    `json:"sections"`
    ColorPalette []string     `json:"colorPalette"`
}
```

### 3. Repository Layer

#### [NEW] `internal/repository/postgres/rsvp_repository.go`
- Implement RSVP CRUD
- `GetAllByContextID` → JOIN guests ON guests.context_id

#### [NEW] `internal/repository/postgres/context_asset_repository.go`
- Implement context-asset linking

#### [MODIFY] `internal/repository/postgres/guest_repository.go`
- Optimize `GetBySlug` query — pastikan pakai index

#### [MODIFY] `internal/repository/postgres/context_repository.go`
- Add pagination support (limit, offset, search parameters)

#### [MODIFY] `internal/repository/postgres/theme_repository.go`
- Add pagination support

### 4. Usecase Layer

#### [NEW] `internal/usecase/rsvp_usecase.go`
- Business logic for RSVP submission (validate attendance value, limit guest_count)
- GetAllByContextID — return RSVP data joined with guest names

#### [MODIFY] `internal/usecase/context_usecase.go`
- Add method `GetBySlugWithRelations` — preload theme + context-specific assets

#### [MODIFY] `internal/usecase/guest_usecase.go`
- Add `GetBySlug` method yang langsung pakai repository

### 5. Delivery/HTTP Layer

#### [MODIFY] `internal/delivery/http/public_handler.go`
**Critical Refactor:**
```go
// BEFORE (inefficient):
guests, _ := h.guestUsecase.GetAllByContextID(ctx, clientCtx.ID)
for _, g := range guests {
    if g.Slug == guestSlug { ... }
}

// AFTER (direct query):
guest, err := h.guestUsecase.GetBySlug(ctx, clientCtx.ID, guestSlug)
```
- Fetch only context-specific assets (via context_assets table), bukan semua assets
- Return structured RSVP status kalau ada

#### [NEW] `internal/delivery/http/rsvp_handler.go`
```
POST /v1/public/rsvp           — Guest submit RSVP (public, no auth)
GET  /v1/contexts/:id/rsvps    — Admin view all RSVPs for a context (auth)
```

#### [MODIFY] `internal/delivery/http/router.go`
- Register RSVP routes
- Add rate limiting middleware for public endpoints

#### [MODIFY] `internal/delivery/http/context_handler.go`
- Support pagination query params: `?page=1&limit=10&search=keyword`

#### [MODIFY] `internal/delivery/http/theme_handler.go`
- Support pagination query params

#### [NEW] `internal/delivery/http/validation.go`
- Centralized request validation helpers
- Custom error response format: `{ "error": { "code": "...", "message": "..." } }`

### 6. Seeder Refactor (`cmd/seeder/main.go`)

**Problem:** File 3468 lines dengan raw HTML/CSS strings hardcoded.

**Solution:**
- Extract theme HTML templates ke folder `cmd/seeder/templates/`
  - `cmd/seeder/templates/royal-gold.html`
  - `cmd/seeder/templates/modern-sinis.html`
  - `cmd/seeder/templates/vintage-romance.html`
- Extract theme JSON data ke `cmd/seeder/data/`
  - `cmd/seeder/data/royal-gold.json`
  - `cmd/seeder/data/modern-sinis.json`
  - `cmd/seeder/data/vintage-romance.json`
- `main.go` hanya baca file dan insert ke DB, max 200 lines

### 7. API Server (`cmd/api/main.go`)

#### [MODIFY] `cmd/api/main.go`
- Register RSVP handler
- Register ContextAsset handler
- Ensure proper graceful shutdown

---

## NEW API ENDPOINTS (Summary)

```
# RSVP (Public)
POST /v1/public/rsvp                    — Submit RSVP confirmation

# RSVP (Admin, Auth)
GET  /v1/contexts/:id/rsvps             — List all RSVPs for context

# Context Assets (Admin, Auth)  
POST /v1/contexts/:id/assets            — Link asset to context
DELETE /v1/contexts/:id/assets/:assetId — Unlink asset from context
GET  /v1/contexts/:id/assets            — List assets for context
```

---

## VERIFICATION PLAN

### Automated
```bash
# Build all binaries
go build -o cmd/api/api cmd/api/main.go
go build -o cmd/seeder/seeder cmd/seeder/main.go

# Run migration
psql $DATABASE_URL -f migrations/002_upgrade.sql

# Run seeder
./cmd/seeder/seeder

# Run API server
./cmd/api/api
```

### Manual
- Test RSVP endpoint: `curl -X POST localhost:8080/v1/public/rsvp -d '{"guest_id": 1, "attendance": "attending", "guest_count": 2, "message": "Selamat!"}'`
- Verify pagination: `curl localhost:8080/v1/themes?page=1&limit=10`
- Check public wedding loads context-specific assets only
- Verify error response format is consistent

---

## OUTPUT EXPECTATIONS FOR AI

1. Gunakan Clean Architecture — domain → repository → usecase → delivery.
2. Setiap file harus lengkap dan siap compile.
3. Tidak ada hardcoded HTML/CSS di Go source code.
4. Template HTML disimpan di file terpisah (`templates/`).
5. Theme JSON schema disimpan di file terpisah (`data/`).
6. Semua endpoint harus ada request validation.
7. Error response harus konsisten format.
8. Pagination built-in untuk semua list endpoints.
9. RSVP feature complete (submit + view + join guest data).
10. Context-Asset relation tabel baru untuk asset scoping.
11. File seeder max 200 lines (baca template files).
12. Public handler optimized (direct slug query, scoped assets).
13. Semua query pakai prepared statements (SQL injection safe).
14. Rate limit middleware untuk public endpoints.
15. No dead code, no duplicate code, no unused imports.
