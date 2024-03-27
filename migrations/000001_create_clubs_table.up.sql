CREATE TABLE IF NOT EXISTS clubs (
    club_id bigserial PRIMARY KEY,
    is_active boolean NOT NULL DEFAULT TRUE,
    is_verified boolean NOT NULL DEFAULT FALSE,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    deleted_at timestamp(0) with time zone,
    code text NOT NULL,
    name text NOT NULL,
    address text NOT NULL,
    observations text NOT NULL DEFAULT '',
    city text NOT NULL DEFAULT '',
    country text NOT NULL DEFAULT ''
);
