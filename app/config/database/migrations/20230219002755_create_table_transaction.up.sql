CREATE TABLE transactions(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    updated_at TIMESTAMP NULL ,
    deleted_at TIMESTAMP NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    contract_number VARCHAR(255) NOT NULL,
    otr INT NOT NULL,
    admin_fee INT NOT NULL,
    instalment INT NOT NULL,
    bank_interest INT NOT NULL,
    asset_name VARCHAR(255) NOT NULL,
    message LONGTEXT NULL,
    CONSTRAINT FK_transactions_user_id_users_id FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX idx_contract_number ON transactions (contract_number);
CREATE INDEX idx_aset_name ON transactions (asset_name);