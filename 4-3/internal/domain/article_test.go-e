package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/u-masato/blogger/4-2/internal/domain"
)

func TestNewTitle(t *testing.T) {
	t.Run("Valid Title", func(t *testing.T) {
		title, err := domain.NewTitle("Valid Title")
		assert.NoError(t, err)
		assert.Equal(t, domain.ArticleTitle("Valid Title"), title)
	})

	t.Run("Empty Title", func(t *testing.T) {
		title, err := domain.NewTitle("")
		assert.NoError(t, err)
		assert.Equal(t, domain.ArticleTitle("No Title"), title)
	})

	t.Run("Too Long Title", func(t *testing.T) {
		longTitle := string(make([]byte, domain.MaxTitleLength+1))
		_, err := domain.NewTitle(longTitle)
		assert.Error(t, err)
		assert.EqualError(t, err, "title length must be less than 255")
	})
}

func TestNewContent(t *testing.T) {
	t.Run("Valid Content", func(t *testing.T) {
		content, err := domain.NewContent("Valid Content")
		assert.NoError(t, err)
		assert.Equal(t, domain.Content("Valid Content"), content)
	})

	t.Run("Empty Content", func(t *testing.T) {
		_, err := domain.NewContent("")
		assert.Error(t, err)
		assert.EqualError(t, err, "content cannot be empty")
	})
}

func TestNewAuthor(t *testing.T) {
	t.Run("Valid Author", func(t *testing.T) {
		author, err := domain.NewAuthor("Valid Author")
		assert.NoError(t, err)
		assert.Equal(t, domain.Author("Valid Author"), author)
	})

	t.Run("Empty Author", func(t *testing.T) {
		_, err := domain.NewAuthor("")
		assert.Error(t, err)
		assert.EqualError(t, err, "author cannot be empty")
	})

	t.Run("Too Long Author", func(t *testing.T) {
		longAuthor := string(make([]byte, domain.MaxAuthorLength+1))
		_, err := domain.NewAuthor(longAuthor)
		assert.Error(t, err)
		assert.EqualError(t, err, "author length must be less than 64")
	})
}

func TestCreateArticle(t *testing.T) {
	successCases := []struct {
		id      uint
		name    string
		title   string
		content string
		author  string
	}{
		{
			id:      1,
			name:    "Valid Article",
			title:   "Valid Title",
			content: "Valid Content",
			author:  "Valid Author",
		},
	}

	errorCases := []struct {
		id      uint
		name    string
		title   string
		content string
		author  string
		errMsg  string
	}{
		{
			id:      3,
			name:    "Empty Content",
			title:   "Valid Title",
			content: "",
			author:  "Valid Author",
			errMsg:  "content cannot be empty",
		},
		{
			id:      4,
			name:    "Empty Author",
			title:   "Valid Title",
			content: "Valid Content",
			author:  "",
			errMsg:  "author cannot be empty",
		},
		{
			id:      5,
			name:    "Too Long Title",
			title:   string(make([]byte, domain.MaxTitleLength+1)),
			content: "Valid Content",
			author:  "Valid Author",
			errMsg:  "title length must be less than 255",
		},
	}

	t.Run("Success Cases", func(t *testing.T) {
		for _, tt := range successCases {
			t.Run(tt.name, func(t *testing.T) {
				article, err := domain.CreateArticle(tt.id, tt.title, tt.content, tt.author)
				assert.NoError(t, err)
				assert.Equal(t, domain.ArticleID(tt.id), article.ID)
				expectedTitle := domain.ArticleTitle(tt.title)
				assert.Equal(t, expectedTitle, article.Title)
				assert.Equal(t, domain.Content(tt.content), article.Content)
				assert.Equal(t, domain.Author(tt.author), article.Author)

				// 時間が一致するかどうかはうまく確認できない
			})
		}
	})

	t.Run("Error Cases", func(t *testing.T) {
		for _, tt := range errorCases {
			t.Run(tt.name, func(t *testing.T) {
				_, err := domain.CreateArticle(tt.id, tt.title, tt.content, tt.author)
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errMsg)
			})
		}
	})
}
