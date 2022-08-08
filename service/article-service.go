package service

import (
	"context"
	"fmt"
	"log"

	"github.com/Fajar-Islami/ais_code_test/dto"
	"github.com/Fajar-Islami/ais_code_test/entity"
	"github.com/Fajar-Islami/ais_code_test/repository/article"
)

type ArticleService interface {
	GetArticles(dto.GetArticleDTO) []entity.Article
	PostArticle(dto.CreateArticleDTO) (entity.Article, error)
}

type articleServiceImpl struct {
	context          context.Context
	articleRepo      article.ArticleRepository
	artilceRedisRepo article.RedisArtilceRepository
}

const basePrefix = "article:"
const timeExpire = 1

func NewArticleService(context context.Context, articleRepo article.ArticleRepository, artilceRedisRepo article.RedisArtilceRepository) ArticleService {
	return &articleServiceImpl{
		context:          context,
		articleRepo:      articleRepo,
		artilceRedisRepo: artilceRedisRepo,
	}
}

func (as *articleServiceImpl) GetArticles(data dto.GetArticleDTO) (res []entity.Article) {
	dataFilter := entity.Article{
		Author: data.Author,
		Search: data.Search,
		Limit:  data.Limit,
		Page:   data.Page,
	}

	if dataFilter.Author != "" || dataFilter.Search != "" {
		// subprefix untuk keys di redis
		strKeys := ""
		if dataFilter.Author != "" {
			strKeys = strKeys + "author=" + dataFilter.Author + ":"
		}
		if dataFilter.Search != "" {
			strKeys = strKeys + "search=" + dataFilter.Search + ":"
		}
		if dataFilter.Limit > 0 {
			strKeys = fmt.Sprint(strKeys, "limit=", dataFilter.Limit, ":")
		}
		if dataFilter.Page > 0 {
			strKeys = fmt.Sprint(strKeys, "page=", dataFilter.Page, ":")
		}

		timeStr := fmt.Sprint(timeExpire, "m")

		strKeys = as.keyWithPrefix(strKeys, timeStr)

		// Mengecek apakah keys sudah ada di redis
		data, err := as.artilceRedisRepo.GetArticleByQueryCtx(strKeys)
		if err != nil {
			log.Println(err)
		}

		if data == nil {
			// kalau data tidak ada diredis maka akan mengambil dari postgre
			res = as.articleRepo.GetAll(dataFilter)

			// set hasil respon postgre ke reids
			if err := as.artilceRedisRepo.SetArticleCtx(strKeys, timeExpire, &res); err != nil {
				log.Println(err)
			}
		} else {
			// kalau data sudah ada diredis, langsung dikembalikan
			res = data
		}

		return

	} else {
		// mengambil data dari database
		res = as.articleRepo.GetAll(dataFilter)
		return res
	}

}
func (as *articleServiceImpl) PostArticle(data dto.CreateArticleDTO) (entity.Article, error) {
	articleData := entity.Article{
		Author: data.Author,
		Title:  data.Title,
		Body:   data.Body,
	}

	return as.articleRepo.Insert(articleData)
}

func (as *articleServiceImpl) keyWithPrefix(subprefix, time string) string {
	return fmt.Sprintf("%s:%s:%s", basePrefix, subprefix, time)
}
