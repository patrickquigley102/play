CREATE DATABASE play;

CREATE TABLE play.players (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    score INT
);

INSERT INTO play.players (name, score) VALUES ("pq", 10);
INSERT INTO play.players (name, score) VALUES ("qp", 5);