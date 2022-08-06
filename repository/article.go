package repository

import (
	"errors"

	"github.com/Fajar-Islami/ais_code_test/entity"
	"gorm.io/gorm"
)

type articleRepositoryImpl struct {
	psql *gorm.DB
}

type ArticleRepository interface {
	Insert(entity.Article) (entity.Article, error)
	GetAll(entity.Article) []entity.Article
	GetById(id uint64) entity.Article
}

func NewArticleRepository(dbConn *gorm.DB) ArticleRepository {
	return &articleRepositoryImpl{
		psql: dbConn,
	}
}

func (a *articleRepositoryImpl) Insert(data entity.Article) (res entity.Article, err error) {
	if a.psql.Debug().Save(&data).Error != nil {
		return res, errors.New("Failed to insert data")
	}

	a.psql.Find(&data)
	res = data
	return
}
func (a *articleRepositoryImpl) GetAll(filterData entity.Article) (res []entity.Article) {
	base := a.psql.Debug()

	if filterData.Author != "" {
		base = base.Where("author = ?", filterData.Author)
	}
	if filterData.Body != "" {
		base = base.Where("body like ? ", "%"+filterData.Body+"%")
	}
	if filterData.Title != "" {
		base = base.Where("title like ? ", "%"+filterData.Title+"%")
	}

	base.Find(&res)
	return
}
func (a *articleRepositoryImpl) GetById(id uint64) (res entity.Article) {
	a.psql.Find(&res, id)
	return res
}
