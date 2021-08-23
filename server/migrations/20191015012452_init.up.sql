CREATE TABLE "users" 
(
    id uuid PRIMARY KEY,
    name varchar(64) NOT NULL,
    avatar varchar(64) NOT NULL,
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
    author_id uuid REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    title varchar(256) NOT NULL,
    art bytea NOT NULL,
    created_at timestamp with time zone NOT NULL
);

CREATE TABLE "likes"
(
    user_id uuid REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    post_id uuid REFERENCES posts(id) ON DELETE CASCADE NOT NULL,
    CONSTRAINT like_pkey PRIMARY KEY (user_id, post_id)
);

CREATE TABLE "follows"
(
    followed_id uuid REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    follower_id uuid REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    CONSTRAINT follow_pkey PRIMARY KEY (followed_id, follower_id),
    CONSTRAINT mustnt_follow_self CHECK (followed_id <> follower_id)
);
