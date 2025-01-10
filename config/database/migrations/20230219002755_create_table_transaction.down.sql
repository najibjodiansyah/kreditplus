ALTER TABLE transactions DROP FOREIGN KEY FK_user_id;
ALTER TABLE transactions DROP INDEX idx_uid;
ALTER TABLE transactions DROP INDEX idx_contract_number;
ALTER TABLE transactions DROP INDEX idx_aset_name;
DROP TABLE IF EXISTS transactions;