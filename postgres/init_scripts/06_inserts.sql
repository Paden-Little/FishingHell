-- Insert Fish
INSERT INTO fishing_hell.fisher.fish (id, name, taxonomy, low_weight, high_weight, peak, spread)
VALUES
    (1, 'Rainbow Trout', 'Oncorhynchus mykiss', 1.5, 7.0, 4.0, 1.5),
    (2, 'Largemouth Bass', 'Micropterus salmoides', 2.0, 12.0, 6.5, 2.0),
    (3, 'Northern Pike', 'Esox lucius', 5.0, 20.0, 12.0, 3.0),
    (4, 'Yellow Perch', 'Perca flavescens', 0.5, 2.5, 1.2, 0.7),
    (5, 'Walleye', 'Sander vitreus', 1.0, 8.0, 3.0, 1.2);

-- Insert Pools
INSERT INTO fishing_hell.fisher.pool (name)
VALUES
    ('Clear Lake'),
    ('Green River');

-- Insert Catch Values (assumes fish and pool IDs already in `fish` and `pool` tables)
INSERT INTO fishing_hell.fisher.catch_value (fishID, poolID, value)
VALUES
    (1, 'Clear Lake', 15.0),
    (2, 'Clear Lake', 20.0),
    (3, 'Green River', 25.0),
    (4, 'Green River', 12.5),
    (5, 'Clear Lake', 18.0);

-- Insert Bait
INSERT INTO fishing_hell.shop.bait (id, name, value, uses)
VALUES
    (101, 'Worm', 1.5, 10),
    (102, 'Minnow', 2.0, 5);

-- Insert Pole (sample entries)
INSERT INTO fishing_hell.shop.pole (id, name, ticks, price)
VALUES
    (1, 'Basic Pole', 10, 15.0),
    (2, 'Advanced Pole', 20, 50.0);

-- Insert String (sample entries)
INSERT INTO fishing_hell.shop.string (id, name, weight, price)
VALUES
    (1, 'Standard Line', 2.0, 5.0),
    (2, 'Heavy-Duty Line', 5.0, 10.0);

-- Update Attraction Values for Fish with Bait
INSERT INTO fishing_hell.fisher.attraction (fishID, baitID, value)
VALUES
    (1, 101, 1.75),  -- Rainbow Trout attraction to Worm
    (1, 102, 0.80),  -- Rainbow Trout attraction to Minnow
    (2, 101, 1.60),  -- Largemouth Bass attraction to Worm
    (2, 102, 0.85),  -- Largemouth Bass attraction to Minnow
    (3, 101, 1.55),  -- Northern Pike attraction to Worm
    (3, 102, 0.90),  -- Northern Pike attraction to Minnow
    (4, 101, 1.50),  -- Yellow Perch attraction to Worm
    (4, 102, 0.65),  -- Yellow Perch attraction to Minnow
    (5, 101, 1.70),  -- Walleye attraction to Worm
    (5, 102, 0.75);  -- Walleye attraction to Minnow

INSERT INTO fishing_hell.account.user (username, email, password)
VALUES (
        'admin', 'admin@admin.com', 'admin'
       );

-- Here angler_one owns some Worms and Minnows, angler_two owns Shrimp
INSERT INTO fishing_hell.shop.baitOwned (userID, baitID, owned) VALUES
((SELECT id FROM fishing_hell.account.user WHERE username = 'admin'), 101, 10);

INSERT INTO fishing_hell.shop.baitOwned (userID, baitID, owned) VALUES
((SELECT id FROM fishing_hell.account.user WHERE username = 'admin'), 102, 10);

INSERT INTO fishing_hell.shop.poleowned (userid, poleid) VALUES (
(SELECT id FROM fishing_hell.account.user WHERE username = 'admin'), 1
);

INSERT INTO fishing_hell.shop.poleowned (userid, poleid) VALUES (
(SELECT id FROM fishing_hell.account.user WHERE username = 'admin'), 2
);

INSERT INTO fishing_hell.fisher.caught (id, fishid, weight, whencaught, frompool, baitused, caughtby, owner) VALUES (
(SELECT id FROM fishing_hell.account.user WHERE username = 'admin'), 1, 69.87, current_timestamp, 'Clear Lake', 102,
(SELECT id FROM fishing_hell.account.user WHERE username = 'admin')
, (SELECT id FROM fishing_hell.account.user WHERE username = 'admin')
);

