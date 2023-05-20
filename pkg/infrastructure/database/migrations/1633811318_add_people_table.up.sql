-- +migrate Up
create table people (
    id uuid default uuid_generate_v4() constraint people_pk primary key,
    created_at timestamptz default current_timestamp,
    name varchar(100),
    father_id uuid,
    mother_id uuid,
    tree_id uuid,
    foreign key (father_id) references people (id),
    foreign key (mother_id) references people (id),
    foreign key (tree_id) references trees (id)
);