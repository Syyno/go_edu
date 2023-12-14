-- +goose Up
-- +goose StatementBegin
alter table users
add column password varchar (50) not null
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table users
drop column password
-- +goose StatementEnd
