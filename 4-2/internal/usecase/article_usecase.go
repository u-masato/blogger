package usecase

import (
	"context"

	"github.com/u-masato/blogger/4-2/internal/domain"
)

type IArticleRepository interface {
	Get(ctx context.Context, id domain.ArticleID) (*domain.Article, error)
	Create(ctx context.Context, article *domain.Article) error
}

type Usecase struct {
	articleRepo IArticleRepository
}

func NewArticleUsecase(articleRepo IArticleRepository) *Usecase {
	return &Usecase{
		articleRepo: articleRepo,
	}
}

func (uc *Usecase) GetArticleByID(ctx context.Context, id domain.ArticleID) (*domain.Article, error) {
	a, err := uc.articleRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return a, nil
}

var articleNextID = 1

func (uc *Usecase) CreateArticle(ctx context.Context, title, content, author string) error {
	article, err := domain.CreateArticle(
		uint(articleNextID),
		title,
		content,
		author,
	)
	if err != nil {
		return err
	}
	articleNextID++
	return uc.articleRepo.Create(ctx, article)
}
