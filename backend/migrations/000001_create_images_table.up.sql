CREATE TABLE images
(
    id                SERIAL PRIMARY KEY,
    filepath          TEXT NOT NULL,
    upload_date       TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    modification_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);