package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct{}

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

var articlesMap = make(map[int]*Article)
var articleNextID = 1

func NewArticleController() *ArticleController {
	return &ArticleController{}
}

func (ctrl *ArticleController) GetAllArticles(c *gin.Context) {}

func (ctrl *ArticleController) CreateArticle(c *gin.Context) {
	var article Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article.ID = articleNextID
	articlesMap[article.ID] = &article
	articleNextID++

	c.JSON(http.StatusCreated, article)
}

func (ctrl *ArticleController) GetArticle(c *gin.Context) {}

func (ctrl *ArticleController) UpdateArticle(c *gin.Context) {}

func (ctrl *ArticleController) DeleteArticle(c *gin.Context) {}
