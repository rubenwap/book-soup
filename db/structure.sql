create table authors (
  id uuid primary key,
  name text not null,
  url text not null
);

create table books (
  id uuid primary key,
  url text not null,
  author_id uuid

);

ALTER TABLE books 
ADD FOREIGN KEY (author_id) REFERENCES authors(id);

create table book_record (
    id uuid primary key,
    name text not null,
    description text not null
);

ALTER TABLE book_record
ADD FOREIGN KEY (id) REFERENCES books(id);
