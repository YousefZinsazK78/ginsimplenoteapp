create table if not exists user_profile(
     id bigserial primary key,
     firstname text not null,
     lastname text not null,
     email text UNIQUE,
     gender text CHECK(gender IN ('male','female')) NOT NULL,
     created_at timestamp without time zone
);

create table if not exists youtube_account(
     id bigint primary key references user_profile(id),
     created_at timestamp without time zone
);

create table if not exists youtube_channel(
     id bigserial primary key,
     youtube_account_id bigint not null references youtube_account(id),
     channel_name text not null unique,
     created_at timestamp without time zone
);

create table if not exists channel_subscriber(
     youtube_account_id bigint references youtube_account(id),
     youtube_channel_id bigint references youtube_channel(id),
     created_at timestamp without time zone,
     primary key (youtube_account_id, youtube_channel_id)
);

INSERT INTO user_profile(firstname, lastname, email, gender, created_at) VALUES 
	('yousef','zinsaz','yz.1378@gmail.com', 'male', CURRENT_TIMESTAMP),
	('sina','kashani','sina.1378@gmail.com', 'male', CURRENT_TIMESTAMP),
	('artina','irani','artina.1378@gmail.com', 'female', CURRENT_TIMESTAMP),
	('samira','yavari','samira.1378@gmail.com', 'female', CURRENT_TIMESTAMP);


INSERT INTO youtube_account(id,created_at) VALUES 
(1,CURRENT_TIMESTAMP),
(2,CURRENT_TIMESTAMP),
(3,CURRENT_TIMESTAMP);

INSERT INTO youtube_account(id,youtube_account_id, channel_name, created_at) VALUES 
(1,1, 'mariamBeauty', CURRENT_TIMESTAMP),
(2,2, 'joeTeck', '2020-11-02 23:40:30.80343'),
(3,4,'alexTutorials',CURRENT_TIMESTAMP);

INSERT INTO channel_subscriber(youtube_account_id, youtube_channel_id, created_at) VALUES 
(1,2, '2020-11-25 23:23:34.23433'),
(1,3, '2020-11-25 21:22:30.11433'),
(2,1, '2020-11-15 20:20:30.53433'),
(2,2, '2020-11-15 23:05:10.32433');
