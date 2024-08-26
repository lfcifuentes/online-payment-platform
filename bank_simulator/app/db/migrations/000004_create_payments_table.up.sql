CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    bank_id BIGINT NOT NULL,
    payment_method_id BIGINT NOT NULL,
    receiver_id BIGINT NOT NULL,
    amount NUMERIC(15, 2) NOT NULL,
    status VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);