CREATE TABLE carts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGSERIAL NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    product_id BIGSERIAL NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    quantity INT NOT NULL CHECK (quantity > 0),
    price BIGINT NOT NULL CHECK (price > 0),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE(user_id, product_id)
);

CREATE INDEX idx_carts_created_at ON carts(created_at);