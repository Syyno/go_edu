-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id bigserial primary key,
    name varchar(50) not null,
    email varchar(50) not null,
    role int not null default 0,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp null
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
