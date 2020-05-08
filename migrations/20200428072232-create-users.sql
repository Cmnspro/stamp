-- +migrate Up
CREATE TABLE users (
    id uuid NOT NULL DEFAULT uuid_generate_v4 (),
    username varchar(255),
    "password" text,
    is_active bool NOT NULL,
    scopes text[] NOT NULL,
    last_authenticated_at timestamptz,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_username_key UNIQUE (username)
);

-- +migrate Down
DROP TABLE IF EXISTS users;

