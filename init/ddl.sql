DROP SCHEMA IF EXISTS h24s_04;

CREATE SCHEMA h24s_04;

USE h24s_04;

DROP TABLE IF EXISTS Genre,
Slide;

CREATE TABLE
  Genre (id CHAR(36) PRIMARY KEY, genrename TEXT NOT NULL);

CREATE TABLE
  Slide (
    id CHAR(36) PRIMARY KEY,
    dl_url TEXT NOT NULL,
    thumb_url TEXT,
    title TEXT NOT NULL,
    genre_id CHAR(36) NOT NULL,
    FOREIGN KEY (genre_id) REFERENCES Genre (id),
    posted_at DATETIME NOT NULL,
    description TEXT
  );
