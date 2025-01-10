CREATE INDEX idx_nik ON users (nik);
ALTER TABLE users DROP INDEX idx_nik;
DROP TABLE IF EXISTS users;