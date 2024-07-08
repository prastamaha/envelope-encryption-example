CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(25) UNIQUE,
    encrypted_name BYTEA NOT NULL,
    encrypted_gender BYTEA NOT NULL,
    encrypted_phone BYTEA NOT NULL,
    encrypted_address BYTEA NOT NULL,
    encrypted_dek BYTEA NOT NULL,
    consented BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX users_username_idx ON users (username);
