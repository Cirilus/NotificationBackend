-- +goose Up
CREATE TABLE IF NOT EXISTS account (
    id uuid primary key,
    telegram text null,
    email text null
);

-- +goose Down
DROP TABLE account;
