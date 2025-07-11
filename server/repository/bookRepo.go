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
	query := `INSERT INTO books (id, user_id, title, url)
              VALUES ($1, $2, $3, $4)`
	_, err := tr.db.Exec(c, query, book.ID, book.UserID, book.Title, book.Url)
	return err
}

func (tr *BookRepo) FetchByUserID(c context.Context, userID string) ([]domain.Book, error) {
	query := `SELECT id, user_id, title, url
              FROM books
              WHERE user_id = $1`
	rows, err := tr.db.Query(c, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []domain.Book
	for rows.Next() {
		var book domain.Book
		err := rows.Scan(&book.ID, &book.UserID, &book.Title, &book.Url)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(books) == 0 {
		return []domain.Book{}, nil
	}

	return books, nil
}
