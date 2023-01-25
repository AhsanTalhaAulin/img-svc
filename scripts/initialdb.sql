CREATE DATABASE IF NOT EXISTS img_db;
USE img_db;

CREATE TABLE IF NOT EXISTS images (
    id int not null AUTO_INCREMENT,
    name VARCHAR(60) NOT NULL,
    uuid VARCHAR(40) NOT NULL,
    lat FLOAT(10, 7) NOT NULL,
    lon FLOAT(10, 7) NOT NULL,
    created_at timestamp NOT null,
    primary key(id)
);




CREATE USER IF NOT EXISTS 'img_user'@'%' IDENTIFIED BY '12345678';
GRANT ALL PRIVILEGES ON *.* TO 'img_user'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;