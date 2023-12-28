-- Create a new database called 'DatabaseName'
CREATE TABLE users ( 
    id serial primary key,
    uuid varchar(64) not null unique,
    name varchar(255),
    email varchar(255) not null unique,
    password varchar(255) not null,
    created_at timestamp not null
);

-- Create a new database called 'sessions'
CREATE TABLE sessions (
    id serial primary key,
    uuid varchar(64) not null unique,
    email varchar(255),
    user_id integer references users(id),
    created_at timestamp not null
);

create table threads(
    id serial primary key,
    uuid varchar(64) not null unique,
    user_id integer references users(id),
    topic text,
    created_at timestamp not null
);

create table posts(
    id serial primary key,
    user_id integer references users(id),
    uuid varchar(64) not null unique,
    body text,
    thread_id integer references threads(id),
    created_at timestamp not null

);