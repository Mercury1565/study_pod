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

func NewInstanceRouter(env *bootstrap.Env, timeout time.Duration, dbPool *pgxpool.Pool, group *gin.RouterGroup) {
	InstanceRepo := repository.NewInstanceRepo(dbPool)
	InstanceController := &controller.InstanceController{
		InstanceUseCase: usecase.NewInstanceUsecase(InstanceRepo, timeout),
	}

	group.POST("/instance", InstanceController.Create)
	group.GET("/instance/:id", InstanceController.FetchByID)
	group.GET("/instance/user", InstanceController.FetchByUserID)
}
