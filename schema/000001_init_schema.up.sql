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
    date        timestamp   DEFAULT (now()),
    description varchar(20) DEFAULT '',
    user_to     UUID REFERENCES userbalance (id)
);
