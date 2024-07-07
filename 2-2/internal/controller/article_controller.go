package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleController struct{}

type Article struct {
	ID      int
	Title   string
	Content string
	Author  string
}

var articlesMap = make(map[int]*Article)
var articleNextID = 1

func NewArticleController() *ArticleController {
	return &ArticleController{}
}

func (ctrl *ArticleController) GetAllArticles(c *gin.Context) {}

func (ctrl *ArticleController) CreateArticle(c *gin.Context) {
	var article Article

	article.ID = articleNextID
	article.Title = c.PostForm("title")
	article.Content = c.PostForm("content")
	article.Author = c.PostForm("author")
	articlesMap[article.ID] = &article
	articleNextID++

	c.JSON(http.StatusCreated, article)
}

func (ctrl *ArticleController) GetArticle(c *gin.Context) {
	id := c.Param("id")
	articleID, err := strconv.Atoi(id)
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
