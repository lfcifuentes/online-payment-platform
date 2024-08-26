CREATE TABLE payment_methods (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    last_4 VARCHAR(4) NOT NULL,
    brand VARCHAR(255) NOT NULL,
    exp_month INT NOT NULL,
    exp_year INT NOT NULL,
    bank_id INT NOT NULL,
    card_number VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);