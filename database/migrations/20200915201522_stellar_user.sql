-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE  stellar_users (
  id serial,
  account_id varchar(255) NOT NULL,
  stellar_address varchar(255) NOT NULL,
  memo_type varchar(255)  NOT NULL,
  memo varchar(255)  NOT NULL,
  PRIMARY KEY (id)
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS stellar_users;