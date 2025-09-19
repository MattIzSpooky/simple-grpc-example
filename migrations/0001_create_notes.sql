-- Creates the notes table
CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE notes (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       description TEXT NOT NULL,
                       created TIMESTAMPTZ NOT NULL DEFAULT now(),
                       updated TIMESTAMPTZ
);