package app_test

import (
	"github.com/Fajar-Islami/ais_code_test/controller"
	"github.com/Fajar-Islami/ais_code_test/infrastructure/container"
	"github.com/Fajar-Islami/ais_code_test/infrastructure/db"
	redisClient "github.com/Fajar-Islami/ais_code_test/infrastructure/redis"
	"github.com/Fajar-Islami/ais_code_test/repository/article"
	"github.com/Fajar-Islami/ais_code_test/service"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func createTestArticleApp() (controller.ArticleController, *gorm.DB) {
	var (
		cont                  = container.New("../.env")
		pgsqlDb *gorm.DB      = db.SetupDatabaseConnection(*cont.Pgsql)
		redisDb *redis.Client = redisClient.NewRedisClient(*cont.Redis)

		articleRepo       article.ArticleRepository      = article.NewArticleRepository(pgsqlDb)
		articleRedisRepo  article.RedisArtilceRepository = article.NewRedisRepo(redisDb)
		artilceService    service.ArticleService         = service.NewArticleService(articleRepo, articleRedisRepo)
		articleController controller.ArticleController   = controller.NewArticleController(artilceService)
	)

	return articleController, pgsqlDb
}
