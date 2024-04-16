# ajedrez-be
back-end code for chess website

## Project setup
- have Go installed
- have a PostgreSQL database running
- have within your profile (.zshrc for example) a exported variable `CHESS_DB_DSN` with your Postgres connection string
- [have go-migrate installed](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) 
- set up database, tables, user and extensions with:
```bash
make setup-dev
```
### NOTEs
- if you are using a different database, you will need to change the `CHESS_DB_DSN` variable in the `Makefile` to match your database connection string

- For any errors create a new issue in the repository and I will try to help you out and update README

TODO: add more setup instructions for Linux and Windows
 - especially for the migrate cli tool and other dependencies

## DB Column Explanations
*NOTE* this is currently changing frequently as the project is in development

### clubs
#### Description
TODO
```
    Column    |            Type             | Collation | Nullable |                Default
--------------+-----------------------------+-----------+----------+----------------------------------------
 club_id      | bigint                      |           | not null | nextval('clubs_club_id_seq'::regclass)
 is_active    | boolean                     |           | not null | true
 is_verified  | boolean                     |           | not null | false
 created_at   | timestamp(0) with time zone |           | not null | now()
 updated_at   | timestamp(0) with time zone |           | not null | now()
 deleted_at   | timestamp(0) with time zone |           |          |
 code         | text                        |           | not null |
 name         | text                        |           | not null |
 address      | text                        |           | not null |
 observations | text                        |           |          |
 city         | text                        |           | not null |
 country      | text                        |           | not null | 'Spain'::text
Indexes:
    "clubs_pkey" PRIMARY KEY, btree (club_id)
Referenced by:
    TABLE "users" CONSTRAINT "fk_admin_of" FOREIGN KEY (club_admin_of) REFERENCES clubs(club_id) ON DELETE CASCADE
    TABLE "users" CONSTRAINT "fk_user_club" FOREIGN KEY (club_id) REFERENCES clubs(club_id) ON DELETE CASCADE
```

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
    id bigserial PRIMARY KEY,
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
```
                                                Table "public.users"
        Column         |            Type             | Collation | Nullable |                Default
-----------------------+-----------------------------+-----------+----------+----------------------------------------
 user_id               | bigint                      |           | not null | nextval('users_user_id_seq'::regclass)
 is_active             | boolean                     |           | not null | false
 is_verified           | boolean                     |           | not null | false
 is_admin_of_club      | boolean                     |           | not null | false
 club_admin_of         | bigint                      |           |          |
 created_at            | timestamp(0) with time zone |           | not null | now()
 updated_at            | timestamp(0) with time zone |           | not null | now()
 deleted_at            | timestamp(0) with time zone |           |          |
 first_name            | text                        |           | not null |
 last_name             | text                        |           | not null |
 dob                   | date                        |           |          |
 sex                   | text                        |           |          |
 username              | text                        |           |          |
 email                 | text                        |           |          |
 password              | text                        |           |          |
 password_reset_token  | text                        |           |          |
 avatar                | text                        |           |          |
 club_id               | bigint                      |           |          |
 club_role_id          | bigint                      |           |          |
 about_me              | text                        |           |          |
 is_arbiter            | boolean                     |           | not null | false
 is_coach              | boolean                     |           | not null | false
 price_per_hour        | integer                     |           | not null | 0
 chess_com_username    | text                        |           | not null | ''::text
 lichess_username      | text                        |           | not null | ''::text
 chess24_username      | text                        |           | not null | ''::text
 country               | text                        |           | not null | 'SPAIN'::text
 province              | text                        |           | not null | ''::text
 city                  | text                        |           | not null | ''::text
 neighborhood          | text                        |           | not null | ''::text
 elo_fide_standard     | integer                     |           |          |
 elo_fide_rapid        | integer                     |           |          |
 elo_national_standard | integer                     |           |          |
 elo_national_rapid    | integer                     |           |          |
 elo_regional_standard | integer                     |           |          |
 club_user_code        | text                        |           |          |
 chess_age_category    | text                        |           |          |
 elo_regional_rapid    | integer                     |           |          |
 Indexes:
    "users_pkey" PRIMARY KEY, btree (user_id)
    "users_club_user_code_unique" UNIQUE CONSTRAINT, btree (club_user_code)
Foreign-key constraints:
    "fk_admin_of" FOREIGN KEY (club_admin_of) REFERENCES clubs(club_id) ON DELETE CASCADE
    "fk_user_club" FOREIGN KEY (club_id) REFERENCES clubs(club_id) ON DELETE CASCADE
```
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
- `elo_fide_standard` - user's FIDE standard rating
- `elo_fide_rapid` - user's FIDE rapid rating
- `elo_national_standard` - user's national standard rating
- `elo_national_rapid` - user's national rapid rating
- `elo_regional_standard` - user's regional standard rating
- `elo_regional_rapid` - user's regional rapid rating
```sql
CREATE TABLE IF NOT EXISTS users (
    user_id bigserial PRIMARY KEY,
    is_active boolean NOT NULL DEFAULT FALSE,
    is_verified boolean NOT NULL DEFAULT FALSE,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    soft_deleted bool NOT NULL DEFAULT FALSE,

    user_code text NOT NULL DEFAULT uuid_generate_v1(),
    first_name text NOT NULL DEFAULT '',
    last_name text NOT NULL DEFAULT '',
    username text NOT NULL DEFAULT '',
    password text NOT NULL DEFAULT '',
    password_reset_token text NOT NULL DEFAULT '',
    email text NOT NULL DEFAULT '',
    avatar text NOT NULL DEFAULT '',
    dob date NOT NULL DEFAULT '1900-01-01',
    about_me text NOT NULL DEFAULT '',
    sex text NOT NULL DEFAULT '',

    club_id bigint NOT NULL DEFAULT 0, -- fk to club table
    chess_age_category text NOT NULL DEFAULT '',

    elo_fide_standard integer NOT NULL DEFAULT 1200,
    elo_fide_rapid integer NOT NULL DEFAULT 1200,
    elo_fide_blitz integer NOT NULL DEFAULT 1200,
    elo_fide_bullet integer NOT NULL DEFAULT 1200,

    elo_national_standard integer NOT NULL DEFAULT 1200,
    elo_national_rapid integer NOT NULL DEFAULT 1200,
    elo_national_blitz integer NOT NULL DEFAULT 1200,
    elo_national_bullet integer NOT NULL DEFAULT 1200,

    elo_regional_standard integer NOT NULL DEFAULT 1200,
    elo_regional_rapid integer NOT NULL DEFAULT 1200,
    elo_regional_blitz integer NOT NULL DEFAULT 1200,
    elo_regional_bullet integer NOT NULL DEFAULT 1200,

    is_arbiter boolean NOT NULL DEFAULT FALSE,
    is_coach boolean NOT NULL DEFAULT FALSE,
    price_per_hour float NOT NULL DEFAULT 0,
    currency text NOT NULL DEFAULT '',
    chess_com_username text NOT NULL DEFAULT '',
    lichess_username text NOT NULL DEFAULT '',
    chess24_username text NOT NULL DEFAULT '',

    country text NOT NULL DEFAULT '',
    province text NOT NULL DEFAULT '',
    city text NOT NULL DEFAULT '',
    neighborhood text NOT NULL DEFAULT '',

    version integer NOT NULL DEFAULT 0
);
```

### notes
In order for the person to be able to have an active account, they need to sign up
for the website.

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


## API Endpoints

### Users
- `POST /user/create` - create user
```json
{
    "code":398409,
    "first_name":" Christian",
    "last_name":"Francia",
    "sex":"M",
    "title":"",
    "chess_age_category":"Senior",
    "club_id":65,
    "country":"Spain",
    "elo_fide_standard":0,
    "elo_fide_rapid":0,
    "elo_national_standard":0,
    "elo_national_rapid":0,
    "elo_regional_standard":1028,
    "elo_regional_rapid":1154
}
```
- `POST /user/login` - login user

