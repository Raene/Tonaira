-- +goose Up
-- SQL in this section is executed when the migration is applied.
ALTER TABLE transactions
  ADD naira FLOAT;


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE transactions
DROP COLUMN naira;