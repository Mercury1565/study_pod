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

func NewBookRouter(env *bootstrap.Env, timeout time.Duration, dbPool *pgxpool.Pool, group *gin.RouterGroup) {
	BookRepo := repository.NewBookRepo(dbPool)
	BookController := &controller.BookController{
		BookUseCase: usecase.NewBookUsecase(BookRepo, timeout),
	}

	group.POST("/book", BookController.Create)
	group.GET("/book/:id", BookController.FetchByID)
}
