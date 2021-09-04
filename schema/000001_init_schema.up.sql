-- CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
-- todo: uuid
-- todo: no init users
CREATE TABLE IF NOT EXISTS userbalance
(
--     id      UUID NOT NULL PRIMARY KEY,
    id      INT NOT NULL PRIMARY KEY,
    balance DECIMAL CHECK (balance > 0)
);

CREATE TABLE IF NOT EXISTS transactions
(
--     user_id   UUID        NOT NULL REFERENCES userbalance (id) ON DELETE CASCADE,
    user_id   INT         NOT NULL REFERENCES userbalance (id) ON DELETE CASCADE,
    operation varchar(16) NOT NULL,
    sum DECIMAL,
    user_to   INT
);

INSERT INTO userbalance (id, balance)
values (1, null),
       (2, 100),
       (3, 200);


