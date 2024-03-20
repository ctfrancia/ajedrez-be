ALTER TABLE users DROP CONSTRAINT username_unique_check UNIQUE (username);
ALTER TABLE users DROP CONSTRAINT email_unique_check UNIQUE (email);
