CREATE SCHEMA IF NOT EXISTS user_service;

CREATE TABLE user_service.users (
    id           uuid primary key      DEFAULT gen_random_uuid(),
    phone        varchar(50)  not null UNIQUE,
    full_name    varchar(255) not null,
    password     varchar(255) not null,
    address      TEXT,
    date_created timestamp    not null default now()
);