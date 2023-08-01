-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS account (
    id uuid primary key,
    telegram text null,
    email text null
);

-- +goose Down
-- +goose StatementBegin
DROP TABLE account;
