ALTER TABLE users ADD CONSTRAINT username_unique_check UNIQUE (username);
ALTER TABLE users ADD CONSTRAINT email_unique_check UNIQUE (email);
