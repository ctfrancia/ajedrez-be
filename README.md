# ajedrez-be
back-end code for chess website

## Project setup
- have Go installed
- have a PostgreSQL database running
    - create a database called `chess`. `psql=# CREATE DATABASE chess;`
    - connect to the database. `psql=# \c chess`
    - create a user called `chess` with password `che55`. `chess=# CREATE ROLE chess WITH LOGIN PASSWORD 'che55';`
    - create extension. `CREATE EXTENSION IF NOT EXISTS citext;`
        - extension is used for case-insensitive text
    - make sure you can connect fine: `psql --host=localhost --dbname=chess --username=chess`
    - to find out where your conf file is located, run `psql postgres -c 'SHOW config_file;'` or `sudo -u postgres psql -c 'SHOW config_file;'`
        - this is for further configuration/tweaking of the file for localhost performance


## DB Column Explanations
- for migrations and seeds, see `migrations/` and `seeds/` directories (brew install golang-migrate)

### `users`
- `id` - primary key
- `is_active` - boolean to determine if user is active
- `created_at` - date user was created
- `updated_at` - date user was last updated
- `deleted_at` - date user was deleted
- `first_name` - user's first name
- `last_name` - user's last name
- `dob` - user's date of birth
    - this is for tournament search purposes
- `sex` - user's personal identity(male or female)
    - this is for tournament search purposes
- `username` - unique username
- `email` - unique email
- `password` - hashed password
- `avatar` - user's profile picture URL
- `club_id` - foreign key to `clubs` table
- `club_role_id` - foreign key to `club_roles` table
- `about_me` - user's personal bio
- `elo_fide` - user's FIDE rating
- `elo_national` - user's national rating
- `elo_regional` - user's regional rating
- `chess_com_username` - user's chess.com username
- `lichess_username` - user's lichess username
- `chess24_username` - user's chess24 username
- `country` - user's country of residence (SPAIN, USA, etc.) DEFAULT: SPAIN
- `province` - user's province of residence (MADRID, BARCELONA, etc.)
- `city` - user's city of residence (ALCOBENDAS, SANT CUGAT, etc.)
- `neighborhood` - user's neighborhood of residence (LA MORALEJA, VALLVIDRERA, etc.)
- `version` - version of this database schema

### `clubs`

### `games`

### `tournaments`

### `club_roles`
