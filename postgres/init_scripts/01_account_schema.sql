CREATE SCHEMA account;

CREATE EXTENSION "uuid-ossp";

CREATE TABLE fishing_hell.account.user(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    cash float NOT NULL DEFAULT 0,
    username varchar(32) NOT NULL UNIQUE,
    email varchar(320) NULL,
    password varchar(30) NOT NULL
)
