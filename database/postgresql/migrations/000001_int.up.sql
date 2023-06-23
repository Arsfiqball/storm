CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(24) NOT NULL,
    password_hash TEXT NOT NULL
);
