-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    items (
        price real NOT NULL,
        rid VARCHAR(255) NOT NULL,
        name VARCHAR(255) NOT NULL,
        sale integer NOT NULL,
        size VARCHAR(50) NOT NULL,
        total_price real NOT NULL,
        nm_id integer NOT NULL,
        brand VARCHAR(255) NOT NULL,
        status integer NOT NULL,
        order_uid VARCHAR(255) REFERENCES orders (order_uid)
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS items;
-- +goose StatementEnd
