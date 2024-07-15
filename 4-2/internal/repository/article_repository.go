package repository

import (
	"context"
	"fmt"

	"github.com/u-masato/blogger/4-2/internal/domain"
)

var articlesMap = make(map[int]*domain.Article)

type ArticleRepository struct{}

func NewArticleRepository() *ArticleRepository {
	return &ArticleRepository{}
}

func (r *ArticleRepository) Get(ctx context.Context, id domain.ArticleID) (*domain.Article, error) {
	article, ok := articlesMap[int(id)]
	if !ok {
		return &domain.Article{}, fmt.Errorf("article not found by id: %d", id)
	}
	return article, nil
}

func (r *ArticleRepository) Create(ctx context.Context, article *domain.Article) error {
	articlesMap[int(article.ID)] = article
	return nil
}
