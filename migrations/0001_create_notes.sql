-- Creates the notes table

-- Enable pgcrypto (for gen_random_uuid)
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Drop table if it exists
DROP TABLE IF EXISTS notes;

-- Create the notes table
CREATE TABLE notes (
                       id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       description TEXT NOT NULL,
                       created TIMESTAMPTZ NOT NULL DEFAULT now(),
                       updated TIMESTAMPTZ
);
