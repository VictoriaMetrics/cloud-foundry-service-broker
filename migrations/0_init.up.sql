create table if not exists users
(
    id       serial primary key,
    name     text not null,
    email    text not null,
    password text not null
);