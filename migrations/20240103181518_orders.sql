-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    orders (
        track_number VARCHAR(255) NOT NULL UNIQUE,
        entry VARCHAR(255) NOT NULL,
        delivery_info jsonb,
        payment_info jsonb,
        locale VARCHAR(10) NOT NULL,
        internal_signature VARCHAR(255),
        customer_id VARCHAR(255) NOT NULL,
        delivery_service VARCHAR(255) NOT NULL,
        shardkey VARCHAR(10) NOT NULL,
        sm_id integer NOT NULL,
        date_created TIMESTAMPTZ NOT NULL,
        oof_shard VARCHAR(10) NOT NULL,
        order_uid VARCHAR(255) NOT NULL UNIQUE
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
-- +goose StatementEnd
