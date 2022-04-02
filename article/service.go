package article

type ServiceImpl struct {
	repo Repository
}

func ProvideService(repo Repository) Service {
	return &ServiceImpl{
		repo: repo,
	}
}

func (s ServiceImpl) GetLists(keyword string, author string) ([]Article, error) {
	lists, err := s.repo.GetList(keyword, author)
	return lists, err
}

func (s ServiceImpl) AddNew(article Article) (err error) {
	return s.repo.Add(article)
}
