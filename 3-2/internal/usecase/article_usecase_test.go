package usecase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/u-masato/blogger/3-2/internal/domain"
	"github.com/u-masato/blogger/3-2/internal/usecase/mock"
)

func TestUsecase_GetArticleByID(t *testing.T) {
	mockRepo := mock.NewMockArticleRepository()
	uc := NewArticleUsecase(mockRepo)

	articleID := domain.ArticleID(1)
	expectedArticle := &domain.Article{
		ID:      articleID,
		Title:   "Title",
		Content: "Content",
		Author:  "Author",
	}
	mockRepo.Articles[articleID] = expectedArticle

	article, err := uc.GetArticleByID(context.Background(), articleID)
	assert.NoError(t, err)
	assert.Equal(t, expectedArticle, article)
}

func TestUsecase_GetArticleByID_NotFound(t *testing.T) {
	mockRepo := mock.NewMockArticleRepository()
	uc := NewArticleUsecase(mockRepo)

	articleID := domain.ArticleID(1)

	article, err := uc.GetArticleByID(context.Background(), articleID)
	assert.Error(t, err)
	assert.Nil(t, article)
}

func TestUsecase_CreateArticle(t *testing.T) {
	mockRepo := mock.NewMockArticleRepository()
	uc := NewArticleUsecase(mockRepo)

	title := "Title"
	content := "Content"
	author := "Author"
	articleID := uint(1)

	err := uc.CreateArticle(context.Background(), title, content, author)
	assert.NoError(t, err)

	createdArticle, err := uc.GetArticleByID(context.Background(), domain.ArticleID(articleID))
	expectedArticle, _ := domain.CreateArticle(articleID, title, content, author)
	assert.NoError(t, err)
	assert.Equal(t, expectedArticle.Title, createdArticle.Title)
	assert.Equal(t, expectedArticle.Content, createdArticle.Content)
	assert.Equal(t, expectedArticle.Author, createdArticle.Author)
}

func TestUsecase_CreateArticle_ValidationError(t *testing.T) {
	mockRepo := mock.NewMockArticleRepository()
	uc := NewArticleUsecase(mockRepo)

	err := uc.CreateArticle(context.Background(), "", "Content", "Author")
	assert.NoError(t, err, "when title is empty, it should not return error")

	err = uc.CreateArticle(context.Background(), "Title", "", "Author")
	assert.Error(t, err)

	err = uc.CreateArticle(context.Background(), "Title", "Content", "")
	assert.Error(t, err)
}
