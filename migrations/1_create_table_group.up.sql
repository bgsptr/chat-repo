CREATE TABLE groups (
  id bigserial PRIMARY KEY,
  group_name varchar NOT NULL,
  created_at timestamptz NOT NULL DEFAULT now(),
  owner_username varchar NOT NULL,
  CONSTRAINT owner_username_unique UNIQUE (owner_username)
);

CREATE TABLE follow_groups (
  username varchar NOT NULL,
  id_group bigserial NOT NULL,
  PRIMARY KEY (username, id_group),
  CONSTRAINT fk_group_id FOREIGN KEY (id_group) REFERENCES groups(id)
);

CREATE INDEX idx_follow_groups_username ON follow_groups (username);

-- CREATE TABLE "transfers" (
--   "id" bigserial PRIMARY KEY,
--   "from_account_id" bigint NOT NULL,
--   "to_account_id" bigint NOT NULL,
--   "amount" bigint NOT NULL,
--   "created_at" timestamptz NOT NULL DEFAULT (now())
-- );