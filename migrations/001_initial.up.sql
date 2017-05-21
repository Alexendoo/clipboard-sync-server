CREATE TABLE users (
  id UUID PRIMARY KEY
);

-- CREATE TYPE DEVICE_TYPE AS ENUM ('chrome', 'android');

CREATE TABLE devices (
  id        UUID PRIMARY KEY,

  name      TEXT NOT NULL,
  fcm_token TEXT NOT NULL,

  user_id   UUID NOT NULL REFERENCES users
);

-- CREATE TABLE sigchain (
--   id              UUID PRIMARY KEY,
--
--   link            BYTEA NOT NULL,
--   sequence_number INT   NOT NULL,
--   user_id         UUID  NOT NULL REFERENCES users
-- );
