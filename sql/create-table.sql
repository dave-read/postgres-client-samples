DROP TABLE IF EXISTS inventory;
CREATE TABLE inventory (id serial PRIMARY KEY, name VARCHAR(50), quantity INTEGER);

INSERT INTO inventory (name, quantity) VALUES ('banana', 150);
INSERT INTO inventory (name, quantity) VALUES ('mango',  100);
INSERT INTO inventory (name, quantity) VALUES ('orange',  50);
