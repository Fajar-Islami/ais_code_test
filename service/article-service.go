package service

import (
	"github.com/Fajar-Islami/ais_code_test/dto"
	"github.com/Fajar-Islami/ais_code_test/entity"
	"github.com/Fajar-Islami/ais_code_test/repository"
)

type ArticleService interface {
	GetArticles(dto.GetArticleDTO) []entity.Article
	PostArticle(dto.CreateArticleDTO) (entity.Article, error)
}

type articleServiceImpl struct {
	articleRepo repository.ArticleRepository
}

func NewArticleService(articleRepo repository.ArticleRepository) ArticleService {
	return &articleServiceImpl{
		articleRepo: articleRepo,
	}
}

func (as *articleServiceImpl) GetArticles(data dto.GetArticleDTO) []entity.Article {
	dataFilter := entity.Article{
		Author: data.Author,
		Search: data.Search,
	}
	return as.articleRepo.GetAll(dataFilter)
}
func (as *articleServiceImpl) PostArticle(data dto.CreateArticleDTO) (entity.Article, error) {
	articleData := entity.Article{
		Author: data.Author,
		Title:  data.Title,
		Body:   data.Body,
	}

	return as.articleRepo.Insert(articleData)
}
