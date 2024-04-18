-- +goose Up
ALTER TABLE tasks
ALTER COLUMN id SET DEFAULT uuid_generate_v4()
;

-- +goose Down
ALTER TABLE tasks
ALTER COLUMN id DROP DEFAULT
;