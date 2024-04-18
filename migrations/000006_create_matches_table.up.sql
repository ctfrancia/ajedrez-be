CREATE TABLE IF NOT EXISTS matches (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    is_locked BOOLEAN NOT NULL DEFAULT FALSE,

    tournament_id INTEGER,
    white_id INTEGER,
    black_id INTEGER,
    winner_id INTEGER,
    round INTEGER,

    pgn TEXT,

    version INTEGER NOT NULL DEFAULT 1,
    FOREIGN KEY (tournament_id) REFERENCES tournaments (id),
    FOREIGN KEY (white_id) REFERENCES users (id),
    FOREIGN KEY (black_id) REFERENCES users (id),
    FOREIGN KEY (winner_id) REFERENCES users (id)
);
