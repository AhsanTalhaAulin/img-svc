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

INSERT INTO images(name, uuid, created_at, lat, lon)
VALUES
("Go.png",	"faa3c180-98e6-11ee-891f-0242ac190004", "2023-12-12 12:07:11",	25.0000000,	34.0000000),	
("Go.png",	"faa3c180-98e6-11ee-8920-0242ac190004", "2023-12-12 12:07:11",	28.0000000,	34.0000000),	
("Go.png",	"ad741f80-98e7-11ee-8921-0242ac190004", "2023-12-12 12:12:11",	25.0000000,	34.0000000),	
("Go.png",	"60447d80-98e8-11ee-8922-0242ac190004", "2023-12-12 12:17:11",	25.0000000,	34.0000000),	
("Go.png",	"60dd1400-98e8-11ee-8923-0242ac190004", "2023-12-12 12:17:12",	25.0000000,	34.0000000);



CREATE USER IF NOT EXISTS 'img_user'@'%' IDENTIFIED BY '12345678';
GRANT ALL PRIVILEGES ON *.* TO 'img_user'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;