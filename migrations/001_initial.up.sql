CREATE TABLE users (
  id INT PRIMARY KEY
);

CREATE TYPE DEVICE_TYPE AS ENUM ('chrome', 'android');

CREATE TABLE devices (
  id          INT PRIMARY KEY,

  name        VARCHAR(80) NOT NULL,
  device_type DEVICE_TYPE NOT NULL,

  user_id     INT         NOT NULL REFERENCES users
);
