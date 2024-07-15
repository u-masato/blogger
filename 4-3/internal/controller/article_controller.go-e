package controller

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/u-masato/blogger/4-2/internal/domain"
)

type ArticleController struct {
	articleUC IArticleUseCase
}

type IArticleUseCase interface {
	GetArticleByID(ctx context.Context, id domain.ArticleID) (*domain.Article, error)
	CreateArticle(ctx context.Context, title, content, author string) error
}

func NewArticleController(au IArticleUseCase) *ArticleController {
	return &ArticleController{
		articleUC: au,
	}
}

func (ctrl *ArticleController) GetAllArticles(c *gin.Context) {}

func (ctrl *ArticleController) CreateArticle(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	author := c.PostForm("author")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title cannot be empty"})
		return
	}
	if content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "content cannot be empty"})
		return
	}
	if author == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "author cannot be empty"})
		return
	}
	// gin.Context から context.Context を取得
	ctx := c.Request.Context()
	if err := ctrl.articleUC.CreateArticle(ctx, title, content, author); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (ctrl *ArticleController) GetArticle(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	// gin.Context から context.Context を取得
	ctx := c.Request.Context()
	article, err := ctrl.articleUC.GetArticleByID(ctx, domain.ArticleID(articleID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if article == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, article)
}

func (ctrl *ArticleController) UpdateArticle(c *gin.Context) {}

func (ctrl *ArticleController) DeleteArticle(c *gin.Context) {}
