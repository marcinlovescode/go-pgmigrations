-- +goose Up
CREATE TABLE IF NOT EXISTS sampletable(
	id text NOT NULL,
	CONSTRAINT sampletable_pkey PRIMARY KEY (id)
);

-- +goose Down
DROP TABLE sampletable;