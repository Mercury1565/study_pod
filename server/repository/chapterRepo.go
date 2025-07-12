package repository

import (
	"Clean_Architecture/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ChapterRepo struct {
	db *pgxpool.Pool
}

func NewChapterRepo(dbPool *pgxpool.Pool) domain.ChapterRepository {
	return &ChapterRepo{
		db: dbPool,
	}
}

func (tr *ChapterRepo) Create(c context.Context, chapter *domain.Chapter) error {
	query := `INSERT INTO chapters (id, book_id, title, description)
              VALUES ($1, $2, $3, $4)`

	_, err := tr.db.Exec(c, query, chapter.ID, chapter.BookID, chapter.Title, chapter.Description)
	return err
}

func (tr *ChapterRepo) FetchByID(c context.Context, id string) (*domain.Chapter, error) {
	query := `SELECT id, book_id, title, description
			  FROM chapters
			  WHERE id = $1`

	row := tr.db.QueryRow(c, query, id)

	var chapter domain.Chapter
	err := row.Scan(&chapter.ID, &chapter.BookID, &chapter.Title, &chapter.Description)
	if err != nil {
		return nil, err
	}

	return &chapter, nil
}
