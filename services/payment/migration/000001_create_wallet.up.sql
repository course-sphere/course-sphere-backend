CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE SCHEMA IF NOT EXISTS payment;

CREATE TABLE IF NOT EXISTS payment.wallets(
    user_id uuid PRIMARY KEY,
    amount bigint NOT NULL DEFAULT 0
);
