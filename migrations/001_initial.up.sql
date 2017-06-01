CREATE TABLE users (
  id TEXT PRIMARY KEY
);

-- CREATE TYPE DEVICE_TYPE AS ENUM ('chrome', 'android');

CREATE TABLE devices (
  id        TEXT PRIMARY KEY,

  name      TEXT NOT NULL,
  fcm_token TEXT NOT NULL,

  user_id   TEXT NOT NULL REFERENCES users
);

CREATE TABLE invites (
  id        TEXT PRIMARY KEY,

  expires   TIMESTAMP NOT NULL,

  source_id TEXT REFERENCES devices
);

CREATE TABLE sigchain (
  link            BYTEA NOT NULL,

  user_id         TEXT  NOT NULL REFERENCES users,
  sequence_number INT   NOT NULL,

  UNIQUE (user_id, sequence_number)
);
