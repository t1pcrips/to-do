-- +goose Up
-- +goose StatementBegin
alter table if exists tasks
add column if not exists user_id integer not null
references users(id) on delete cascade;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table if exists tasks
drop column if exists user_id
-- +goose StatementEnd
