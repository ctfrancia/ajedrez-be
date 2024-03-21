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
- for migrations and seeds, see `migrations/` directories (brew install golang-migrate)

### clubs
#### Description
TODO

### Definitions and descriptions
- `club_id` - primary key
- `is_active` - boolean to determine if club is active
- `created_at` - date club was created
- `updated_at` - date club was last updated
- `deleted_at` - date club was deleted
- `name` - club's name
- `description` - club's description
- `avatar` - club's profile picture URL
- `organizer_id` - foreign key to `users` table
- `email` - club's email
- `phone` - club's phone number
- `website` - club's website
- `country` - club's country of residence (SPAIN, USA, etc.) DEFAULT: SPAIN
- `province` - club's province of residence (MADRID, BARCELONA, etc.)
- `city` - club's city of residence (ALCOBENDAS, SANT CUGAT, etc.)
- `address` - club's address
- `members` -  foreign keys[] to `users` table
```sql
CREATE TABLE IF NOT EXISTS clubs (
    club_id bigserial PRIMARY KEY,
    is_active boolean NOT NULL DEFAULT TRUE,
    is_verified boolean NOT NULL DEFAULT FALSE,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    deleted_at timestamp(0) with time zone,
    code text NOT NULL,
    name text NOT NULL,
    address text NOT NULL,
    observations text,
    city text NOT NULL
);
```

### users
#### Description
TODO

### Definitions and descriptions
- `user_id` - primary key
- `is_active` - boolean to determine if user is active
- `is_verified` - boolean to determine if user is verified
- `is_admin` - boolean to determine if user is admin
- `club_admin` - foreign key to `clubs` table (what club the user is admin of)
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
- `is_arbiter` - boolean to determine if user is an arbiter
- `is_coach` - boolean to determine if user is a coach
- `title` - user's title (GM, IM, FM, etc.)
- `chess_com_username` - user's chess.com username
- `lichess_username` - user's lichess username
- `chess24_username` - user's chess24 username
- `country` - user's country of residence (SPAIN, USA, etc.) DEFAULT: SPAIN
- `province` - user's province of residence (MADRID, BARCELONA, etc.)
- `city` - user's city of residence (ALCOBENDAS, SANT CUGAT, etc.)
- `neighborhood` - user's neighborhood of residence (LA MORALEJA, VALLVIDRERA, etc.)
```sql
CREATE TABLE IF NOT EXISTS users (
    user_id bigserial PRIMARY KEY,
    is_active boolean NOT NULL DEFAULT FALSE,
    is_verified boolean NOT NULL DEFAULT FALSE,
    is_admin_of_club boolean NOT NULL DEFAULT FALSE,
    club_admin_of bigint, -- create fk
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    deleted_at timestamp(0) with time zone,
    first_name text NOT NULL,
    last_name text NOT NULL,
    dob date,
    sex text,
    username text,
    email text,
    password text,
    password_reset_token text,
    avatar text,
    club_id bigint,
    club_role_id bigint, -- create fk
    about_me text,
    elo_fide integer NOT NULL DEFAULT 1200,
    elo_national integer NOT NULL DEFAULT 1200,
    elo_regional integer NOT NULL DEFAULT 1200,
    is_arbiter boolean NOT NULL DEFAULT FALSE,
    is_coach boolean NOT NULL DEFAULT FALSE,
    price_per_hour integer NOT NULL DEFAULT 0,
    chess_com_username text NOT NULL DEFAULT '',
    lichess_username text NOT NULL DEFAULT '',
    chess24_username text NOT NULL DEFAULT '',
    country text NOT NULL DEFAULT 'SPAIN',
    province text NOT NULL DEFAULT '',
    city text NOT NULL DEFAULT '',
    neighborhood text NOT NULL DEFAULT ''
);
```

### tournaments
#### Description
Tournaments represent the parent of the `games` table and will define the rules
and regulations of the subsequent games of the child games.

#### Definitions and descriptions
- `tournament_id` - primary key
- `is_active` - boolean to determine if tournament is active
- `created_at` - date tournament was created
- `updated_at` - date tournament was last updated
- `deleted_at` - date tournament was deleted
- `start_at` - date tournament was started
- `end_at` - date tournament was ended
- `no_of_rounds` - number of rounds in tournament
- `time_control` - time control of tournament (Standard, Rapid, Blitz)
- `clock_type` - clock type of tournament (Analog, Digital)
- `clock_rythm` - clock increment of tournament(50+10, 25+5, etc.)
- `arbiters` - foreign keys[] to `users` table
- `location` - tournament's location
- `fide_valid` - boolean to determine if tournament is FIDE valid
- `national_valid` - boolean to determine if tournament is national valid
- `regional_valid` - boolean to determine if tournament is regional valid
- `organizer_id` - foreign key to `users` table

```sql
CREATE TABLE IF NOT EXISTS tournaments (
    tournament_id bigserial PRIMARY KEY,  
    is_active boolean NOT NULL DEFAULT TRUE,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    deleted_at timestamp(0) with time zone,
    name text NOT NULL,
    description text,
    start_date timestamp(0) with time zone NOT NULL,
    end_date timestamp(0) with time zone NOT NULL,
    no_of_rounds integer NOT NULL DEFAULT 0,
    time_control text NOT NULL, -- create fk to time_control table
    clock_type text NOT NULL,
    clock_rhythm text NOT NULL,
    aribiters bigint[] NOT NULL DEFAULT '{}', -- create fk
    location text,
    organizer_id bigint,
);
```
### `games`
#### Description
A game is also considered a round or a match. For example a tournament can be composed
of 9 rounds. Those 9 rounds can be considered 9 games and will thus be stored in the
`games` table.

### Definitions and descriptions
- `game_id` - primary key
- `is_active` - boolean to determine if game is active
- `created_at` - date game was created
- `updated_at` - date game was last updated
- `deleted_at` - date game was deleted
- `start_at` - date game was started
- `end_at` - date game was ended
- `location` - game's location
- `fide_valid` - boolean to determine if game is FIDE valid
- `national_valid` - boolean to determine if game is national valid
- `regional_valid` - boolean to determine if game is regional valid
- `organizer_id` - foreign key to `users` table
- `organizer_email` - organizer's email
- `organizer_phone` - organizer's phone number
- `players_attending` - foreign keys[] to `users` table of members that are registered to the website
- `club_members_price` - price for club members
- `club_non_members_price` - price for non-club members
- `qr_code` - URL to QR code for game
- `description` - game's description
- `additional_info` - game's additional info
- `tournament_id` - foreign key to `tournaments` table

```sql
CREATE TABLE IF NOT EXISTS games (
    game_id bigserial PRIMARY KEY,  
    is_active boolean NOT NULL DEFAULT TRUE,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    deleted_at timestamp(0) with time zone,
    start_at timestamp(0) with time zone NOT NULL,
    end_at timestamp(0) with time zone,
    location text,
    fide_valid boolean NOT NULL DEFAULT FALSE,
    national_valid boolean NOT NULL DEFAULT FALSE,
    regional_valid boolean NOT NULL DEFAULT FALSE,
    organizer_id bigint,
    organizer_email text NOT NULL,
    organizer_phone text NOT NULL,
    players_attending integer[] NOT NULL DEFAULT '{}',
    club_member_price integer NOT NULL DEFAULT 0,
    club_non_member_price integer NOT NULL DEFAULT 0,
    qr_code text,
    description text,
    additional_info text,
    tournament_id bigint,
);
```

### `club_roles`
#### Description
TODO

### Definitions and descriptions
TODO
