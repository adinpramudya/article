package articles

import "article/internal/models/articles"

func (r *repository) Create(article *articles.Article) error{
	return r.db.Create(article).Error
}

func (r *repository) FindAll(limit, offset int) ([]articles.Article, error){
	var articles []articles.Article
	err:= r.db.Limit(limit).Offset(offset).Find(&articles).Error
	return articles, err
}

func (r *repository) FindByID(id uint) (*articles.Article, error) {
	var article articles.Article
	err := r.db.First(&article, id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

func (r *repository) Update(article *articles.Article) error{
	return r.db.Save(article).Error
}

func (r *repository) Delete(id uint) error{
	return r.db.Delete(&articles.Article{},id).Error
}
