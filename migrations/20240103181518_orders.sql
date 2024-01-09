-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    orders (
        order_uid VARCHAR(255) NOT NULL UNIQUE,
        order_info jsonb
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
