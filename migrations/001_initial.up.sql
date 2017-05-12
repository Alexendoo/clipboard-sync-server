CREATE TABLE users (
  id INT PRIMARY KEY
);

CREATE TABLE devices (
  id          INT PRIMARY KEY,

  name        VARCHAR(80) NOT NULL,
  device_type INT         NOT NULL,

  user_id     INT         NOT NULL REFERENCES users
);
