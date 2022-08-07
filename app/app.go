package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Fajar-Islami/ais_code_test/controller"
	"github.com/Fajar-Islami/ais_code_test/infrastructure/container"
	"github.com/Fajar-Islami/ais_code_test/infrastructure/db"
	"github.com/Fajar-Islami/ais_code_test/repository/article"

	redisClient "github.com/Fajar-Islami/ais_code_test/infrastructure/redis"
	"github.com/Fajar-Islami/ais_code_test/service"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func Start(cont *container.Container, server *gin.Engine) {
	var (
		ctx               context.Context                = context.Background()
		redisDb           *redis.Client                  = redisClient.NewRedisClient(*cont.Redis)
		pgsqlDb           *gorm.DB                       = db.SetupDatabaseConnection(*cont.Pgsql)
		articleRedisRepo  article.RedisArtilceRepository = article.NewRedisRepo(ctx, redisDb)
		articleRepo       article.ArticleRepository      = article.NewArticleRepository(ctx, pgsqlDb)
		artilceService    service.ArticleService         = service.NewArticleService(ctx, articleRepo, articleRedisRepo)
		articleController controller.ArticleController   = controller.NewArticleController(ctx, artilceService)
	)

	defer db.CloseDatabaseConnection(pgsqlDb)
	defer redisDb.Close()

	articleRouter := server.Group("api/article")
	{
		articleRouter.GET("/", articleController.GetListArticle)
		articleRouter.POST("/", articleController.NewArticle)
	}

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	port := ":" + fmt.Sprint(cont.Apps.Port)
	server.Run(port)
}
