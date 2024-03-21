ALTER TABLE users DROP CONSTRAINT users_club_user_code_unique;
-- fixed constraint prev. version had a wrong column name
ALTER TABLE users ADD CONSTRAINT users_club_user_code_unique UNIQUE (club_user_code);
