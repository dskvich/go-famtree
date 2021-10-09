create table users(
   id uuid primary key default uuid_generate_v4(),
   created_at timestamptz not null default now(),
   first_name varchar,
   last_name varchar
)