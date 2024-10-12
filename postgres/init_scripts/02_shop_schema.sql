CREATE SCHEMA shop;

CREATE TABLE fishing_hell.shop.bait(
    id int primary key,
    name varchar(20) NOT NULL,
    value float NOT NULL CHECK ( value > 0 ),
    uses int NOT NULL
);

CREATE TABLE fishing_hell.shop.pole(
    id int primary key,
    name varchar(30) NOT NULL,
    ticks int NOT NULL,
    price float NOT NULL
);

CREATE TABLE fishing_hell.shop.string(
    id int primary key,
    name varchar(40) NOT NULL,
    weight float NOT NULL,
    price float NOT NULL
);
-- IN CODE: first check for whether the primary key of this user and this bait already exist
-- (does this user already own some of this bait), if so, just add the number of bait.uses to
-- baitOwned.owned. If not then create it and set owned equal to bait.uses.
-- Whenever bait.uses == 0, delete the row.
CREATE TABLE fishing_hell.shop.baitOwned(
    userID uuid REFERENCES fishing_hell.account.user(id),
    baitID int REFERENCES fishing_hell.shop.bait(id),
    PRIMARY KEY (userID, baitID),
    owned int NOT NULL
);

-- Step 1: Create the function to delete the row when owned becomes 0
CREATE OR REPLACE FUNCTION delete_row_if_owned_zero()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.owned = 0 THEN
        DELETE FROM fishing_hell.shop.baitOwned
        WHERE userID = OLD.userID AND baitID = OLD.baitID;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Step 2: Create the trigger to call the function on update
CREATE TRIGGER delete_if_owned_zero
AFTER UPDATE ON fishing_hell.shop.baitOwned
FOR EACH ROW
WHEN (NEW.owned = 0)
EXECUTE FUNCTION delete_row_if_owned_zero();

-- Once you buy a pole you own it forever and cannot repurchase it
-- Does not come strung
CREATE TABLE fishing_hell.shop.poleOwned(
    userID uuid REFERENCES fishing_hell.account.user(id),
    poleID int REFERENCES fishing_hell.shop.pole(id),
    strung boolean DEFAULT false,
    PRIMARY KEY (userID, poleID)
);

-- Cannot own multiples of string
-- Can equip string to pole which will then remove it from your inventory and put it on the pole
CREATE TABLE fishing_hell.shop.stringOwned(
    userID uuid REFERENCES fishing_hell.account.user(id),
    stringID int REFERENCES fishing_hell.shop.string(id),
    PRIMARY KEY (userID, stringID)
)