package bootstrap

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	Env   *Env
	Pool  *pgxpool.Pool
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Pool = NewPostgresPool(app.Env)
	return *app
}

func (app *Application) ClosePostgresConnection() {
	if app.Pool != nil {
		app.Pool.Close()
	}
}
