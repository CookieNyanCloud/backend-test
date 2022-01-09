CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS userbalance
(
    id      UUID PRIMARY KEY,
    balance DECIMAL CHECK (balance >= 0) DEFAULT 0
);

CREATE TABLE IF NOT EXISTS transactions
(
    user_id     UUID        NOT NULL REFERENCES userbalance (id) ON DELETE CASCADE,
    operation   varchar(16) NOT NULL,
    sum         DECIMAL     NOT NULL,
    date        timestamp   NOT NULL DEFAULT (now()),
    description varchar(20)          DEFAULT '',
    user_to     UUID REFERENCES userbalance (id)
--     idempotency_key UUID        NOT NULL UNIQUE
);
