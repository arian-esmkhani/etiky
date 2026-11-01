package models

type Content struct {
	ID               uint64     `db:"id"`
	ArticleID        uint64     `db:"article_id"`
	Title            string     `db:"title"`
	TitleTag         string     `db:"title_tag"`
	Paragraph        []string   `db:"paragraph"`
	ParagraphTag     []string   `db:"paragraph_tag"`
	NewLine          []bool     `db:"new_line"`
	ImgUrl           string     `db:"img_url"`
	Order			 uint8      `db:"order"`
}