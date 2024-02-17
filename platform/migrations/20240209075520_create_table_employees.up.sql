CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE employees
(
    id            UUID                     DEFAULT uuid_generate_v4() PRIMARY KEY,
    created_at    TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at    TIMESTAMP NULL,
    email         VARCHAR(255) NOT NULL UNIQUE,
    first_name     VARCHAR(255) NOT NULL,
    last_name     VARCHAR(255) NOT NULL,
    position      VARCHAR(255) NOT NULL,
    salary        INTEGER      NOT NULL,
    department_id INTEGER      NOT NULL,
    password      varchar(255) NOT NULL
);