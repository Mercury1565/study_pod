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

func NewPodcastRouter(env *bootstrap.Env, timeout time.Duration, dbPool *pgxpool.Pool, group *gin.RouterGroup) {
	PodcastRepo := repository.NewPodcastRepo(dbPool)
	PodcastController := &controller.PodcastController{
		PodcastUseCase: usecase.NewPodcastUsecase(PodcastRepo, timeout),
	}

	group.POST("/podcast", PodcastController.Create)
	group.GET("/podcast/:id", PodcastController.FetchByID)
}
