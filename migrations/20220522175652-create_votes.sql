-- +migrate Up
CREATE TABLE votes (
    id uuid NOT NULL DEFAULT uuid_generate_v4 (),
    domain_stamp_id uuid NOT NULL,
    user_id uuid NOT NULL,
    approved boolean NOT NULL,
    rating bigint,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    CONSTRAINT votes_pkey PRIMARY KEY (id),
    CONSTRAINT votes_user_id_domain_stamp_id_unique UNIQUE (domain_stamp_id, user_id)
);

CREATE INDEX idx_votes_fk_user_id ON votes USING btree (user_id);

CREATE INDEX idx_votes_fk_domain_stamp_id ON votes USING btree (domain_stamp_id);

ALTER TABLE votes
    ADD CONSTRAINT votes_user_id_fkey FOREIGN KEY (user_id) REFERENCES users (id) ON UPDATE CASCADE ON DELETE CASCADE;

ALTER TABLE votes
    ADD CONSTRAINT votes_domain_stamp_id_fkey FOREIGN KEY (domain_stamp_id) REFERENCES domain_stamps (id) ON UPDATE CASCADE ON DELETE CASCADE;

-- +migrate Down
DROP TABLE IF EXISTS votes;

