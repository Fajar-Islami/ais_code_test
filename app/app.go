package app

import (
	"fmt"
	"net/http"

	"github.com/Fajar-Islami/ais_code_test/controller"
	"github.com/Fajar-Islami/ais_code_test/infrastructure/container"
	"github.com/Fajar-Islami/ais_code_test/infrastructure/db"
	redisClient "github.com/Fajar-Islami/ais_code_test/infrastructure/redis"
	"github.com/Fajar-Islami/ais_code_test/repository"
	"github.com/Fajar-Islami/ais_code_test/service"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func Start(cont *container.Container, server *gin.Engine) {
	var (
		redisDb           *redis.Client                = redisClient.NewRedisClient(*cont.Redis)
		pgsqlDb           *gorm.DB                     = db.SetupDatabaseConnection(*cont.Pgsql)
		articleRepo       repository.ArticleRepository = repository.NewArticleRepository(pgsqlDb)
		artilceService    service.ArticleService       = service.NewArticleService(articleRepo)
		articleController controller.ArticleController = controller.NewArticleController(artilceService)
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
