CREATE TABLE users (
  id TEXT PRIMARY KEY
);

-- CREATE TYPE DEVICE_TYPE AS ENUM ('chrome', 'android');

CREATE TABLE devices (
  public_key BYTEA PRIMARY KEY,

  name       TEXT NOT NULL,
  fcm_token  TEXT NOT NULL,

  user_id    TEXT NOT NULL REFERENCES users
);

-- CREATE TABLE invites (
--   id        TEXT PRIMARY KEY,
--
--   expires   TIMESTAMP NOT NULL,
--
--   source_id BYTEA REFERENCES devices
-- );

CREATE TABLE sigchain (
  link      BYTEA NOT NULL,

  signature BYTEA NOT NULL,
  signed_by BYTEA NOT NULL REFERENCES devices,

  user_id   TEXT  NOT NULL REFERENCES users,
  seq_no    INT   NOT NULL,

  UNIQUE (user_id, seq_no)
);
