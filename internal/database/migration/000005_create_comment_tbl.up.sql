CREATE TABLE IF NOT EXISTS commenttbl(
     ID BIGSERIAL PRIMARY KEY,
     PostID INT NOT NULL REFERENCES posttbl(id),
     AuthorID INT NOT NULL REFERENCES usertbl(id),
     Body TEXT NOT NULL,
     CreatedAt timestamp with time zone default current_timestamp(),
     UpdatedAt timestamp with time zone
);