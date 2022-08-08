package article

import (
	"context"
	"errors"

	"github.com/Fajar-Islami/ais_code_test/entity"
	"gorm.io/gorm"
)

type articleRepositoryImpl struct {
	psql    *gorm.DB
	context context.Context
}

type ArticleRepository interface {
	Insert(entity.Article) (entity.Article, error)
	GetAll(entity.Article) []entity.Article
	GetById(id uint64) entity.Article
}

func NewArticleRepository(context context.Context, dbConn *gorm.DB) ArticleRepository {
	return &articleRepositoryImpl{
		context: context,
		psql:    dbConn,
	}
}

func (a *articleRepositoryImpl) Insert(data entity.Article) (res entity.Article, err error) {

	if a.psql.Save(&data).Error != nil {
		return res, errors.New("Failed to insert data")
	}

	a.psql.Find(&data)
	res = data
	return
}
func (a *articleRepositoryImpl) GetAll(filterData entity.Article) (res []entity.Article) {
	base := a.psql

	if filterData.Author != "" {
		base = base.Where("author = ?", filterData.Author)
	}
	if filterData.Search != "" {
		likeString := "%" + filterData.Search + "%"
		base = base.Where("body like ? OR title like ? ", likeString, likeString)
	}

	if filterData.Limit != 0 {
		base = base.Limit(int(filterData.Limit))
	}

	if filterData.Page != 0 {

		offset := (filterData.Page - 1) * filterData.Limit
		base = base.Offset(int(offset))
	}

	base.Debug().Order("created DESC").Find(&res)
	return
}
func (a *articleRepositoryImpl) GetById(id uint64) (res entity.Article) {
	a.psql.Find(&res, id)
	return res
}
