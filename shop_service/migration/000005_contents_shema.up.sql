CREATE TABLE contents (
    id BIGSERIAL PRIMARY KEY,
    article_id BIGINT NOT NULL REFERENCES articles(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    title_tag VARCHAR(100) NOT NULL,
    paragraph TEXT[],
    paragraph_tag TEXT[],
    new_line BOOLEAN[],
    img_url TEXT NOT NULL,
    "order" SMALLINT NOT NULL
);

CREATE INDEX contents_article_id_idx
    ON contents (article_id);

CREATE INDEX contents_title_trgm_idx
    ON contents USING GIN (title gin_trgm_ops);

CREATE INDEX contents_paragraph_gin_idx
    ON contents USING GIN (paragraph);

CREATE INDEX contents_paragraph_tag_gin_idx
    ON contents USING GIN (paragraph_tag);