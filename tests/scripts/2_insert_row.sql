-- +goose Up
-- +goose StatementBegin
INSERT INTO sampletable VALUES ('1337');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM sampletable WHERE Id = '1337';
-- +goose StatementEnd
