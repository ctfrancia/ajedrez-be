ALTER TABLE users ADD CONSTRAINT fk_user_club FOREIGN KEY (club_id) REFERENCES clubs (club_id) ON DELETE CASCADE;
ALTER TABLE users ADD CONSTRAINT fk_admin_of FOREIGN KEY (club_admin_of) REFERENCES clubs (club_id) ON DELETE CASCADE;
