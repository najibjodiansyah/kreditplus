CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
    updated_at TIMESTAMP NULL ,
    deleted_at TIMESTAMP NULL,
    nik VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    full_name VARCHAR(255) NOT NULL,
    legal_name VARCHAR(255) NOT NULL,
    birth_date DATE NULL,
    birth_place VARCHAR(255) NULL,
    wages INT NOT NULL,
    photo_ktp VARCHAR(255) NULL,
    photo_selfie VARCHAR(255) NULL,
    CONSTRAINT UC_users UNIQUE (id,nik,password)
);

CREATE INDEX idx_nik ON users (nik);