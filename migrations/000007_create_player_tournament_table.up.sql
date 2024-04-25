-- this table references the users table and the tournaments table in a many-to-many relationship
CREATE TABLE IF NOT EXISTS player_tournament (
  player_id INT REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
  tournament_id INT REFERENCES tournaments(id) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT player_tournament_pk PRIMARY KEY (player_id, tournament_id)
);
