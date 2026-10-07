CREATE EXTENSION IF NOT EXISTS pg_trgm;

-- Table
CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    image_url TEXT NOT NULL,
    mark TEXT NOT NULL,
    price BIGINT NOT NULL,
    discount SMALLINT DEFAULT 0,
    sales_count BIGINT DEFAULT 0,
    count INT DEFAULT 0,
    images TEXT[],
    features TEXT[],
    tags TEXT[],
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

-- -- Gin
CREATE INDEX idx_products_tags ON products USING GIN (tags);
CREATE INDEX idx_products_features ON products USING GIN (features);
CREATE INDEX idx_products_name_trgm ON products USING GIN (name gin_trgm_ops);
CREATE INDEX idx_products_price ON products (price) WHERE deleted_at IS NULL AND count > 0;