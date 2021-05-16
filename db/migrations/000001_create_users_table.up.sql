CREATE TABLE IF NOT EXISTS users
(
    id           serial PRIMARY KEY,
    name         VARCHAR(50) UNIQUE NOT NULL,
    email        VARCHAR(50) UNIQUE NOT NULL,
    phone_number VARCHAR(20) UNIQUE NOT NULL
);
