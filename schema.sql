CREATE TABLE IF NOT EXISTS authors (
    id serial primary key,
    name varchar not null,
    url varchar not null
);

CREATE TABLE IF NOT EXISTS books (
    id serial primary key,
    isbn varchar unique not null,
    title varchar not null,
    number_of_pages integer,
    cover_image varchar,
    subjects varchar[],
    publish_date date
);

CREATE TABLE IF NOT EXISTS authorships (
    id serial primary key,
    book_id integer references books(id) NOT NULL,
    author_id integer references authors(id) NOT NULL
);

