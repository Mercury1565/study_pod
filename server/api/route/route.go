package route

import (
	"Clean_Architecture/api/middelware"
	"Clean_Architecture/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Setup(env *bootstrap.Env, timeout time.Duration, dbPool *pgxpool.Pool, gin *gin.Engine) {
		publicRouter := gin.Group("")
		protectedRouter := gin.Group("")

		// Middleware to verify AccessToken
		protectedRouter.Use(middelware.JWTAuthMiddleware(env.AccessTokenSecret))

		// All public APIs
		NewSignupRouter(env, timeout, dbPool, publicRouter)
		NewLoginRouter(env, timeout, dbPool, publicRouter)
		NewRefreshTokenRouter(env, timeout, dbPool, publicRouter)

		// All private APIs
		NewProfileRouter(env, timeout, dbPool, protectedRouter)
		NewBookRouter(env, timeout, dbPool, protectedRouter)
		NewChapterRouter(env, timeout, dbPool, protectedRouter)
		NewPodcastRouter(env, timeout, dbPool, protectedRouter)
		NewInstanceRouter(env, timeout, dbPool, protectedRouter)
}
