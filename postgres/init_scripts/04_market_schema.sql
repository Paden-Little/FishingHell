CREATE SCHEMA market;

CREATE TABLE fishing_hell.market.fish(
    id uuid primary key,
    fish uuid REFERENCES fishing_hell.fisher.caught(id),
    timeSold timestamp NOT NULL,
    value float NOT NULL
);

CREATE TABLE fishing_hell.market.bait(
    id uuid primary key,
    bait int REFERENCES fishing_hell.shop.bait(id),
    timeSold timestamp NOT NULL,
    value float NOT NULL
);