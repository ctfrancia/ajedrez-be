CREATE TABLE IF NOT EXISTS tournaments (
    id BIGSERIAL PRIMARY KEY,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    is_verified BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    updated_by BIGINT REFERENCES users,
    deleted_at TIMESTAMP(0) with time zone,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    is_online BOOLEAN NOT NULL DEFAULT FALSE,
    online_link TEXT NOT NULL DEFAULT '',
    is_otb BOOLEAN NOT NULL DEFAULT FALSE,
    is_hybrid BOOLEAN NOT NULL DEFAULT FALSE,

    is_team BOOLEAN NOT NULL DEFAULT FALSE,
    is_individual BOOLEAN NOT NULL DEFAULT FALSE,
    is_team_individual BOOLEAN NOT NULL DEFAULT FALSE,

    is_rated BOOLEAN NOT NULL DEFAULT FALSE,
    is_unrated BOOLEAN NOT NULL DEFAULT FALSE,

    match_making TEXT NOT NULL DEFAULT '', -- swiss, round-robin, knockout, etc

    is_private BOOLEAN NOT NULL DEFAULT FALSE, -- public cannot watch
    is_public BOOLEAN NOT NULL DEFAULT FALSE, -- public can watch

    member_cost INTEGER NOT NULL DEFAULT 0, -- 2000 = 20.00
    public_cost INTEGER NOT NULL DEFAULT 0,
    currency CITEXT NOT NULL DEFAULT 'EUR',

    is_open BOOLEAN NOT NULL DEFAULT FALSE,
    is_closed BOOLEAN NOT NULL DEFAULT FALSE,

    code UUID NOT NULL DEFAULT uuid_generate_v1(),
    name TEXT NOT NULL,
    poster TEXT NOT NULL DEFAULT '',
    dates TIMESTAMP(0) with time zone[] NOT NULL DEFAULT '{}',
    -- start_date TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    -- end_date TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    -- rest_date TIMESTAMP(0) with time zone,
    location TEXT NOT NULL DEFAULT '',
    registration_start_date TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    registration_end_date TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    age_category TEXT NOT NULL DEFAULT '', -- if empty then all ages
    time_control TEXT NOT NULL DEFAULT '', -- 45+10, 90+30, etc
    type TEXT NOT NULL DEFAULT '', -- classical, rapid, blitz, bullet, etc
    rounds INTEGER NOT NULL DEFAULT 0,

    organizer BIGINT REFERENCES clubs,      -- can be a club 
    user_organizer BIGINT REFERENCES users, -- or a user
    contact_email TEXT NOT NULL DEFAULT '',
    contact_phone TEXT NOT NULL DEFAULT '', -- include +country code

    country TEXT NOT NULL DEFAULT '',
    province TEXT NOT NULL DEFAULT '',
    city TEXT NOT NULL DEFAULT '',
    address TEXT NOT NULL DEFAULT '',
    postal_code TEXT NOT NULL DEFAULT '',
    observations TEXT NOT NULL DEFAULT '',
    is_cancelled BOOLEAN NOT NULL DEFAULT FALSE,

    players INTEGER[] NOT NULL DEFAULT '{}', -- user.ids
    teams INTEGER[] NOT NULL DEFAULT '{}', -- team.ids
    
    max_attendees INTEGER NOT NULL DEFAULT 0,
    min_attendees INTEGER NOT NULL DEFAULT 0,

    completed BOOLEAN NOT NULL DEFAULT FALSE,

    is_draft BOOLEAN NOT NULL DEFAULT TRUE,
    is_published BOOLEAN NOT NULL DEFAULT FALSE,
    version INTEGER NOT NULL DEFAULT 1
);
