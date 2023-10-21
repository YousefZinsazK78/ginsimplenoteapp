CREATE TABLE IF NOT EXISTS imagetbl(
     id SERIAL PRIMARY KEY,
     img_url TEXT NOT NULL,
     post_id INTEGER NOT NULL REFERENCES posttbl(id),
     user_id INTEGER NOT NULL REFERENCES usertbl(id)
);