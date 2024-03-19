CREATE TABLE IF NOT EXISTS tournaments (
    id bigserial PRIMARY KEY,  
    is_active boolean NOT NULL DEFAULT TRUE,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    deleted_at timestamp(0) with time zone,
    name text NOT NULL,
    description text,
    start_date timestamp(0) with time zone NOT NULL,
    end_date timestamp(0) with time zone NOT NULL,
    location text,
    organizer_id bigint,
);
