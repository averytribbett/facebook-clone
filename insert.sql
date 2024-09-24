-- Inserting data into users
INSERT INTO users (id, username, first_name, last_name, bio) VALUES
(0, 'cbeckers', 'Cade', 'Beckers', 'this is a test bio.'),
(1, 'mbrown', 'Melissa', 'Brown', 'Another test bio.');

-- Inserting data into posts
INSERT INTO posts (post_id, user_id, post_text) VALUES
(1, 0, 'this is a test post.');

-- Inserting data into reactions
INSERT INTO reactions (post_id, user_id, reaction) VALUES
(1, 0, 'thumbs up');

-- Inserting data into replies
INSERT INTO replies (post_id, user_id, reply_text) VALUES
(1, 0, 'this is a test reply.');

-- Inserting data into friends
INSERT INTO friends (user_id, friend_id) VALUES
(0, 1);
