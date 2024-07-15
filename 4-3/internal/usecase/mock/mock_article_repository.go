package mock

import (
	"context"
	"errors"

	"github.com/u-masato/blogger/4-3/internal/domain"
)

type MockArticleRepository struct {
	Articles map[domain.ArticleID]*domain.Article
	Err      error
}

func NewMockArticleRepository() *MockArticleRepository {
	return &MockArticleRepository{
		Articles: make(map[domain.ArticleID]*domain.Article),
	}
}

func (m *MockArticleRepository) Get(_ context.Context, id domain.ArticleID) (*domain.Article, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	article, ok := m.Articles[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return article, nil
}

func (m *MockArticleRepository) Create(_ context.Context, article *domain.Article) error {
	if m.Err != nil {
		return m.Err
	}
	m.Articles[article.ID] = article
	return nil
}

func (m *MockArticleRepository) Update(ctx context.Context, article *domain.Article) error {
	if m.Err != nil {
		return m.Err
	}
	m.Articles[article.ID] = article
	return nil
}
