create table users (
    id uuid default uuid_generate_v4() constraint users_pk primary key,
    created_at timestamptz default current_timestamp,
    login varchar(100) not null,
    role varchar(32) not null,
    lang varchar(2) default 'en' not null,
    name varchar(100),
    email varchar(50),
    password_hash varchar(100)
)