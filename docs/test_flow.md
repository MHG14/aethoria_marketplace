# Market Dragon Test Guide

Base URL:

```text
http://localhost:8080/api/v1
```

Send the following header with every request:

```text
Content-Type: application/json
```

---

## Setup

Start the project:

```bash
docker compose up --build
```

After the containers start, the database will automatically load the schema and seed data.

Guild IDs:

* Dragon Lords = 1
* Shadow Guild = 2
* Iron Wolves = 3

Item IDs:

* Iron Sword (common, owner=1) = 1
* Health Potion (common, owner=1) = 2
* Leather Armor (common, owner=2) = 3
* Elven Bow (rare, owner=2) = 4
* Mithril Shield (rare, owner=3) = 5
* Soul Reaver (legendary, owner=1) = 6
* Eye of Dragon (legendary, owner=2) = 7

---

# 1. Check Seed Data

```text
GET /guilds/1
GET /guilds/2
GET /guilds/3
```

Expected result:

Each guild should return its balance, daily_limit, and daily_spent = 0.

```text
GET /items
```

Expected result:

The API should return 7 items, and all of them should have `status=available`.

---

# 2. Top Up Wallet

```text
POST /guilds/1/topup
{ "amount": 10000 }

POST /guilds/2/topup
{ "amount": 10000 }

POST /guilds/3/topup
{ "amount": 10000 }
```

Expected result:

The wallet balance of each guild should increase.

```text
GET /guilds/1/wallet
```

Expected result:

```
total_money = 20000
available_balance = 20000
```

---

# 3. Listing Flow (Common and Rare Items)

## 3.1 Create a Listing

Dragon Lords lists the Iron Sword (item_id=1) for 200.

```text
POST /listings
{
    "item_id":   1,
    "seller_id": 1,
    "price":     200
}
```

Expected result:

The listing should be created with `status=active`.
Item 1 should now have `status=listed`.

```text
GET /items/1
```

Expected result:

```
status=listed
```

---

## 3.2 Try to List a Legendary Item (Should Fail)

```text
POST /listings
{
    "item_id":   6,
    "seller_id": 1,
    "price":     9999
}
```

Expected result:

422 error

```
legendary items cannot be listed, use auction instead
```

---

## 3.3 Try to List Someone Else's Item (Should Fail)

Shadow Guild tries to list an item owned by Dragon Lords.

```text
POST /listings
{
    "item_id":   2,
    "seller_id": 2,
    "price":     100
}
```

Expected result:

403 error

```
you do not own this item
```

---

## 3.4 Buy Your Own Listing (Should Fail)

```text
POST /listings/1/buy
{
    "buyer_id": 1
}
```

Expected result:

403 error

```
cannot buy your own listing
```

---

## 3.5 Successful Purchase

Shadow Guild buys the Iron Sword.

```text
POST /listings/1/buy
{
    "buyer_id": 2
}
```

Expected result:

A trade record should be returned with:

* buyer_id = 2
* seller_id = 1
* price = 200

Check item ownership:

```text
GET /items/1
```

Expected result:

```
owner_id=2
status=available
```

Check wallets:

```text
GET /guilds/2/wallet
```

Expected result:

```
total_money = 9800
```

```text
GET /guilds/1/wallet
```

Expected result:

```
total_money = 10200
```

Check transactions:

```text
GET /guilds/2/transactions
```

Expected result:

One transaction with:

```
type=debit
amount=200
```

```text
GET /guilds/1/transactions
```

Expected result:

One transaction with:

```
type=credit
amount=200
```

---

## 3.6 Buy the Same Listing Again (Should Fail)

```text
POST /listings/1/buy
{
    "buyer_id": 3
}
```

Expected result:

409 error

```
item is not available
```

---

## 3.7 Cancel a Listing

Dragon Lords lists the Health Potion and then cancels it.

```text
POST /listings
{
    "item_id":   2,
    "seller_id": 1,
    "price":     50
}
```

Save the new listing ID (expected to be 2).

```text
DELETE /listings/2
{
    "seller_id": 1
}
```

Expected result:

```
204 No Content
```

```text
GET /items/2
```

Expected result:

```
status=available
```

---

## 3.8 Cancel Someone Else's Listing (Should Fail)

Iron Wolves tries to cancel Dragon Lords' listing.

```text
DELETE /listings/2
{
    "seller_id": 3
}
```

Expected result:

403 error

```
you do not own this listing
```

---

## 3.9 Not Enough Money (Should Fail)

Create a guild with a small balance.

```text
POST /guilds
{
    "name":          "Broke Guild",
    "daily_limit":   99999,
    "initial_money": 10
}
```

Save the new guild ID (expected to be 4).

Now Shadow Guild lists the Leather Armor.

```text
POST /listings
{
    "item_id":   3,
    "seller_id": 2,
    "price":     500
}
```

The new listing ID should be 3.

Try to buy it.

```text
POST /listings/3/buy
{
    "buyer_id": 4
}
```

Expected result:

422 error

```
insufficient funds
```

---

## 3.10 Daily Limit (Should Fail)

Create another guild.

```text
POST /guilds
{
    "name":          "Limited Guild",
    "daily_limit":   100,
    "initial_money": 99999
}
```

