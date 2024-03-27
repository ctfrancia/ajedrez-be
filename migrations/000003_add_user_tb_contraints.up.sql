ALTER TABLE users ADD CONSTRAINT fk_user_club FOREIGN KEY (club_id) REFERENCES clubs (club_id) ON DELETE CASCADE;
ALTER TABLE users ADD CONSTRAINT users_club_user_code_unique UNIQUE (user_code);
