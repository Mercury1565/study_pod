package repository

import (
	"Clean_Architecture/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type BookRepo struct {
	db *pgxpool.Pool
}

func NewBookRepo(dbPool *pgxpool.Pool) domain.BookRepository {
	return &BookRepo{
		db: dbPool,
	}
}

func (tr *BookRepo) Create(c context.Context, book *domain.Book) error {
	query := `INSERT INTO books (id, title, url)
              VALUES ($1, $2, $3, $4)`
	_, err := tr.db.Exec(c, query, book.ID, book.Title, book.Url)
	return err
}

func (tr *BookRepo) FetchByID(c context.Context, id string) (*domain.Book, error) {
	query := `SELECT id, title, url
			  FROM books
			  WHERE id = $1`
	row := tr.db.QueryRow(c, query, id)

	var book domain.Book
	err := row.Scan(&book.ID, &book.Title, &book.Url)
	if err != nil {
		return nil, err
	}

	return &book, nil
}
