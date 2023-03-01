-- Active: 1661154899678@@127.0.0.1@3306

SHOW DATABASEs;

CREATE DATABASE nostra;

USE nostra;

CREATE Table
    product (
        id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        price INT NOT NULL,
        quantity INT NOT NULL DEFAULT 0,
        created_at BIGINT NOT NULL,
        updated_at BIGINT NOT NULL
    ) engine = InnoDB;

show tables;

ALTER Table product MODIFY updated_at BIGINT NULL;

DESC product;

INSERT INTO
    product(
        name,
        price,
        quantity,
        created_at
    )
VALUES (
        "Laptop Lenovo",
        9000000,
        10,
        1674574399076
    ), (
        "Laptop Asus",
        8000000,
        11,
        1674574399076
    );

SELECT * FROM product;

INSERT INTO
    product(
        name,
        price,
        quantity,
        created_at
    )
VALUES (
        "Laptop Lenovo",
        9000000,
        10,
        1674574399076
    ), (
        "Laptop Asus",
        8000000,
        11,
        1674574399076
    );

UPDATE product SET updated_at = 1674576973 WHERE id = 3 ;

SELECT * from product;

SELECT * from product;

UPDATE product
SET
    name = "Laptop Lenovo core i3",
    price = 900000,
    quantity = 15,
    updated_at = 1675009935427
WHERE id = 1;

UPDATE product SET name = "Laptop Lenovo inte core i3" WHERE id = 1;

SELECT * FROM product;

CREATE Table
    user (
        id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
        username VARCHAR(150) NOT NULL,
        password VARCHAR(200) NOT NULL
    ) engine = InnoDB;

INSERT INTO user(username, password) VALUES ("admin", "SecretAdmin");

SELECT * FROM user;

SELECT id, username W 