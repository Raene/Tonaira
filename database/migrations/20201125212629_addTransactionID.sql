-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE callback_payloads
  ADD transaction_id INT;


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE callback_payloads
DROP COLUMN transaction_id;