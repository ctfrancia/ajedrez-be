CREATE TABLE IF NOT EXISTS dates (
  id BIGSERIAL PRIMARY KEY,
  date TIMESTAMP(0) WITH TIME ZONE UNIQUE
);
