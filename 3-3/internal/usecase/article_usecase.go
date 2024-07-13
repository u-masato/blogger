package usecase

import (
	"context"

	"github.com/u-masato/blogger/3-3/internal/domain"
)

type IArticleRepository interface {
	Get(id domain.ArticleID) (*domain.Article, error)
	Create(article *domain.Article) error
}

type IArticlePresenter interface {
	Progress(percentage int)
	Complete()
	Present(a *domain.Article)
}

type Usecase struct {
	articleRepo      IArticleRepository
	articlePresenter IArticlePresenter
}

func NewArticleUsecase(articleRepo IArticleRepository, articlePre IArticlePresenter) *Usecase {
	return &Usecase{
		articleRepo:      articleRepo,
		articlePresenter: articlePre,
	}
}

func (uc *Usecase) GetArticleByID(ctx context.Context, id domain.ArticleID) (*domain.Article, error) {
	a, err := uc.articleRepo.Get(id)
	if err != nil {
		return nil, err
	}
	uc.articlePresenter.Present(a)
	return a, nil
}

var articleNextID = 1

func (uc *Usecase) CreateArticle(ctx context.Context, title, content, author string) error {
	uc.articlePresenter.Progress(10)
	article, err := domain.CreateArticle(
		uint(articleNextID),
		title,
		content,
		author,
	)
	if err != nil {
		return err
	}
	uc.articlePresenter.Progress(50)
	articleNextID++
	err = uc.articleRepo.Create(article)
	if err != nil {
		return err
	}
	uc.articlePresenter.Complete()
	return nil
}
