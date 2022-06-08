CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id INT NOT NULL AUTO_INCREMENT primary key,
    nome VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL unique,
    senha VARCHAR(100) NOT NULL,
    nick varchar(50) NOT NULL unique,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=INNODB;