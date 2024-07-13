package repository

import (
	"fmt"

	"github.com/u-masato/blogger/3-2/internal/domain"
)

var articlesMap = make(map[int]*domain.Article)

type ArticleRepository struct{}

func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{}
}

func (r *ArticleRepository) Get(id domain.ArticleID) (*domain.Article, error) {
	article, ok := articlesMap[int(id)]
	if !ok {
		return &domain.Article{}, fmt.Errorf("article not found by id: %d", id)
	}
	return article, nil
}

func (r *ArticleRepository) Create(article *domain.Article) error {
	articlesMap[int(article.ID)] = article
	return nil
}
