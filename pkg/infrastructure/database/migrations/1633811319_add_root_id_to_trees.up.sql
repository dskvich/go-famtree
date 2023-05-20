-- +migrate Up
alter table trees add column root_id uuid references people(id) on delete cascade;