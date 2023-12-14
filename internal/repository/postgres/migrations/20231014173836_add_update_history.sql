-- +goose Up
-- +goose StatementBegin
create table if not exists users_update_history(
    id bigserial primary key,
    user_id bigserial,
    email_old varchar(50) not null,
    email_new varchar(50) not null,
    name_old varchar(50) not null,
    name_new varchar(50) not null,
    role_old int,
    role_new int,
    constraint fk_user foreign key (user_id) references users(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users_update_history
-- +goose StatementEnd
