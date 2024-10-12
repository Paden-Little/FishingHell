CREATE SCHEMA fisher;

CREATE TABLE fishing_hell.fisher.fish (
    id int PRIMARY KEY,
    name varchar(100) NOT NULL,
    taxonomy varchar(200) NOT NULL,
    low_weight float NOT NULL CONSTRAINT low_weight_must_be_positive CHECK (low_weight > 0),
    high_weight float NOT NULL CONSTRAINT high_weight_must_be_positive CHECK (high_weight > 0),
    peak float NOT NULL CONSTRAINT peak_must_be_zero_or_positive CHECK (peak >= 0),
    spread float NOT NULL CONSTRAINT spread_must_be_zero_or_positive CHECK (spread >= 0)
);

CREATE TABLE fishing_hell.fisher.pool (
    name varchar(50) PRIMARY KEY
);

CREATE TABLE fishing_hell.fisher.catch_value (
    fishID int REFERENCES fisher.fish(id) NOT NULL,
    poolID varchar(50) REFERENCES fisher.pool(name) NOT NULL,
    value float NOT NULL CONSTRAINT value_must_be_zero_or_positive CHECK (value >= 0)
);

CREATE TABLE fishing_hell.fisher.caught (
    -- ADD refernce to caught by, owner, baitUsed
    caughtID uuid primary key,
    fishID int REFERENCES fisher.fish(id) NOT NULL,
    weight float NOT NULL,
    baitUsed int NOT NULL,
    caughtBy uuid NOT NULL,
    owner uuid NOT NULL,
    frozen boolean default false
);