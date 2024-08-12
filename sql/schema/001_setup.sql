-- +goose UP

CREATE TABLE setup(
    id SERIAL PRIMARY KEY,
    completed BOOLEAN NOT NULL
);

INSERT INTO setup(id, completed)
VALUES(1, FALSE);

-- +goose Down

DROP TABLE setup;
