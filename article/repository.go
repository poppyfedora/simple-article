package article

import (
	"github.com/jinzhu/gorm"
	"time"
)

type RepositoryImpl struct {
	db *gorm.DB
}

func ProvideRepository(db *gorm.DB) Repository {
	return &RepositoryImpl{
		db: db,
	}
}

func (r RepositoryImpl) GetList(keyword string, author string) (lists []Article, err error) {
	err = r.db.Select("*").Order("ID DESC").
		Or("title LIKE ?", "%"+keyword+"%").
		Or("body LIKE ?", "%"+keyword+"%").
		Or("author LIKE ?", "%"+author+"%").
		Take(&lists).Error
	return
}

func (r RepositoryImpl) Add(article Article) (err error) {
	article.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return r.db.Create(article).Error
}
