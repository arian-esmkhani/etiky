CREATE TYPE user_role AS ENUM ('customer', 'admin');

CREATE TABLE users (
    id         BIGSERIAL PRIMARY KEY,
    number     VARCHAR(11) UNIQUE NOT NULL,
    password   TEXT NOT NULL,
    role       user_role NOT NULL DEFAULT 'customer',
    status     BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX idx_users_user_name ON users(user_name) WHERE deleted_at IS NULL;
CREATE INDEX idx_users_deleted_at ON users(deleted_at);