CREATE TABLE IF NOT EXISTS tournaments (
    id bigserial PRIMARY KEY,
    is_active boolean NOT NULL DEFAULT TRUE,
    is_verified boolean NOT NULL DEFAULT FALSE,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_by bigint REFERENCES users,
    deleted_at timestamp(0) with time zone,
    is_deleted boolean NOT NULL DEFAULT FALSE,

    is_online boolean NOT NULL DEFAULT FALSE,
    is_over_the_board boolean NOT NULL DEFAULT FALSE,
    is_hybrid boolean NOT NULL DEFAULT FALSE,

    is_team boolean NOT NULL DEFAULT FALSE,
    is_individual boolean NOT NULL DEFAULT FALSE,

    is_rated boolean NOT NULL DEFAULT FALSE,
    is_unrated boolean NOT NULL DEFAULT FALSE,

    is_private boolean NOT NULL DEFAULT FALSE,
    is_public boolean NOT NULL DEFAULT FALSE,

    public_cost integer NOT NULL DEFAULT 0,
    private_cost integer NOT NULL DEFAULT 0,
    currency text NOT NULL DEFAULT '',

    is_open boolean NOT NULL DEFAULT FALSE,
    is_invitational boolean NOT NULL DEFAULT FALSE,

    code text NOT NULL DEFAULT uuid_generate_v1(),
    name text NOT NULL,
    start_date timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    end_date timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    organizer bigint REFERENCES clubs,
    country text NOT NULL DEFAULT '',
    province text NOT NULL DEFAULT '',
    city text NOT NULL DEFAULT '',
    address text NOT NULL,
    poster text NOT NULL DEFAULT '',
    observations text NOT NULL DEFAULT ''
);
