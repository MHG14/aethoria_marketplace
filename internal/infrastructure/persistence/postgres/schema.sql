CREATE TYPE item_type AS ENUM ('common', 'rare', 'legendary');
CREATE TYPE item_status AS ENUM ('available', 'listed', 'auctioned');
CREATE TYPE listing_status AS ENUM ('active', 'sold', 'cancelled');
CREATE TYPE auction_status AS ENUM ('active', 'finished', 'cancelled');
CREATE TYPE trade_type AS ENUM ('auction', 'listing');
CREATE TYPE tx_type AS ENUM ('reserve', 'release', 'debit', 'credit');
CREATE TYPE ref_type AS ENUM ('auction', 'listing', 'trade');

CREATE TABLE guilds (
    id             BIGSERIAL PRIMARY KEY,
    name           TEXT NOT NULL UNIQUE,
    total_money    BIGINT NOT NULL DEFAULT 0,
    reserved_money BIGINT NOT NULL DEFAULT 0,
    daily_limit    BIGINT NOT NULL DEFAULT 0,
    daily_spent    BIGINT NOT NULL DEFAULT 0
);

CREATE TABLE items (
    id       BIGSERIAL PRIMARY KEY,
    name     TEXT NOT NULL,
    type     item_type NOT NULL,
    status   item_status NOT NULL DEFAULT 'available',
    owner_id BIGINT NOT NULL REFERENCES guilds(id)
);

CREATE TABLE listings (
    id         BIGSERIAL PRIMARY KEY,
    item_id    BIGINT NOT NULL REFERENCES items(id),
    seller_id  BIGINT NOT NULL REFERENCES guilds(id),
    buyer_id   BIGINT REFERENCES guilds(id),
    price      BIGINT NOT NULL,
    status     listing_status NOT NULL DEFAULT 'active',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE auctions (
    id                BIGSERIAL PRIMARY KEY,
    item_id           BIGINT NOT NULL REFERENCES items(id),
    seller_id         BIGINT NOT NULL REFERENCES guilds(id),
    starting_price    BIGINT NOT NULL,
    highest_bid       BIGINT NOT NULL DEFAULT 0,
    highest_bidder_id BIGINT REFERENCES guilds(id),
    end_time          TIMESTAMPTZ NOT NULL,
    original_end_time TIMESTAMPTZ NOT NULL,
    status            auction_status NOT NULL DEFAULT 'active',
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE bids (
    id           BIGSERIAL PRIMARY KEY,
    auction_id   BIGINT NOT NULL REFERENCES auctions(id),
    guild_id     BIGINT NOT NULL REFERENCES guilds(id),
    amount       BIGINT NOT NULL,
    is_cancelled BOOLEAN NOT NULL DEFAULT FALSE,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE trades (
    id         BIGSERIAL PRIMARY KEY,
    item_id    BIGINT NOT NULL REFERENCES items(id),
    seller_id  BIGINT NOT NULL REFERENCES guilds(id),
    buyer_id   BIGINT NOT NULL REFERENCES guilds(id),
    price      BIGINT NOT NULL,
    type       trade_type NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE wallet_transactions (
    id         BIGSERIAL PRIMARY KEY,
    guild_id   BIGINT NOT NULL REFERENCES guilds(id),
    type       tx_type NOT NULL,
    amount     BIGINT NOT NULL,
    ref_type   ref_type NOT NULL,
    ref_id     BIGINT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);