ALTER TABLE limits DROP FOREIGN KEY FK_user_id;
ALTER TABLE limits DROP INDEX idx_uid;
DROP TABLE IF EXISTS limits;