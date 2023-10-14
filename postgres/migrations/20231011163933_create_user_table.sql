-- +goose Up
create table users
(
    id               serial primary key,
    name             varchar   not null,
    email            varchar   not null,
    role             int       not null,
    password         varchar   not null,
    password_confirm varchar   not null,
    created_at       timestamp not null default now(),
    updated_at       timestamp
);

-- +goose Down
drop table users;
