CREATE TABLE IF NOT EXISTS usertbl (
     ID BIGSERIAL PRIMARY KEY,
     RoleID INT NOT NULL REFERENCES roletbl(id) default 3,
     Username TEXT NOT NULL UNIQUE,
     Password TEXT NOT NULL,
     Email TEXT NOT NULL,
     created_at timestamp with time zone default current_timestamp,
     updated_at timestamp with time zone null
);