-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE city (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    population BIGINT NOT NULL DEFAULT 0
);

CREATE INDEX idx_city_id ON city(id);
CREATE UNIQUE INDEX idx_city_name ON city(name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_city_name;
DROP INDEX IF EXISTS idx_city_id;
DROP TABLE IF EXISTS city;
-- +goose StatementEnd
