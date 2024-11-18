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
    value float NOT NULL CONSTRAINT value_must_be_zero_or_positive CHECK (value >= 0),
    PRIMARY KEY (fishID, poolID)
);

CREATE TABLE fishing_hell.fisher.caught (
    id uuid primary key default uuid_generate_v4(),
    fishID int REFERENCES fisher.fish(id) NOT NULL,
    weight float NOT NULL,
    whenCaught timestamp NOT NULL,
    fromPool varchar(50) NOT NULL REFERENCES fishing_hell.fisher.pool(name),
    baitUsed int NOT NULL REFERENCES fishing_hell.shop.bait(id),
    caughtBy uuid NOT NULL REFERENCES fishing_hell.account.user(id),
    owner uuid NOT NULL REFERENCES fishing_hell.account.user(id),
    frozen boolean default false
);

CREATE TABLE fishing_hell.fisher.attraction (
    id SERIAL PRIMARY KEY,  -- Automatically generated unique ID for each row
    fishID INT NOT NULL REFERENCES fishing_hell.fisher.fish(id),   -- Foreign key referencing a fish
    baitID INT NOT NULL REFERENCES fishing_hell.shop.bait(id),   -- Foreign key referencing a bait type
    value FLOAT NOT NULL  -- A value representing the attraction strength (e.g., probability or weight)
);