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
  ("Go.png", "c62eda74-4558-4fc3-8350-381628ee1899", 24.098678, 90.235898),
  ("Go.png", "b2a7a78b-5a3e-48a9-8e91-cf48af0e8d8c", 25.098678, 90.235898),
  ("Go.png", "02bb92a2-5ed5-44d7-9710-fa4750418328", 26.098678, 90.235898),
  ("Go.png", "a9a5df14-0270-475d-bdb4-ed3be5ca445a", 23.098678, 91.235898),
  ("Go.png", "cbe01b9f-52b6-4a55-bb62-703810e7918f", 23.098678, 92.235898),
  ("Go.png", "501dd279-a9bf-4594-9430-4257999c6bd0", 24.098678, 92.235898),
  ("Go.png", "e02b4f64-3b55-4af5-8b00-8009348455f8", 25.098678, 92.235898),
  ("Go.png", "1dfbce08-ebc3-49c5-88d4-0f3a9d158651", 23.098678, 93.235898),
  ("Go.png", "fd46f602-8679-4a65-ad2c-08b6a9ea52a1", 85.0, -180.0),
  ("Go.png", "8bc88b73-32df-4a28-867b-63413551a2cc", 85.0, 180.0),
  ("Go.png", "b175838c-2c1e-4930-b849-140adf212561", -85.0, -180.0),
  ("Go.png", "f946fe81-9df1-411f-8433-dc79515565f1", -85.0, 180.0),
  ("Go.png", "dd64d611-9cb6-46cf-969f-69ba8f20d96a", 0.0, -180.0),
  ("Go.png", "875d2e3c-b828-4f9a-869f-434bb8299772", 0.0, 180.0),
  ("Go.png", "bb271c3c-1799-4dad-b2e8-201f35eb93ee", 85.0, 0.0),
  ("Go.png", "385391e7-d746-4a0d-892c-1ee2d449ba80", -85.0, 0.0),
  ("Go.png", "786513b1-16cd-4f4b-bbc3-c8242775d3fe", 0.0, 0.0);


CREATE USER IF NOT EXISTS 'img_user'@'%' IDENTIFIED BY '12345678';
GRANT ALL PRIVILEGES ON *.* TO 'img_user'@'%' WITH GRANT OPTION;
FLUSH PRIVILEGES;