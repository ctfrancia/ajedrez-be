CREATE TABLE IF NOT EXISTS tournaments (
    id BIGSERIAL PRIMARY KEY AUTOINCREMENT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    is_verified BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    updated_by BIGINT REFERENCES users,
    deleted_at TIMESTAMP(0) with time zone,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,

    is_online BOOLEAN NOT NULL DEFAULT FALSE,
    is_over_the_board BOOLEAN NOT NULL DEFAULT FALSE,
    is_hybrid BOOLEAN NOT NULL DEFAULT FALSE,

    is_team BOOLEAN NOT NULL DEFAULT FALSE,
    is_individual BOOLEAN NOT NULL DEFAULT FALSE,

    is_rated BOOLEAN NOT NULL DEFAULT FALSE,
    is_unrated BOOLEAN NOT NULL DEFAULT FALSE,

    is_private BOOLEAN NOT NULL DEFAULT FALSE,
    is_public BOOLEAN NOT NULL DEFAULT FALSE,

    public_cost INTEGER NOT NULL DEFAULT 0,
    private_cost INTEGER NOT NULL DEFAULT 0,
    currency TEXT NOT NULL DEFAULT '',

    is_open BOOLEAN NOT NULL DEFAULT FALSE,
    is_invitational BOOLEAN NOT NULL DEFAULT FALSE,

    code TEXT NOT NULL DEFAULT uuid_generate_v1(),
    name TEXT NOT NULL,
    poster TEXT NOT NULL DEFAULT '',
    start_date TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    end_date TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    registration_start_date TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    registration_end_date TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    age_category TEXT NOT NULL DEFAULT '', -- if empty then all ages

    organizer BIGINT REFERENCES clubs,      -- can be a club 
    player_manager BIGINT REFERENCES users, -- or a user

    country TEXT NOT NULL DEFAULT '',
    province TEXT NOT NULL DEFAULT '',
    city TEXT NOT NULL DEFAULT '',
    address TEXT NOT NULL DEFAULT '',
    observations TEXT NOT NULL DEFAULT '',

    players BIGINT[] NOT NULL DEFAULT '{}',
    teams BIGINT[] NOT NULL DEFAULT '{}',

    max_atendees INTEGER NOT NULL DEFAULT 0,
    min_atendees INTEGER NOT NULL DEFAULT 0,

    version INTEGER NOT NULL DEFAULT 1
);
