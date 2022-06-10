CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS seguidores;
DROP TABLE IF EXISTS usuarios;


CREATE TABLE usuarios(
    id INT NOT NULL AUTO_INCREMENT primary key,
    nome VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL unique,
    senha VARCHAR(100) NOT NULL,
    nick varchar(50) NOT NULL unique,
    criado_em TIMESTAMP DEFAULT CURRENT_TIMESTAMP
) ENGINE=INNODB;



CREATE TABLE seguidores (
    usuario_id INT NOT NULL, 
    FOREIGN KEY(usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    seguidor_id INT NOT NULL, 
    FOREIGN KEY(seguidor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    PRIMARY KEY(usuario_id, seguidor_id)
) ENGINE=INNODB;