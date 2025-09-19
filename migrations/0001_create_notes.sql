-- Creates the notes table
CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE IF NOT EXISTS notes
(
    id
    UUID
    PRIMARY
    KEY,
    description
    TEXT
    NOT
    NULL,
    created
    TIMESTAMPTZ
    NOT
    NULL
    DEFAULT
    now
(
),
    updated TIMESTAMPTZ
    );