-- +migrate Up
CREATE TABLE users (
    id serial primary key,
    name varchar(50),
    password text,
    email text
);

-- +migrate Down
DROP TABLE users;

-- +migrate Up
CREATE TABLE artists (
    id serial primary key,
    name varchar(50) UNIQUE,
    picture text,
    fans int
);

-- +migrate Down
DROP TABLE artists;

-- +migrate Up
CREATE TABLE albums (
    id serial primary key,
    title varchar(50) UNIQUE,
    upc text,
    cover text,
    artist_id int,
    CONSTRAINT fk_artist
        FOREIGN KEY(artist_id)
            REFERENCES artists(id)
);

-- +migrate Down
DROP TABLE albums;

-- +migrate Up
CREATE TABLE tracks (
    id serial primary key,
    title varchar(50) UNIQUE,
    duration int,
    rank int,
    album_id int,
    CONSTRAINT fk_album
        FOREIGN KEY(album_id)
            REFERENCES albums(id)
);

-- +migrate Down
DROP TABLE tracks;
