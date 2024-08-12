-- +goose UP

CREATE TABLE users(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    api_key VARCHAR(64) UNIQUE NOT NULL,
    usertype TEXT NOT NULL
);

-- +goose Down

DROP TABLE users;
