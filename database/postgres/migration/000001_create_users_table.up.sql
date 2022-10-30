CREATE TABLE IF NOT EXISTS users (
    id SERIAL,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(256) NOT NULL,
    PRIMARY KEY (id)
)