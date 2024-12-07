-- +goose Up
-- +goose StatementBegin
create table tasks (
                       id serial primary key,
                       title VARCHAR(255) not null,
                       is_done boolean default false,
                       created_at timestamp not null default now(),
                       updated_at timestamp default null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists tasks;
-- +goose StatementEnd
