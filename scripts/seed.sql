-- guilds
INSERT INTO guilds (name, total_money, reserved_money, daily_limit, daily_spent)
VALUES
    ('Dragon Lords',  10000, 0, 5000, 0),
    ('Shadow Guild',  8000,  0, 4000, 0),
    ('Iron Wolves',   6000,  0, 3000, 0);

-- items
-- common items (owned by Dragon Lords = id 1)
INSERT INTO items (name, type, status, owner_id)
VALUES
    ('Iron Sword',      'common',    'available', 1),
    ('Health Potion',   'common',    'available', 1),
    ('Leather Armor',   'common',    'available', 2);

-- rare items
INSERT INTO items (name, type, status, owner_id)
VALUES
    ('Elven Bow',       'rare',      'available', 2),
    ('Mithril Shield',  'rare',      'available', 3);

-- legendary items (one per guild, these are the auction candidates)
INSERT INTO items (name, type, status, owner_id)
VALUES
    ('Soul Reaver',     'legendary', 'available', 1),
    ('Eye of Dragon',   'legendary', 'available', 2);