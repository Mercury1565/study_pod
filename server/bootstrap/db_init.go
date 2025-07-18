package bootstrap

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func initializeSchema(ctx context.Context, pool *pgxpool.Pool) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(255) PRIMARY KEY,
			name VARCHAR(255),
			email VARCHAR(255) UNIQUE,
			password VARCHAR(255)
		)`,

		`CREATE TABLE IF NOT EXISTS books (
			id VARCHAR(255) PRIMARY KEY,
			title VARCHAR(255),
			url VARCHAR(255)
		)`,

		`CREATE TABLE IF NOT EXISTS podcasts (
		    id VARCHAR(255) PRIMARY KEY,
		    title VARCHAR(255),
		    description VARCHAR(255),
		    url VARCHAR(255)
		)`,

		`CREATE TABLE IF NOT EXISTS chapters (
		    id VARCHAR(255) PRIMARY KEY,
		    book_id VARCHAR(255),
			chapter INTEGER,
		    title VARCHAR(255),
		    description VARCHAR(255),
		    FOREIGN KEY (book_id) REFERENCES books(id)
		)`,

		`CREATE TABLE IF NOT EXISTS instances (
		    id VARCHAR(255) PRIMARY KEY,
		    user_id VARCHAR(255),
		    book_id VARCHAR(255),
		    chapter_id VARCHAR(255),
		    podcast_id VARCHAR(255),
		    FOREIGN KEY (user_id) REFERENCES users(id),
		    FOREIGN KEY (book_id) REFERENCES books(id),
		    FOREIGN KEY (chapter_id) REFERENCES chapters(id),
		    FOREIGN KEY (podcast_id) REFERENCES podcasts(id),
		)`,
	}
	for _, query := range queries {
		_, err := pool.Exec(ctx, query)
		if err != nil {
			return fmt.Errorf("failed to execute query: %v, error: %w", query, err)
		}
	}

	return nil
}
