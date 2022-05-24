-- +migrate Up
CREATE TABLE domains (
    id uuid NOT NULL DEFAULT uuid_generate_v4 (),
    domain text NOT NULL,
    parent_id uuid,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    CONSTRAINT domains_pkey PRIMARY KEY (id),
    CONSTRAINT domains_domain_unique UNIQUE (DOMAIN)
);

CREATE INDEX idx_domains_parent_id_fk ON domains USING btree (parent_id);

ALTER TABLE domains
    ADD CONSTRAINT domains_parent_id_fkey FOREIGN KEY (parent_id) REFERENCES domains (id);

-- +migrate Down
DROP TABLE IF EXISTS domains;

