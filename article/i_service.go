package article

type Service interface {
	GetLists(keyword string, author string) ([]Article, error)
	AddNew(article Article) (err error)
}
