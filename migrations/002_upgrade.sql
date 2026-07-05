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

-- Context-Asset junction table
CREATE TABLE IF NOT EXISTS context_assets (
    id SERIAL PRIMARY KEY,
    context_id INTEGER NOT NULL REFERENCES contexts(id) ON DELETE CASCADE,
    asset_id INTEGER NOT NULL REFERENCES assets(id) ON DELETE CASCADE,
    role VARCHAR(50) DEFAULT 'general', -- 'background', 'music', 'photo', 'font', 'video'
    sort_order INTEGER DEFAULT 0,
    CONSTRAINT unique_context_asset UNIQUE (context_id, asset_id)
);

-- Add indexes for database search performance
CREATE INDEX IF NOT EXISTS idx_guests_context_slug ON guests(context_id, slug);
CREATE INDEX IF NOT EXISTS idx_contexts_slug ON contexts(slug);
CREATE INDEX IF NOT EXISTS idx_themes_slug ON themes(slug);
CREATE INDEX IF NOT EXISTS idx_rsvps_guest_id ON rsvps(guest_id);
CREATE INDEX IF NOT EXISTS idx_context_assets_context ON context_assets(context_id);
