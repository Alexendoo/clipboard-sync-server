CREATE TABLE users (
  id SERIAL PRIMARY KEY
);

CREATE TYPE DEVICE_TYPE AS ENUM ('chrome', 'android');

CREATE TABLE devices (
  id          UUID PRIMARY KEY,

  name        VARCHAR(80) NOT NULL,
  device_type DEVICE_TYPE NOT NULL,

  user_id     INT         NOT NULL REFERENCES users
);

CREATE TABLE sigchain (
  link            BYTEA NOT NULL,
  sequence_number INT   NOT NULL,
  user_id         INT   NOT NULL REFERENCES users,


  PRIMARY KEY (user_id, sequence_number)
);
