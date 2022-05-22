-- +migrate Up
CREATE TABLE stamps (
    id uuid NOT NULL DEFAULT uuid_generate_v4 (),
    name text NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    CONSTRAINT stamps_pkey PRIMARY KEY (id)
);

-- +migrate Down
DROP TABLE IF EXISTS stamps;

