create table trees(
    id uuid default uuid_generate_v4() constraint trees_pk primary key,
    created_at timestamptz default current_timestamp,
    name varchar(100),
    description varchar(1000)
)