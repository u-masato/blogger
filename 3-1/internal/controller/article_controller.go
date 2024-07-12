package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/u-masato/blogger/3-1/internal/domain"
)

type ArticleController struct {
	articleRepository IArticleRepository
}

var articleNextID = 1

type IArticleRepository interface {
	Get(id domain.ArticleID) (*domain.Article, error)
	Create(article *domain.Article) error
}

func NewArticleController(articleRepository IArticleRepository) *ArticleController {
	return &ArticleController{
		articleRepository: articleRepository,
	}
}

func (ctrl *ArticleController) GetAllArticles(c *gin.Context) {}

func (ctrl *ArticleController) CreateArticle(c *gin.Context) {
	article, err := domain.CreateArticle(
		uint(articleNextID),
		c.PostForm("title"),
		c.PostForm("content"),
		c.PostForm("author"),
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctrl.articleRepository.Create(article)
	articleNextID++

	c.JSON(http.StatusCreated, article)
}

func (ctrl *ArticleController) GetArticle(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	article, err := ctrl.articleRepository.Get(domain.ArticleID(articleID))
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
