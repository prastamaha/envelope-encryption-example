CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(25) UNIQUE,
    name VARCHAR(255) NOT NULL,
    gender VARCHAR(10) NOT NULL,
    phone VARCHAR(25) NOT NULL,
    address VARCHAR(255) NOT NULL,
    consented BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX users_username_idx ON users (username);
