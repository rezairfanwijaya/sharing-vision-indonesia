package article

type IService interface {
	Save()
	GetByID()
	Update()
	Delete()
}

type service struct {
	repoArticle IRepository
}

func NewService(repoArticle IRepository) *service {
	return &service{repoArticle}
}
