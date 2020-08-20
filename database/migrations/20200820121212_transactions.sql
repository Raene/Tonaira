-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE transactions (
    id int NOT NULL,
    accountNumber varchar(255) NOT NULL,
    bank varchar(255) NOT NULL,
    sender varchar(255),
    senderEmail varchar(255),
    amount int,
    currency varchar(255),
    coin    int,
    crypto  varchar(255),
    createdAt TIMESTAMP NOT NULL DEFAULT now(),
    PRIMARY KEY (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS transactions;
