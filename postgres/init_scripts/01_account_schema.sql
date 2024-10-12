CREATE SCHEMA account;

CREATE TABLE fishing_hell.account.user(
    id uuid PRIMARY KEY,
    username varchar(32) NOT NULL,
    email varchar(320) NOT NULL,
    password varchar(30) NOT NULL
)
