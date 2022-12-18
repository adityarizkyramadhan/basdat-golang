CREATE DATABASE basdat_golang

USE basdat_golang

CREATE TABLE pelanggan(
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL
);

CREATE TABLE bus(
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    quota INT NOT NULL,
    destination VARCHAR(255),
    derpature VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL
);

CREATE TABLE book(
    id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    id_user INT,
    id_bus INT,
    created_at DATETIME NOT NULL,
    FOREIGN KEY (id_user) REFERENCES pelanggan(id),
    FOREIGN KEY (id_bus) REFERENCES bus(id)
)
