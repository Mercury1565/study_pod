package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(env *Env) *pgxpool.Pool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbName := env.DBName
	dbUser := env.DBUser
	dbPassword := env.DBPassword

	// Construct the PostgreSQL connection string
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Create a connection pool
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Fatal("Failed to parse PostgreSQL config: ", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL: ", err)
	}

	// Check the connection
	err = pool.Ping(ctx)
	if err != nil {
		log.Fatal("Failed to ping PostgreSQL: ", err)
	}

	// Initialize database schema
	err = initializeSchema(ctx, pool)
	if err != nil {
		log.Fatal("Failed to initialize database schema: ", err)
	}

	fmt.Println("Connected to PostgreSQL")
	return pool
}
