CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    surname VARCHAR(20) NOT NULL,
    patronymic VARCHAR(20) NOT NULL,
    age INT NOT NULL,
    gender VARCHAR(10) NOT NULL,
    nationality VARCHAR(20) NOT NULL
);