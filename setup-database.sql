CREATE DATABASE play;

CREATE TABLE play.players (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    score INT
);

INSERT INTO play.players (name, score) VALUES ("pq", 10);
INSERT INTO play.players (name, score) VALUES ("qp", 5);

CREATE DATABASE play_test;

CREATE TABLE play_test.players (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    score INT
);

INSERT INTO play_test.players (name, score) VALUES ("pq", 10);
