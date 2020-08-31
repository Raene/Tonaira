-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE callback_payloads (
    id serial,
    address VARCHAR(255),
    value   FLOAT,
    tx_hash  VARCHAR(255),
    status  boolean DEFAULT  false,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    PRIMARY KEY (id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS callback_payloads;
