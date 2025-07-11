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

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, dbPool *pgxpool.Pool, group *gin.RouterGroup) {
	userRepo := repository.NewUserRepo(dbPool)
	signupController := &controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(userRepo, timeout),
		Env:           env,
	}

	group.POST("/signup", signupController.Signup)
}
