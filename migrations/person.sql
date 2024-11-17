create table Person (
    id serial primary key,
    name varchar(255) not null,
    email varchar(255) not null unique,
    age int not null
)