CREATE TABLE IF NOT EXISTS posttbl(
     id BIGSERIAL PRIMARY KEY,
     Title TEXT NOT NULL,
     Subtitle VARCHAR(250) NOT NULL,
     Body TEXT NOT NULL,
     AuthorID INT NOT NULL REFERENCES usertbl(id),
     created_at timestamp with time zone default current_timestamp,
     updated_at timestamp with time zone null
);