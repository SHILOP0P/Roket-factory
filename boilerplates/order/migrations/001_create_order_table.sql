-- +goose Up
create table orders (
    order_uuid uuid primary key,
    user_uuid uuid not null,
    part_uuids uuid[] not null,
    total_price numeric(10, 2) not null,
    transaction_uuid uuid,
    payment_method integer,
    order_status varchar(20) not null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

-- +goose Down
DROP TABLE IF EXISTS orders;