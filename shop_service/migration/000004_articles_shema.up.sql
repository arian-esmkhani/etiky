CREATE TABLE articles (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    views_count BIGINT NOT NULL DEFAULT 0,
    author_id BIGINT NOT NULL,
    status VARCHAR(30) NOT NULL,
    img_url TEXT,
    video_url TEXT,
    tags TEXT[],
    -- SEO
    slug VARCHAR(255) NOT NULL,
    meta_title VARCHAR(70) NOT NULL,
    meta_description TEXT NOT NULL,
    meta_robots VARCHAR(100) NOT NULL,
    canonical_mode VARCHAR(50) NOT NULL,
    canonical_url TEXT NOT NULL,
    breadcrumb_title TEXT NOT NULL,
    breadcrumb_parent TEXT NOT NULL,
    breadcrumb_url TEXT NOT NULL,
    schema_type VARCHAR(100) NOT NULL,
    -- Open Graph
    og_title TEXT,
    og_description TEXT,
    og_image TEXT,
    -- Date
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    published_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE INDEX articles_tags_gin_idx
    ON articles USING GIN (tags);

CREATE INDEX articles_title_trgm_idx
    ON articles USING GIN (title gin_trgm_ops);

CREATE INDEX articles_description_trgm_idx
    ON articles USING GIN (description gin_trgm_ops);