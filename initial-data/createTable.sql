CREATE USER 'products'@'%' IDENTIFIED BY 'products';
GRANT ALL PRIVILEGES ON * . * TO 'products'@'%';

CREATE TABLE products (
  id varchar(25),
  name  varchar(25),
  description varchar(255)
);