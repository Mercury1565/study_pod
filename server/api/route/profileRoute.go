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

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, dbPool *pgxpool.Pool, group *gin.RouterGroup) {
	userRepo := repository.NewUserRepo(dbPool)
	profileController := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(userRepo, timeout),
	}

	group.GET("/profile", profileController.FetchByID)
}
