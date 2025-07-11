package main

import (
	"Clean_Architecture/api/route"
	"Clean_Architecture/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	env := app.Env

	// Use the PostgreSQL connection pool
	dbPool := app.Pool
	defer app.ClosePostgresConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	// Update route setup to use dbPool instead of MongoDB database
	route.Setup(env, timeout, dbPool, gin)
	gin.Run(env.ServerAddress)
}
