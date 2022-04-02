//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=article

package article

type Repository interface {
	GetList(keyword string, author string) (lists []Article, err error)
	Add(article Article) (err error)
}
