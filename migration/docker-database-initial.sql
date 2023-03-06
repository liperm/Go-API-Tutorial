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
    code varchar unique not null,
    active boolean default true,
    quantity integer,
    value float default 0.0
);

create table go_api.customer_item(
    id serial primary key,
    customer_id int not null,
    item_id int not null,
    active boolean default true,
    
    constraint fk_customer
        foreign key (customer_id)
            references go_api.customer(id),
    constraint fk_item
        foreign key (item_id)
            references go_api.item(id)
);