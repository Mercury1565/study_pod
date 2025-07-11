package repository

import (
	"Clean_Architecture/domain"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(dbPool *pgxpool.Pool) domain.UserRepository {
	return &userRepo{
		db: dbPool,
	}
}

func (ur *userRepo) Create(c context.Context, user *domain.User) error {
	query := `INSERT INTO users (id, name, email, password)
              VALUES ($1, $2, $3, $4)`
	_, err := ur.db.Exec(c, query, user.UserID, user.Name, user.Email, user.Password)
	return err
}

func (ur *userRepo) GetByEmail(c context.Context, email string) (domain.User, error) {
	var user domain.User
	query := `SELECT id, name, email, password
              FROM users
              WHERE email = $1`
	err := ur.db.QueryRow(c, query, email).Scan(&user.UserID, &user.Name, &user.Email, &user.Password,)
	if err == pgx.ErrNoRows {
		return domain.User{}, fmt.Errorf("user with email %s not found", email)
	}
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (ur *userRepo) GetByID(c context.Context, id string) (domain.User, error) {
	var user domain.User
	query := `SELECT id, name, email, password
              FROM users
              WHERE id = $1`
	err := ur.db.QueryRow(c, query, id).Scan(&user.UserID, &user.Name, &user.Email, &user.Password,)
	if err == pgx.ErrNoRows {
		return domain.User{}, fmt.Errorf("user with ID %s not found", id)
	}
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
