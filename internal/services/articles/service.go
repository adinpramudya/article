package articles

import (
	"article/internal/configs"
	"article/internal/models/articles"
)

type repository interface{
	Create(article *articles.Article) error
	FindAll(limit, offset int) ([]articles.Article, error)
	FindByID(id uint) (*articles.Article, error)
	Update(article *articles.Article) error
	Delete(id uint) error
}

type service struct{
	cfg			*configs.Config
	repository	repository
}

func NewService(cfg *configs.Config, repository repository) *service{
	return &service{
		cfg: cfg,
		repository: repository,
	}
}