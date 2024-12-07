-- +goose Up
-- +goose StatementBegin
create table if not exists users (
                                     id serial primary key,
                                     email text unique not null,
                                     password text not null,
                                     created_at timestamp not null default now(),
                                     updated_at timestamp default null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users;
-- +goose StatementEnd
