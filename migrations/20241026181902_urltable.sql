-- +goose Up
-- +goose StatementBegin
CREATE TABLE urls (
   id_url VARCHAR(8)  PRIMARY KEY,
   original_url TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE urls;
-- +goose StatementEnd
