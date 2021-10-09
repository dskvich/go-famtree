create table trees(
  id uuid primary key default uuid_generate_v4(),
  created_at timestamptz not null default now(),
  name varchar,
  description varchar
)