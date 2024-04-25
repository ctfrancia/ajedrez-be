CREATE TABLE IF NOT EXISTS date_tournament (
    date_id INT REFERENCES dates(id) ON DELETE NO ACTION,
    tournament_id INT REFERENCES tournaments(id) ON DELETE NO ACTION,
    CONSTRAINT date_tournament_pk PRIMARY KEY (date_id, tournament_id)
);