The new guild ID should be 5.

```text
POST /listings/3/buy
{
    "buyer_id": 5
}
```

Expected result:

422 error

```
daily limit exceeded
```

---

# 4. Auction Flow (Legendary Items Only)

## 4.1 Create an Auction

Dragon Lords starts an auction for Soul Reaver.

```text
POST /auctions
{
    "item_id":        6,
    "seller_id":      1,
    "starting_price": 1000,
    "duration_hours": 24
}
```

Expected result:

Auction is created with:

```
status=active
auction_id=1
```

```text
GET /items/6
```

Expected result:

```
status=auctioned
```

---

## 4.2 Create Another Auction for the Same Item (Should Fail)

Expected result:

409 error

```
an active auction already exists for this item
```

---

## 4.3 Auction a Non-Legendary Item (Should Fail)

Expected result:

400 error

```
only legendary items can be auctioned
```

---

## 4.4 Seller Places a Bid (Should Fail)

```text
POST /auctions/1/bids
{
    "guild_id": 1,
    "amount": 1000
}
```

Expected result:

403 error

```
seller cannot bid on own item
```

---

## 4.5 First Valid Bid

Shadow Guild places the first bid.

```text
POST /auctions/1/bids
{
    "guild_id": 2,
    "amount": 1000
}
```

Expected result:

* bid_id = 1
* highest_bid = 1000
* highest_bidder_id = 2

Check wallet:

```text
GET /guilds/2/wallet
```

Expected result:

```
reserved_money=1000
```

Check transactions:

```text
GET /guilds/2/transactions
```

Expected result:

A transaction with:

```
type=reserve
amount=1000
```

---

## 4.6 Bid Below the Minimum (Should Fail)

The next minimum bid is 1050.

```text
POST /auctions/1/bids
{
    "guild_id": 3,
    "amount": 1049
}
```

Expected result:

422 error

```
bid must be at least 5% above current highest
```

---

## 4.7 Outbid Another Player

Iron Wolves places a higher bid.

```text
POST /auctions/1/bids
{
    "guild_id": 3,
    "amount": 1050
}
```

Expected result:

* bid_id = 2
* highest_bid = 1050
* highest_bidder_id = 3

Check Shadow Guild wallet:

```
reserved_money=0
```

Check Iron Wolves wallet:

```
reserved_money=1050
```

---

## 4.8 Cancel the Highest Bid (Should Fail)

```text
DELETE /auctions/1/bids/2
{
    "guild_id": 3
}
```

Expected result:

403 error

```
cannot cancel bid while you are the highest bidder
```

---

## 4.9 Cancel a Losing Bid

```text
DELETE /auctions/1/bids/1
{
    "guild_id": 2
}
```

Expected result:

```
204 No Content
```

The reserved balance should already be released.

---

## 4.10 Cancel Someone Else's Bid (Should Fail)

```text
DELETE /auctions/1/bids/2
{
    "guild_id": 2
}
```

Expected result:

403 error

```
you do not own this bid
```

---

# 5. Closing an Auction

The scheduler checks expired auctions every 30 seconds.

To test this without waiting 24 hours, insert an expired auction directly into the database.

```bash
docker compose exec postgres psql -U postgres -d aethoria -c "
INSERT INTO auctions (
    item_id, seller_id, starting_price,
    highest_bid, highest_bidder_id,
    end_time, original_end_time, status
) VALUES (
    7, 2, 500,
    600, 3,
    NOW() - INTERVAL '1 minute',
    NOW() - INTERVAL '1 minute',
    'active'
);"
```

Wait up to 30 seconds.

```text
GET /auctions/2
```

Expected result:

```
status=finished
```

```text
GET /items/7
```

Expected result:

```
owner_id=3
status=available
```

Check both wallets and transaction history.

The buyer should pay 600, the seller should receive 600, and both should have the correct transaction records.

---

## 5.1 Auction With No Bids

Insert another expired auction.

```bash
docker compose exec postgres psql -U postgres -d aethoria -c "
INSERT INTO auctions (
    item_id, seller_id, starting_price,
    end_time, original_end_time, status
) VALUES (
    5, 3, 999,
    NOW() - INTERVAL '1 minute',
    NOW() - INTERVAL '1 minute',
    'active'
);"
```

Wait up to 30 seconds.

```text
GET /auctions/3
```

Expected result:

```
status=cancelled
```

```text
GET /items/5
```

Expected result:

```
status=available
```

The item should be returned to the seller.

---

# 6. Concurrency Test

Run two bid requests at the same time.

```bash
curl -s -X POST http://localhost:8080/api/v1/auctions/1/bids \
  -H "Content-Type: application/json" \
  -d '{"guild_id": 2, "amount": 1200}' &

curl -s -X POST http://localhost:8080/api/v1/auctions/1/bids \
  -H "Content-Type: application/json" \
  -d '{"guild_id": 3, "amount": 1200}' &

wait
```

Expected result:

Only one request should succeed.
The other request should return an error.
Both wallets should remain in a valid state.

---

# 7. Unit Tests

```bash
go test ./internal/domain/auction/... -v
go test ./internal/domain/guild/... -v
go test ./internal/infrastructure/adapters/oracle/... -v
```

Expected result:

All tests should pass.
