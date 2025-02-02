CREATE TABLE transactions
(
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    account_id VARCHAR(36) NOT NULL,
    amount NUMERIC NOT NULL,
    operation_type_id INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (account_id) REFERENCES accounts (id)
);