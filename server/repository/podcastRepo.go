package repository

import (
	"Clean_Architecture/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PodcastRepo struct {
	db *pgxpool.Pool
}

func NewPodcastRepo(dbPool *pgxpool.Pool) domain.PodcastRepository {
	return &PodcastRepo{
		db: dbPool,
	}
}

func (tr *PodcastRepo) Create(c context.Context, podcast *domain.Podcast) error {
	query := `INSERT INTO podcasts (id, title, description, url)
              VALUES ($1, $2, $3, $4)`

	_, err := tr.db.Exec(c, query, podcast.ID, podcast.Title, podcast.Description, podcast.Url)
	return err
}

func (tr *PodcastRepo) FetchByID(c context.Context, id string) (*domain.Podcast, error) {
	query := `SELECT id, title, description, url
			  FROM podcasts
			  WHERE id = $1`

	row := tr.db.QueryRow(c, query, id)
	var podcast domain.Podcast
	err := row.Scan(&podcast.ID, &podcast.Title, &podcast.Description, &podcast.Url)
	if err != nil {
		return nil, err
	}

	return &podcast, nil
}
