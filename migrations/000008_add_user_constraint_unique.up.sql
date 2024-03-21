-- this constraint is wrong, fixed in 09
ALTER TABLE users ADD CONSTRAINT users_club_user_code_unique UNIQUE (username);
