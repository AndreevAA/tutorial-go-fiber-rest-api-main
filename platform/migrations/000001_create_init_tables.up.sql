-- Add UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET TIMEZONE="Europe/Moscow";

-- Create books table
CREATE TABLE books (
    id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
    updated_at TIMESTAMP NULL,
    title VARCHAR (255) NOT NULL,
    author VARCHAR (255) NOT NULL,
    book_status INT NOT NULL,
    book_attrs JSONB NOT NULL
);

create table users (
  id uuid default uuid_generate_v4 ()  ,
  login text not null primary key ,
  password text,
  lvl_access int
);

-- Add indexes
CREATE INDEX active_books ON books (title) WHERE book_status = 1;

insert into books(title, author, book_status, book_attrs) values ('asda', 'asd', 1, '{}');

migrate -database ${POSTGRESQL_URL='postgres://postgres:qwerty@localhost:5432/example?sslmode=disable'} -path db/migrations up;
migrate -database ${POSTGRESQL_URL='postgres://postgres:qwerty@localhost:5432/example?sslmode=disable'} -path db/migrations up psql example -c "\d users"
