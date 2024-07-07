package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/u-masato/blogger/2-3/internal/domain"
)

func TestNewTitle(t *testing.T) {
	t.Run("Valid Title", func(t *testing.T) {
		title, err := domain.NewTitle("Valid Title")
		assert.NoError(t, err)
		assert.Equal(t, domain.Title("Valid Title"), title)
	})

	t.Run("Empty Title", func(t *testing.T) {
		title, err := domain.NewTitle("")
		assert.NoError(t, err)
		assert.Equal(t, domain.Title("No Title"), title)
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
