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
