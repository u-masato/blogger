package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/u-masato/blogger/2-3/internal/domain"
)

type ArticleController struct{}

var articlesMap = make(map[int]*domain.Article)
var articleNextID = 1

func NewArticleController() *ArticleController {
	return &ArticleController{}
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

	articlesMap[int(article.ID)] = article
	articleNextID++

	c.JSON(http.StatusCreated, article)
}

func (ctrl *ArticleController) GetArticle(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	article := articlesMap[articleID]
	if article == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, article)
}

func (ctrl *ArticleController) UpdateArticle(c *gin.Context) {}

func (ctrl *ArticleController) DeleteArticle(c *gin.Context) {}
