DROP SCHEMA IF EXISTS sample;
CREATE SCHEMA sample;
USE sample;

DROP TABLE IF EXISTS customers;

CREATE TABLE customers
(
  id       INT(10),
  name     VARCHAR(40)
);

INSERT INTO customers (id, name) VALUES (1, "Ben White");
INSERT INTO customers (id, name) VALUES (1, "Pater Rice");
INSERT INTO customers (id, name) VALUES (1, "Bob Marry");