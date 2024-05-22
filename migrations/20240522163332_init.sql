-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "user"
(
    id bigint PRIMARY KEY NOT NULL,
    created_at timestamp(0) without time zone NOT NULL,
    updated_at timestamp(0) without time zone NOT NULL,
    login varchar NOT NULL,
    password varchar NOT NULL,
    CONSTRAINT user_unique UNIQUE (login)
);

CREATE TABLE IF NOT EXISTS "order"
(
    id bigint PRIMARY KEY NOT NULL,
    user_id bigint NOT NULL,
    `number` varchar NOT NULL,
    created_at timestamp(0) without time zone NOT NULL,
    CONSTRAINT fk_user_orders FOREIGN KEY(user_id) REFERENCES user(id),
    CONSTRAINT order_unique UNIQUE (`number`)
)

CREATE TABLE IF NOT EXISTS "user_balance"
(
    user_id bigint PRIMARY KEY NOT NULL,
    accrual double precision NOT NULL,
    withdrawal double precision NOT NULL
)

CREATE TYPE order_status AS ENUM (
    'NEW',
    'PROCESSING',
    'INVALID',
    'PROCESSED'
);
CREATE TABLE IF NOT EXISTS "accrual"
(
    order_id bigint PRIMARY KEY NOT NULL,
    status order_status NOT NULL DEFAULT 'NEW',
    `value` double precision NOT NULL,
    created_at timestamp(0) without time zone NOT NULL,
    updated_at timestamp(0) without time zone NOT NULL,
)

CREATE TABLE IF NOT EXISTS "withdrawal"
(
    id bigint PRIMARY KEY NOT NULL,
    order_id bigint NOT NULL,
    `value` double precision NOT NULL,
    created_at timestamp(0) without time zone NOT NULL,
    CONSTRAINT fk_order_withdrawals FOREIGN KEY(order_id) REFERENCES order(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.user;
DROP TABLE IF EXISTS public.order;
DROP TABLE IF EXISTS public.user_balance;
DROP TABLE IF EXISTS public.accrual;
DROP TABLE IF EXISTS public.withdrawal;
-- +goose StatementEnd