-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE transactions (
    id serial,
    account_number varchar(255) NOT NULL,
    bank varchar(255) NOT NULL,
    sender varchar(255),
    sender_email varchar(255),
    amount int NOT NULL,
    network  varchar(255) NOT NULL,
    address  varchar(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    PRIMARY KEY (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS transactions;
