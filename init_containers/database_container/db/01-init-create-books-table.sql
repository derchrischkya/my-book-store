CREATE TABLE books (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    title VARCHAR ( 50 ) NOT NULL,
    author VARCHAR ( 50 ) NOT NULL,
    category VARCHAR ( 50 ) NOT NULL,
    isbn VARCHAR ( 50 ) UNIQUE NOT NULL
);