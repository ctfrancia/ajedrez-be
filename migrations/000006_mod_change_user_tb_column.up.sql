ALTER TABLE users DROP COLUMN IF EXISTS chess_category;
ALTER TABLE users ADD COLUMN IF NOT EXISTS chess_age_category text;
