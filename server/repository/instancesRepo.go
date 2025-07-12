package repository

import (
	"Clean_Architecture/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type InstanceRepo struct {
	db *pgxpool.Pool
}

func NewInstanceRepo(dbPool *pgxpool.Pool) domain.InstanceRepository {
	return &InstanceRepo{
		db: dbPool,
	}
}

func (tr *InstanceRepo) Create(c context.Context, instance *domain.Instance) error {
	query := `INSERT INTO instances (id, user_id, book_id, chapter_id, podcast_id)
              VALUES ($1, $2, $3, $4)`
	_, err := tr.db.Exec(c, query, instance.ID, instance.BookID, instance.ChapterID, instance.PodcastID)
	return err
}

func (tr *InstanceRepo) FetchByID(c context.Context, id string) (*domain.Instance, error) {
	query := `SELECT id, user_id, book_id, chapter_id, podcast_id
			  FROM instances
			  WHERE id = $1`
	row := tr.db.QueryRow(c, query, id)

	var instance domain.Instance
	err := row.Scan(&instance.ID, &instance.UserID, &instance.BookID, &instance.ChapterID, &instance.PodcastID)
	if err != nil {
		return nil, err
	}

	return &instance, nil
}

func (tr *InstanceRepo) FetchByUserID(c context.Context, id string) ([]domain.Instance, error) {
	query := `SELECT id, user_id, book_id, chapter_id, podcast_id
              FROM instances
              WHERE id = $1`
	rows, err := tr.db.Query(c, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var instances []domain.Instance
	for rows.Next() {
		var instance domain.Instance
		err := rows.Scan(&instance.ID, &instance.UserID, &instance.BookID, &instance.ChapterID, &instance.PodcastID)
		if err != nil {
			return nil, err
		}
		instances = append(instances, instance)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(instances) == 0 {
		return []domain.Instance{}, nil
	}

	return instances, nil
}