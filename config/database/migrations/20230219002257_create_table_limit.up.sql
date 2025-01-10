CREATE TABLE limits(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    updated_at TIMESTAMP NULL ,
    deleted_at TIMESTAMP NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    tenor INT NOT NULL,
    limitations INT NOT NULL,
    CONSTRAINT FK_limits_user_id_users_id FOREIGN KEY (user_id)
    REFERENCES users(id)
);