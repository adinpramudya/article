package articles

import (
	"article/internal/enums"
	"article/internal/models/articles"
	"errors"
)

func(s *service)CreateArticle(article *articles.Article) error{
	if !enums.IsValidStatus(article.Status){
		return errors.New("invalid status")
	}
	return s.repository.Create(article)
}

func(s *service)GetAllArticles(limit, offset int) ([] articles.Article,error){
	return s.repository.FindAll(limit, offset)
}

func(s *service)GetArticleByID(id uint)(*articles.Article, error){
	return s.repository.FindByID(id)
}

func(s *service)UpdateArticle(id uint, updated *articles.Article) error{
		article, err := s.repository.FindByID(id)
	if err != nil {
		return err
	}

	article.Title = updated.Title
	article.Content = updated.Content
	article.Category = updated.Category
	article.Status = updated.Status

	if !enums.IsValidStatus(article.Status) {
		return errors.New("invalid status")
	}

	return s.repository.Update(article)
}

func (s *service) DeleteArticle(id uint) error {
	return s.repository.Delete(id)
}