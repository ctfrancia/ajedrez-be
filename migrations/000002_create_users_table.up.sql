CREATE TABLE IF NOT EXISTS users (
    user_id bigserial PRIMARY KEY,
    is_active boolean NOT NULL DEFAULT FALSE,
    is_verified boolean NOT NULL DEFAULT FALSE,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    soft_deleted bool NOT NULL DEFAULT FALSE,

    user_code text NOT NULL DEFAULT uuid_generate_v1(),
    first_name text NOT NULL DEFAULT '',
    last_name text NOT NULL DEFAULT '',
    username text NOT NULL DEFAULT '',
    password text NOT NULL DEFAULT '',
    password_reset_token text NOT NULL DEFAULT '',
    email text NOT NULL DEFAULT '',
    avatar text NOT NULL DEFAULT '',
    dob date NOT NULL DEFAULT '1900-01-01',
    about_me text NOT NULL DEFAULT '',
    sex text NOT NULL DEFAULT '',

    club_id bigint NOT NULL DEFAULT 0, -- fk to club table
    chess_age_category text NOT NULL DEFAULT '',

    elo_fide_standard integer NOT NULL DEFAULT 1200,
    elo_fide_rapid integer NOT NULL DEFAULT 1200,
    elo_fide_blitz integer NOT NULL DEFAULT 1200,
    elo_fide_bullet integer NOT NULL DEFAULT 1200,

    elo_national_standard integer NOT NULL DEFAULT 1200,
    elo_national_rapid integer NOT NULL DEFAULT 1200,
    elo_national_blitz integer NOT NULL DEFAULT 1200,
    elo_national_bullet integer NOT NULL DEFAULT 1200,

    elo_regional_standard integer NOT NULL DEFAULT 1200,
    elo_regional_rapid integer NOT NULL DEFAULT 1200,
    elo_regional_blitz integer NOT NULL DEFAULT 1200,
    elo_regional_bullet integer NOT NULL DEFAULT 1200,

    is_arbiter boolean NOT NULL DEFAULT FALSE,
    is_coach boolean NOT NULL DEFAULT FALSE,
    price_per_hour float NOT NULL DEFAULT 0,
    currency text NOT NULL DEFAULT '',
    chess_com_username text NOT NULL DEFAULT '',
    lichess_username text NOT NULL DEFAULT '',
    chess24_username text NOT NULL DEFAULT '',

    country text NOT NULL DEFAULT '',
    province text NOT NULL DEFAULT '',
    city text NOT NULL DEFAULT '',
    neighborhood text NOT NULL DEFAULT '',

    version integer NOT NULL DEFAULT 0
);
