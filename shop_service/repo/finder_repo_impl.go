package repo

import (
	"context"
	"fiber-app/dto"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type finderRepo struct {
	db *pgxpool.Pool
}

func NewFinderRepo(db *pgxpool.Pool) IFinderRepo {
	return &finderRepo{db: db}
}

// product
func (r *finderRepo) FindCheapestProducts(ctx context.Context,
	limit, offset uint8) ([]dto.PreviewProductDto, error) {
	query := `
		SELECT 
			p.id,
			p.name,
			p.image_url,
			p.price,
			p.discount,
			(p.price * (100 - p.discount) / 100)::BIGINT AS final_price
		FROM products p
		WHERE p.deleted_at IS NULL AND p.count > 0
		ORDER BY final_price ASC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query cheapest products: %v", err)
	}
	defer rows.Close()

	products := make([]dto.PreviewProductDto, 0, limit)
	for rows.Next() {
		var p dto.PreviewProductDto
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.ImgUrl,
			&p.Price,
			&p.Discount,
			&p.FinalPrice,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan cheapest product row: %v", err)
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error for cheapest products: %v", err)
	}

	if len(products) < 1 {
		log.Warnf("No cheapest products found")
	}

	return products, nil
}

func (r *finderRepo) CountCheapestProducts(ctx context.Context) (uint64, error) {
	query := `
		SELECT count(*)
		FROM products p
		WHERE p.deleted_at IS NULL AND p.count > 0
	`

	var count uint64
	err := r.db.QueryRow(ctx, query).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count cheapest products: %v", err)
	}

	if count == 0 {
		log.Warnf("Total count of cheapest products is zero")
	}

	return count, nil
}

func (r *finderRepo) FindProductsByDiscount(ctx context.Context, category string,
	limit, offset uint8) ([]dto.PreviewProductDto, error) {
	query := `
		SELECT 
			p.id,
			p.name,
			p.image_url,
			p.price,
			p.discount,
			(p.price * (100 - p.discount) / 100)::BIGINT AS final_price
		FROM products p
		WHERE p.deleted_at IS NULL AND p.count > 0 AND p.discount > 0
	`

	args := []interface{}{}
	paramIdx := 1

	if category != "" {
		query += fmt.Sprintf(` AND (p.tags @> ARRAY[$%d] OR p.features @> ARRAY[$%d] OR p.mark = $%d)`, paramIdx, paramIdx, paramIdx)
		args = append(args, category)
		paramIdx++
	}

	query += fmt.Sprintf(` ORDER BY p.discount DESC LIMIT $%d OFFSET $%d`, paramIdx, paramIdx+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query discounted products (category: %s): %v", category, err)
	}
	defer rows.Close()

	products := make([]dto.PreviewProductDto, 0, limit)

	for rows.Next() {
		var p dto.PreviewProductDto
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.ImgUrl,
			&p.Price,
			&p.Discount,
			&p.FinalPrice,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan discounted product row (category: %s): %v", category, err)
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error for discounted products (category: %s): %v", category, err)
	}

	if len(products) < 1 {
		log.Warnf("No discounted products found (category: %s)", category)
	}

	return products, nil
}

func (r *finderRepo) CountProductsByDiscount(ctx context.Context, category string) (uint64, error) {
	query := `
		SELECT count(*)
		FROM products p
		WHERE p.deleted_at IS NULL 
			AND p.count > 0
			AND p.discount > 0
	`

	args := []interface{}{}
	paramIdx := 1

	if category != "" {
		query += fmt.Sprintf(` AND (p.tags @> ARRAY[$%d] OR p.features @> ARRAY[$%d] OR p.mark = $%d)`, paramIdx, paramIdx, paramIdx)
		args = append(args, category)
	}

	var count uint64
	err := r.db.QueryRow(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count discounted products (category: %s): %v", category, err)
	}

	if count == 0 {
		log.Warnf("Total count of discounted products is zero (category: %s)", category)
	}

	return count, nil
}

func (r *finderRepo) FindBestSellingProducts(ctx context.Context, category string,
	limit, offset uint8) ([]dto.PreviewProductDto, error) {

	query := `
		SELECT 
			p.id,
			p.name,
			p.image_url,
			p.price,
			p.discount,
			(p.price * (100 - p.discount) / 100)::BIGINT AS final_price
		FROM products p
		WHERE p.deleted_at IS NULL 
			AND p.count > 0
	`

	args := []interface{}{}
	paramIdx := 1

	if category != "" {
		query += fmt.Sprintf(` AND (p.tags @> ARRAY[$%d] OR p.features @> ARRAY[$%d] OR p.mark = $%d)`, paramIdx, paramIdx, paramIdx)
		args = append(args, category)
		paramIdx++
	}

	query += fmt.Sprintf(` ORDER BY p.sales_count DESC LIMIT $%d OFFSET $%d`, paramIdx, paramIdx+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query best selling products (category: %s): %v", category, err)
	}
	defer rows.Close()

	products := make([]dto.PreviewProductDto, 0, limit)

	for rows.Next() {
		var p dto.PreviewProductDto
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.ImgUrl,
			&p.Price,
			&p.Discount,
			&p.FinalPrice,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan best selling product row (category: %s): %v", category, err)
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error for best selling products (category: %s): %v", category, err)
	}

	if len(products) < 1 {
		log.Warnf("No best selling products found (category: %s)", category)
	}

	return products, nil
}

func (r *finderRepo) CountBestSellingProducts(ctx context.Context, category string) (uint64, error) {
	query := `
		SELECT count(*)
		FROM products p
		WHERE p.deleted_at IS NULL 
			AND p.count > 0
	`

	args := []interface{}{}
	paramIdx := 1

	if category != "" {
		query += fmt.Sprintf(` AND (p.tags @> ARRAY[$%d] OR p.features @> ARRAY[$%d] OR p.mark = $%d)`, paramIdx, paramIdx, paramIdx)
		args = append(args, category)
	}

	var count uint64
	err := r.db.QueryRow(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count best selling products (category: %s): %v", category, err)
	}

	if count == 0 {
		log.Warnf("Total count of best selling products is zero (category: %s)", category)
	}

	return count, nil
}

func (r *finderRepo) FindCheapestProductsByFilter(ctx context.Context,
	filter dto.ProductFilterDto, limit, offset uint8) ([]dto.PreviewProductDto, error) {

	query := `
		SELECT 
			p.id,
			p.name,
			p.image_url,
			p.price,
			p.discount,
			(p.price * (100 - p.discount) / 100)::BIGINT AS final_price
		FROM products p
		WHERE p.deleted_at IS NULL
		AND ($1 = '' OR p.mark = $1) AND ($2 = '' OR p.tags @> ARRAY[$2])
		AND ($3 = '' OR p.features @> ARRAY[$3]) AND p.count > 0
		ORDER BY final_price ASC
		LIMIT $4 OFFSET $5
	`

	rows, err := r.db.Query(ctx, query, filter.Mark, filter.Tag, filter.Feature, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query cheapest products by filter '%s': %v", filter, err)
	}
	defer rows.Close()

	products := make([]dto.PreviewProductDto, 0, limit)
	for rows.Next() {
		var p dto.PreviewProductDto
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.ImgUrl,
			&p.Price,
			&p.Discount,
			&p.FinalPrice,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan cheapest product by filter '%s' row: %v", filter, err)
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error for cheapest products by filter '%s': %v", filter, err)
	}

	if len(products) < 1 {
		log.Debugf("Cheapest product by filter len is under zero !")
	}

	return products, nil
}

func (r *finderRepo) CountCheapestProductsByFilter(ctx context.Context, filter dto.ProductFilterDto) (uint64, error) {
	query := `
		SELECT count(*)
		FROM products p
		WHERE p.deleted_at IS NULL 
			AND ($1 = '' OR p.mark = $1) AND ($2 = '' OR p.tags @> ARRAY[$2])
			AND ($3 = '' OR p.features @> ARRAY[$3]) AND p.count > 0
	`

	var count uint64
	err := r.db.QueryRow(ctx, query, filter.Mark, filter.Tag, filter.Feature).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count cheapest products by filter '%s': %v", filter, err)
	}

	if count == 0 {
		log.Warnf("Total count of cheapest products by filter is zero (filter: %s)", filter)
	}

	return count, nil
}

func (r *finderRepo) SearchProducts(ctx context.Context, queryStr, category string,
	lowPrice uint32, maxPrice uint32, limit, offset uint8) ([]dto.PreviewProductDto, error) {

	query := `
		SELECT 
			p.id,
			p.name,
			p.image_url,
			p.price,
			p.discount,
			(p.price * (100 - p.discount) / 100)::BIGINT AS final_price
		FROM products p
		WHERE p.deleted_at IS NULL 
			AND p.count > 0
			AND (
				p.name ILIKE $1 OR p.mark = $1
				OR p.tags @> ARRAY[$1]
				OR p.features @> ARRAY[$1]
			)
	`

	args := []interface{}{"%" + queryStr + "%"}
	paramIdx := 2

	if category != "" {
		query += fmt.Sprintf(" AND (p.tags @> ARRAY[$%d] OR p.features @> ARRAY[$%d])", paramIdx, paramIdx)
		args = append(args, category)
		paramIdx++
	}

	if lowPrice > 0 {
		query += fmt.Sprintf(" AND p.price >= $%d", paramIdx)
		args = append(args, lowPrice)
		paramIdx++
	}

	if maxPrice > 0 {
		query += fmt.Sprintf(" AND p.price <= $%d", paramIdx)
		args = append(args, maxPrice)
		paramIdx++
	}

	query += fmt.Sprintf(" ORDER BY p.sales_count DESC LIMIT $%d OFFSET $%d", paramIdx, paramIdx+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to search products (query: %s, price range: %d-%d): %v", queryStr, lowPrice, maxPrice, err)
	}
	defer rows.Close()

	products := make([]dto.PreviewProductDto, 0, limit)

	for rows.Next() {
		var p dto.PreviewProductDto
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.ImgUrl,
			&p.Price,
			&p.Discount,
			&p.FinalPrice,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan search result row (query: %s): %v", queryStr, err)
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		log.Errorf("rows iteration error for search results (query: %s): %v", queryStr, err)
		return nil, err
	}

	if len(products) < 1 {
		log.Warnf("no search results found for query: %s", queryStr)
	}

	return products, nil
}

func (r *finderRepo) GetSearchSuggestions(ctx context.Context, queryStr string) ([]string, error) {

	query := `
		SELECT 
			p.name
		FROM products p
		WHERE p.deleted_at IS NULL 
			AND p.count > 0
			AND (
				p.name ILIKE $1 OR p.mark = $1
				OR p.tags @> ARRAY[$1]
				OR p.features @> ARRAY[$1]
			)
    LIMIT 3
	`

	rows, err := r.db.Query(ctx, query, "%"+queryStr+"%")
	if err != nil {
		log.Errorf("failed to get search suggestions for '%s': %v", queryStr, err)
		return nil, err
	}
	defer rows.Close()

	suggestions := make([]string, 0, 3)

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			log.Errorf("failed to scan suggestion: %v", err)
			return nil, err
		}
		suggestions = append(suggestions, name)
	}

	return suggestions, nil
}

func (r *finderRepo) CountSearch(ctx context.Context, queryStr, category string,
	lowPrice uint32, maxPrice uint32) (uint64, error) {

	query := `
		SELECT count(*)
		FROM products p
		WHERE p.deleted_at IS NULL 
			AND p.count > 0
			AND (
				p.name ILIKE $1 OR p.mark = $1
				OR p.tags @> ARRAY[$1]
				OR p.features @> ARRAY[$1]
			)
	`

	args := []interface{}{"%" + queryStr + "%"}
	paramIdx := 2

	if category != "" {
		query += fmt.Sprintf(" AND (p.tags @> ARRAY[$%d] OR p.features @> ARRAY[$%d])", paramIdx, paramIdx)
		args = append(args, category)
		paramIdx++
	}

	if lowPrice > 0 {
		query += fmt.Sprintf(" AND p.price >= $%d", paramIdx)
		args = append(args, lowPrice)
		paramIdx++
	}

	if maxPrice > 0 {
		query += fmt.Sprintf(" AND p.price <= $%d", paramIdx)
		args = append(args, maxPrice)
		paramIdx++
	}

	var count uint64
	err := r.db.QueryRow(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count search results (query: %s, price range: %d-%d): %v", queryStr, lowPrice, maxPrice, err)
	}

	if count == 0 {
		log.Warnf("total count of search results is zero for query: %s", queryStr)
	}

	return count, nil
}

// article
// ===================== Queries =====================
const articleQueryTpl = `
	SELECT
		a.id, a.title, a.description, a.views_count, COALESCE(a.author_id, NULL), a.status,
		COALESCE(a.img_url, ''), COALESCE(a.video_url, ''), COALESCE(a.tags, ARRAY[]::text[]),
		a.slug, a.meta_title, a.meta_description, a.meta_robots,
		a.canonical_mode, a.canonical_url,
		a.breadcrumb_title, a.breadcrumb_parent, a.breadcrumb_url, a.schema_type,
		a.published_at, a.updated_at,
		c.id IS NOT NULL,
		COALESCE(c.title, ''),
		COALESCE(c.title_tag, ''),
		COALESCE(c.paragraph, ARRAY[]::text[]),
		COALESCE(c.paragraph_tag, ARRAY[]::text[]),
		COALESCE(c.new_line, ARRAY[]::boolean[]),
		COALESCE(c.img_url, ''),
		COALESCE(c."order", 0)
	FROM articles a
	LEFT JOIN contents c ON c.article_id = a.id
	WHERE a.deleted_at IS NULL AND %s
	ORDER BY c."order" ASC
`

var (
	articleByUrlQuery   = fmt.Sprintf(articleQueryTpl, "a.slug = $1")
	articleByIdQuery    = fmt.Sprintf(articleQueryTpl, "a.id = $1")
	articleByTitleQuery = fmt.Sprintf(articleQueryTpl, "a.title = $1")
	notPublishedQuery   = fmt.Sprintf(articleQueryTpl, "(a.status = 'draft' OR a.published_at IS NULL)")
)

// ===================== Find Methods =====================

func (r *finderRepo) FindArticleBySlug(ctx context.Context, url string) (*dto.ArticleDto, error) {
	rows, err := r.db.Query(ctx, articleByUrlQuery, url)
	if err != nil {
		return nil, fmt.Errorf("query article by url=%q: %w", url, err)
	}
	defer rows.Close()

	article, err := scanArticle(rows)
	if err != nil {
		return nil, fmt.Errorf("scan article by url=%q: %w", url, err)
	}
	return article, nil
}

func (r *finderRepo) FindArticleById(ctx context.Context, id uint64) (*dto.ArticleDto, error) {
	rows, err := r.db.Query(ctx, articleByIdQuery, id)
	if err != nil {
		return nil, fmt.Errorf("query article by id=%d: %w", id, err)
	}
	defer rows.Close()

	article, err := scanArticle(rows)
	if err != nil {
		return nil, fmt.Errorf("scan article by id=%d: %w", id, err)
	}
	return article, nil
}

func (r *finderRepo) FindArticleByTitle(ctx context.Context, title string) (*dto.ArticleDto, error) {
	rows, err := r.db.Query(ctx, articleByTitleQuery, title)
	if err != nil {
		return nil, fmt.Errorf("query article by title=%q: %w", title, err)
	}
	defer rows.Close()

	article, err := scanArticle(rows)
	if err != nil {
		return nil, fmt.Errorf("scan article by title=%q: %w", title, err)
	}
	return article, nil
}

// ===================== Helpers =====================
func scanArticle(rows pgx.Rows) (*dto.ArticleDto, error) {
	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return nil, fmt.Errorf("rows iteration error: %w", err)
		}
		return nil, fmt.Errorf("article not found")
	}

	article := &dto.ArticleDto{}

	var (
		hasContent bool
		content    dto.ContentDto
	)

	err := rows.Scan(
		&article.ID, &article.Title, &article.Description, &article.ViewsCount,
		&article.AuthorID, &article.Status, &article.ImgUrl, &article.VideoUrl, &article.Tags,
		&article.Slug, &article.MetaTitle, &article.MetaDescription, &article.MetaRobots,
		&article.CanonicalMode, &article.CanonicalUrl,
		&article.BreadcrumbTitle, &article.BreadcrumbParent, &article.BreadcrumbUrl, &article.SchemaType,
		&article.PublishedAt, &article.UpdatedAt,
		&hasContent,
		&content.ContentTitle, &content.TitleTag,
		&content.Paragraph, &content.ParagraphTag, &content.NewLine,
		&content.ContentImgUrl, &content.Order,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to scan first article row: %w", err)
	}

	if hasContent {
		article.Content = append(article.Content, content)
	}

	for rows.Next() {
		var (
			hc bool
			c  dto.ContentDto
		)

		err := rows.Scan(
			new(uint64), new(string), new(string), new(uint64),
			new(uint64), new(string), new(string), new(string), new([]string),
			new(string), new(string), new(string), new(string),
			new(string), new(string),
			new(string), new(string), new(string), new(string),
			new(interface{}), new(interface{}),
			&hc,
			&c.ContentTitle, &c.TitleTag,
			&c.Paragraph, &c.ParagraphTag, &c.NewLine,
			&c.ContentImgUrl, &c.Order,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan article content row: %w", err)
		}

		if hc {
			article.Content = append(article.Content, c)
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return article, nil
}
