package route

import (
	"Clean_Architecture/api/controller"
	"Clean_Architecture/bootstrap"
	"Clean_Architecture/repository"
	"Clean_Architecture/usecase"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/gin-gonic/gin"
)

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, dbPool *pgxpool.Pool, group *gin.RouterGroup) {
	userRepo := repository.NewUserRepo(dbPool)
	loginController := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(userRepo, timeout),
		Env:          env,
	}

	group.POST("/login", loginController.Login)
}
