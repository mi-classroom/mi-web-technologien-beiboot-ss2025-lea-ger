CREATE TABLE folders (
                         id   SERIAL PRIMARY KEY,
                         name TEXT NOT NULL UNIQUE
);

ALTER TABLE images
    ADD COLUMN folder_id INTEGER REFERENCES folders(id) ON DELETE SET NULL;

INSERT INTO folders (id, name) VALUES (0, 'root');