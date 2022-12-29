CREATE DATABASE IF NOT EXISTS img_db;
USE img_db;

CREATE TABLE IF NOT EXISTS images (
    id int not null AUTO_INCREMENT,
    name VARCHAR(60) NOT NULL,
    uuid VARCHAR(40) NOT NULL,
    lat FLOAT(10, 7) NOT NULL,
    lon FLOAT(10, 7) NOT NULL,
    primary key(id)
);

INSERT INTO images(name, uuid, lat, lon)
VALUES
  ("Go.png", "08a097d7-83db-4e0c-84bc-7730e14b0aaa", 23.098678, 90.235898),
  ("Go.png", "08a097d7-83db-4e0c-84bc-7730e14b0aaa", 24.098678, 90.235898),
  ("Go.png", "08a097d7-83db-4e0c-84bc-7730e14b0aaa", 25.098678, 90.235898),
  ("Go.png", "08a097d7-83db-4e0c-84bc-7730e14b0aaa", 26.098678, 90.235898),
  ("Go.png", "08a097d7-83db-4e0c-84bc-7730e14b0aaa", 23.098678, 91.235898),
  ("Go.png", "08a097d7-83db-4e0c-84bc-7730e14b0aaa", 23.098678, 92.235898),
  ("Go.png", "08a097d7-83db-4e0c-84bc-7730e14b0aaa", 24.098678, 92.235898),
  ("Go.png", "08a097d7-83db-4e0c-84bc-7730e14b0aaa", 25.098678, 92.235898),
  ("Go.png", "08a097d7-83db-4e0c-84bc-7730e14b0aaa", 23.098678, 93.235898),
("Go.png", "08a097d7-83db-4e0c-84bc-7730e14b0aaa", 90.0, -180.0),
  ("Go.png", "08a097d7-83db-4e0c-84bc-7730e14b0aaa", 90.0, 180.0),
  ("Go.png", "08a097d7-83db-4e0c-84bc-7730e14b0aaa", -90.0, -180.0),
  ("Go.png", "08a097d7-83db-4e0c-84bc-7730e14b0aaa", -90.0, 180.0),
  ("Go.png", "08a097d7-83db-4e0c-84bc-7730e14b0aaa", 0.0, -180.0),
  ("Go.png", "08a097d7-83db-4e0c-84bc-7730e14b0aaa", 0.0, 180.0),
  ("Go.png", "08a097d7-83db-4e0c-84bc-7730e14b0aaa", 90.0, 0.0),
  ("Go.png", "08a097d7-83db-4e0c-84bc-7730e14b0aaa", -90.0, 0.0),
  ("Go.png", "08a097d7-83db-4e0c-84bc-7730e14b0aaa", 0.0, 0.0);


CREATE USER IF NOT EXISTS 'img_user'@'%' IDENTIFIED BY '12345678';
GRANT ALL PRIVILEGES ON *.* TO 'img_user'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;