CREATE TABLE users
(
    id            serial PRIMARY KEY,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE segments
(

    id   serial PRIMARY KEY,
    name varchar(255) not null unique
);

CREATE TABLE users_segments
(
    id         serial PRIMARY KEY,
    user_id    int references users (id) on delete cascade    not null,
    segment_id int references segments (id) on delete cascade not null
);