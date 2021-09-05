-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
--     user_id   UUID        NOT NULL REFERENCES userbalance (id) ON DELETE CASCADE,

CREATE TABLE IF NOT EXISTS userbalance
(
    id      INT NOT NULL PRIMARY KEY,
    balance DECIMAL CHECK (balance >= 0)
);

CREATE TABLE IF NOT EXISTS transactions
(
    user_id     INT         NOT NULL REFERENCES userbalance (id) ON DELETE CASCADE,
    operation   varchar(16) NOT NULL,
    sum         DECIMAL     NOT NULL,
    date        timestamp   NOT NULL DEFAULT (now()),
    description varchar(50) DEFAULT '',
    user_to     INT
);

-- INSERT INTO userbalance (id, balance)
-- values (1, null),
--        (2, 100),
--        (3, 200);


