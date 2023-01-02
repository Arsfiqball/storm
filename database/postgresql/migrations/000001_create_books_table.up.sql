CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    series VARCHAR(255) NOT NULL,
    volume INT NOT NULL,
    file_url TEXT,
    cover_url TEXT,
    publish_date TIMESTAMP
);
