CREATE TABLE IF NOT EXISTS users (
    email varchar(64) primary key,
    username varchar(32),
    birthday varchar(10),
    sex varchar(6),
    bio varchar(100),
    "password" text,
    created_at timestamp,
    updated_at timestamp,
    UNIQUE(username)
);