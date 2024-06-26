-- uuid-ossp eklentisini etkinleştir
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- users tablosunu oluştur
CREATE TABLE users (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       username VARCHAR(255) UNIQUE NOT NULL,
                       email VARCHAR(255) UNIQUE NOT NULL,
                       password VARCHAR(255) NOT NULL
);

-- posts tablosunu oluştur
CREATE TABLE posts (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       title VARCHAR(255) NOT NULL,
                       content TEXT NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       user_id UUID NOT NULL,
                       FOREIGN KEY (user_id) REFERENCES users(id)
);

-- comments tablosunu oluştur
CREATE TABLE comments (
                          id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                          content TEXT NOT NULL,
                          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          user_id UUID NOT NULL,
                          post_id UUID NOT NULL,
                          FOREIGN KEY (user_id) REFERENCES users(id),
                          FOREIGN KEY (post_id) REFERENCES posts(id)
);

-- likes tablosunu oluştur
CREATE TABLE likes (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       user_id UUID NOT NULL,
                       post_id UUID NOT NULL,
                       FOREIGN KEY (user_id) REFERENCES users(id),
                       FOREIGN KEY (post_id) REFERENCES posts(id)
);


-- Veri ekliyelim

-- users tablosuna örnek veri ekleme
INSERT INTO users (username, email, password) VALUES
                                                  ('john_doe', 'john.doe@example.com', 'hashed_password_1'),
                                                  ('jane_doe', 'jane.doe@example.com', 'hashed_password_2');

-- posts tablosuna örnek veri ekleme
INSERT INTO posts (title, content, user_id) VALUES
                                                ('First Post', 'This is the content of the first post.', (SELECT id FROM users WHERE username = 'john_doe')),
                                                ('Second Post', 'This is the content of the second post.', (SELECT id FROM users WHERE username = 'jane_doe'));

-- comments tablosuna örnek veri ekleme
INSERT INTO comments (content, user_id, post_id) VALUES
                                                     ('Great post!', (SELECT id FROM users WHERE username = 'jane_doe'), (SELECT id FROM posts WHERE title = 'First Post')),
                                                     ('Thanks for sharing!', (SELECT id FROM users WHERE username = 'john_doe'), (SELECT id FROM posts WHERE title = 'Second Post'));

-- likes tablosuna örnek veri ekleme
INSERT INTO likes (user_id, post_id) VALUES
                                         ((SELECT id FROM users WHERE username = 'jane_doe'), (SELECT id FROM posts WHERE title = 'First Post')),
                                         ((SELECT id FROM users WHERE username = 'john_doe'), (SELECT id FROM posts WHERE title = 'Second Post'));


-- Sorgular

SELECT * FROM users;
SELECT * FROM posts;
SELECT * FROM comments;
SELECT * FROM likes;

SELECT id, username, email FROM users;

SELECT
    p.id AS post_id,
    p.title AS post_title,
    p.content AS post_content,
    c.id AS comment_id,
    c.content AS comment_content
FROM
    posts p
        LEFT JOIN
    comments c ON p.id = c.post_id
WHERE
    p.user_id = (SELECT id FROM users WHERE username = 'john_doe');


