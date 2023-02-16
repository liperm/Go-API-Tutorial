create schema go_api;

create table go_api.customer(
    id serial primary key,
    name varchar,
    email varchar not null,
    password varchar not null
);