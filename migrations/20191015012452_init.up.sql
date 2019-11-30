CREATE TABLE "users" 
(
    id uuid PRIMARY KEY,
    name varchar(64) NOT NULL,
    email varchar(254) NOT NULL UNIQUE,
    password varchar(256) NOT NULL,
    created_at timestamp with time zone NOT NULL
);

CREATE TABLE "sessions"
(
    id char(64) PRIMARY KEY,
    user_id uuid REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    expires_at timestamp with time zone NOT NULL
);

CREATE TABLE "posts"
(
    id uuid PRIMARY KEY,
    user_id uuid REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    title varchar(256) NOT NULL,
    data bytea NOT NULL,
    created_at timestamp with time zone NOT NULL
);

