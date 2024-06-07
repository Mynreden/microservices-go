CREATE TABLE posts (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       title TEXT NOT NULL,
                       content TEXT NOT NULL,
                       userId UUID NOT NULL,
                       FOREIGN KEY (userId) REFERENCES users (id)
);
