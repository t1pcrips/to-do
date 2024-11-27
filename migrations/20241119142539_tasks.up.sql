create table tasks (
    id serial primary key,
    title VARCHAR(255) not null,
    is_done boolean default false,
    created_at timestamp not null default now(),
    updated_at timestamp not null default now(),
    deleted_at timestamp default null
);