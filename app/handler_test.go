package app_test

import (
	"github.com/Fajar-Islami/ais_code_test/controller"
	"github.com/Fajar-Islami/ais_code_test/infrastructure/container"
	"github.com/Fajar-Islami/ais_code_test/infrastructure/db"
	"github.com/Fajar-Islami/ais_code_test/repository"
	"github.com/Fajar-Islami/ais_code_test/service"
	"gorm.io/gorm"
)

func createTestArticleApp() (controller.ArticleController, *gorm.DB) {
	var (
		cont                                           = container.New("../.env")
		pgsqlDb           *gorm.DB                     = db.SetupDatabaseConnection(*cont.Pgsql)
		articleRepo       repository.ArticleRepository = repository.NewArticleRepository(pgsqlDb)
		artilceService    service.ArticleService       = service.NewArticleService(articleRepo)
		articleController controller.ArticleController = controller.NewArticleController(artilceService)
	)

	return articleController, pgsqlDb
}
