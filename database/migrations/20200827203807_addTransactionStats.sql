-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE transactions
  ADD status boolean DEFAULT false;


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE transactions
DROP COLUMN status;