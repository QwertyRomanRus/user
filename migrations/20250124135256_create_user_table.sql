-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user" (
    id serial,
    first_name VARCHAR(255) not null,
    last_name VARCHAR(255) not null,
    email VARCHAR(255) not null,
    age smallint,
    created_at timestamp default now(),
    updated_at timestamp default now(),

    primary key (id),
    unique (email)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "user"
-- +goose StatementEnd
