CREATE DATABASE IF NOT EXISTS capstone;
USE capstone;

-- Table: users
CREATE TABLE users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    bio TEXT
);

-- Table: posts
CREATE TABLE posts (
    post_id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    post_text TEXT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Table: reactions
CREATE TABLE reactions (
    post_id INT,
    user_id INT,
    reaction VARCHAR(255),
    FOREIGN KEY (post_id) REFERENCES posts(post_id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Table: replies
CREATE TABLE replies (
    post_id INT,
    user_id INT,
    reply_text TEXT,
    FOREIGN KEY (post_id) REFERENCES posts(post_id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- Table: friends
CREATE TABLE friends (
    user_id INT,
    friend_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (friend_id) REFERENCES users(id)
);
