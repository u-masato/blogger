package usecase

import (
	"context"

	"github.com/u-masato/blogger/4-3/internal/domain"
)

type IArticleRepository interface {
	Get(ctx context.Context, id domain.ArticleID) (*domain.Article, error)
	Create(ctx context.Context, article *domain.Article) error
	Update(ctx context.Context, article *domain.Article) error
}

type ITxAdmin interface {
	Transaction(ctx context.Context, f func(ctx context.Context) error) error
}

type Usecase struct {
	articleRepo IArticleRepository
	tx          ITxAdmin
}

func NewArticleUsecase(articleRepo IArticleRepository, tx ITxAdmin) *Usecase {
	return &Usecase{
		articleRepo: articleRepo,
		tx:          tx,
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

func (uc *Usecase) UpdateArticle(ctx context.Context, id domain.ArticleID, title, content string) error {
	article, err := uc.articleRepo.Get(ctx, id)
	if err != nil {
		return err
	}
	err = article.Update(title, content)
	if err != nil {
		return err
	}
	updateFunc := func(ctx context.Context) error {
		return uc.articleRepo.Update(ctx, article)
	}
	return uc.tx.Transaction(ctx, updateFunc)
}
