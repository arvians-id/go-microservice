CREATE TABLE IF NOT EXISTS products (
    id SERIAL,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_by INTEGER,
    image VARCHAR(256),
    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (id)
)