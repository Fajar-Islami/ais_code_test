package app_test

import (
	"context"

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
		cont                                             = container.New("../.env")
		ctx               context.Context                = context.Background()
		redisDb           *redis.Client                  = redisClient.NewRedisClient(*cont.Redis)
		pgsqlDb           *gorm.DB                       = db.SetupDatabaseConnection(*cont.Pgsql)
		articleRedisRepo  article.RedisArtilceRepository = article.NewRedisRepo(ctx, redisDb)
		articleRepo       article.ArticleRepository      = article.NewArticleRepository(ctx, pgsqlDb)
		artilceService    service.ArticleService         = service.NewArticleService(ctx, articleRepo, articleRedisRepo)
		articleController controller.ArticleController   = controller.NewArticleController(ctx, artilceService)
	)

	return articleController, pgsqlDb
}
