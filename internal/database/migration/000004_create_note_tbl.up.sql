CREATE TABLE IF NOT EXISTS notetbl(
     id serial primary key,
     title text not null,
     body text not null,
     created_at timestamp with time zone default current_timestamp,
     updated_at timestamp with time zone null
);