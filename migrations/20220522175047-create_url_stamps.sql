-- +migrate Up
CREATE TABLE domain_stamps (
    id uuid NOT NULL DEFAULT uuid_generate_v4 (),
    domain_id uuid NOT NULL,
    stamp_id uuid NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    CONSTRAINT domain_stamps_pkey PRIMARY KEY (id),
    CONSTRAINT domain_stamps_domain_id_stamp_id_unique UNIQUE (domain_id, stamp_id)
);

CREATE INDEX idx_domain_stamps_fk_domain_id ON domain_stamps USING btree (domain_id);

CREATE INDEX idx_domain_stamps_fk_stamp_id ON domain_stamps USING btree (stamp_id);

ALTER TABLE domain_stamps
    ADD CONSTRAINT domain_stamps_domain_id_fkey FOREIGN KEY (domain_id) REFERENCES domains (id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE domain_stamps
    ADD CONSTRAINT domain_stamps_stamp_id_fkey FOREIGN KEY (stamp_id) REFERENCES stamps (id) ON UPDATE CASCADE ON DELETE CASCADE;

-- +migrate Down
DROP TABLE IF EXISTS domain_stamps;

