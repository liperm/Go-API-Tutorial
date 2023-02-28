create schema go_api;

create table go_api.customer(
    id serial primary key,
    name varchar,
    email varchar not null,
    password varchar not null,
    active boolean default true 
);

create table go_api.item(
    id serial primary key,
    name varchar not null,
    code varchar not null,
    active boolean default true,
    quantity integer
)