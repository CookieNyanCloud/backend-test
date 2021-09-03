CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS user_balance
(
    id      UUID NOT NULL PRIMARY KEY,
    balance decimal
);

CREATE TABLE IF NOT EXISTS transactions
(
    user_id   UUID        NOT NULL REFERENCES user_balance (id) ON DELETE CASCADE,
    operation varchar(16) NOT NULL,
    user_to   UUID
);


