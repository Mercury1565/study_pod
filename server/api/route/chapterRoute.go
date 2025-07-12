package route

import (
	"Clean_Architecture/api/controller"
	"Clean_Architecture/bootstrap"
	"Clean_Architecture/repository"
	"Clean_Architecture/usecase"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewChapterRouter(env *bootstrap.Env, timeout time.Duration, dbPool *pgxpool.Pool, group *gin.RouterGroup) {
	ChapterRepo := repository.NewChapterRepo(dbPool)
	ChapterController := &controller.ChapterController{
		ChapterUseCase: usecase.NewChapterUsecase(ChapterRepo, timeout),
	}

	group.POST("/chapter", ChapterController.Create)
	group.GET("/chapter/:id", ChapterController.FetchByID)
}
