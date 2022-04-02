package article

type Article struct {
	ID        int64  `json:"id"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	CreatedAt string `json:"created_at"`
}

func (Article) TableName() string {
	return "articles"
}
