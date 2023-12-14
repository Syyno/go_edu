-- +goose Up
-- +goose StatementBegin
alter table users_update_history
add column created_at timestamp not null default current_timestamp;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table users_update_history
drop column created_at;
-- +goose StatementEnd
