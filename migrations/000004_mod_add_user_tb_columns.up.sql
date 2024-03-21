ALTER TABLE users DROP COLUMN IF EXISTS elo_fide;
ALTER TABLE users DROP COLUMN IF EXISTS elo_national;
ALTER TABLE users DROP COLUMN IF EXISTS elo_regional;

ALTER TABLE users ADD COLUMN IF NOT EXISTS chess_category text;

ALTER TABLE users ADD COLUMN IF NOT EXISTS elo_fide_standard INT;
ALTER TABLE users ADD COLUMN IF NOT EXISTS elo_fide_rapid INT;

ALTER TABLE users ADD COLUMN IF NOT EXISTS elo_national_standard INT;
ALTER TABLE users ADD COLUMN IF NOT EXISTS elo_national_rapid INT;

ALTER TABLE users ADD COLUMN IF NOT EXISTS elo_regional_standard INT;

-- duplicate column, fixed in 000008_mod_fix_elo_national_typo.up.sql
-- change column name to elo_national_standard to elo_regional_rapid
ALTER TABLE users ADD COLUMN IF NOT EXISTS elo_national_rapid INT;

