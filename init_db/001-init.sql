CREATE DATABASE books_db DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci;

CREATE USER 'dev'@'%' IDENTIFIED BY 'dev';

GRANT ALL PRIVILEGES ON books_db.* TO 'dev'@'%';

USE books_db;
CREATE TABLE books (
  id int NOT NULL AUTO_INCREMENT,
  title text NOT NULL,
  author text NOT NULL,
  released_year year NOT NULL,
  PRIMARY KEY (id)
);