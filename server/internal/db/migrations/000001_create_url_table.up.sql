-- Up migration
CREATE TABLE url (
    id serial PRIMARY KEY,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    original text NOT NULL,
    short text NOT NULL
);
