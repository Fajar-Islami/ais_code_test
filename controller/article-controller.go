package controller

import (
	"net/http"

	"github.com/Fajar-Islami/ais_code_test/dto"
	"github.com/Fajar-Islami/ais_code_test/helper"
	"github.com/Fajar-Islami/ais_code_test/service"
	"github.com/gin-gonic/gin"
)

type ArticleController interface {
	NewArticle(c *gin.Context)
	GetListArticle(c *gin.Context)
}

type articleController struct {
	articleService service.ArticleService
}

func NewArticleController(articleService service.ArticleService) ArticleController {
	return &articleController{
		articleService: articleService,
	}
}

func (ac *articleController) NewArticle(c *gin.Context) {
	var articleData dto.CreateArticleDTO
	errDTO := c.ShouldBind(&articleData)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Request failed", errDTO.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	newArticle, err := ac.articleService.PostArticle(articleData)
	if err != nil {
		response := helper.BuildErrorResponse("Created article failed", "Something wrong when Created article failed", helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := helper.BuildSuccessResponse("Article created", newArticle)
	c.JSON(http.StatusCreated, response)

}

func (ac *articleController) GetListArticle(c *gin.Context) {
	var filterData dto.GetArticleDTO
	errDTO := c.ShouldBind(&filterData)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Request failed", errDTO.Error(), helper.EmptyObj{})
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	listArticle := ac.articleService.GetArticles(filterData)
	response := helper.BuildSuccessResponse("Get List Articles Success", listArticle)
	c.JSON(http.StatusCreated, response)

}
