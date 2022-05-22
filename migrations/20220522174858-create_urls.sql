-- +migrate Up
CREATE TABLE domains (
    id uuid NOT NULL DEFAULT uuid_generate_v4 (),
    domain text NOT NULL,
    parent_id uuid,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    CONSTRAINT domains_pkey PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE IF EXISTS domains;

