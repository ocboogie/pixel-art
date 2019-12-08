package postgres

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/repositories"
)

type userRepo struct {
	db *sqlx.DB
}

func NewRepositoryUser(db *sqlx.DB) repositories.User {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Find(id string) (*models.User, error) {
	user := models.User{}

	err := r.db.Get(
		&user,
		`SELECT * FROM users WHERE id=$1 LIMIT 1`,
		id,
	)

	if err == sql.ErrNoRows {
		return nil, repositories.ErrUserNotFound
	}

	return &user, err
}

func (r *userRepo) FindByEmail(email string) (*models.User, error) {
	user := models.User{}

	err := r.db.Get(
		&user,
		`SELECT * FROM users WHERE email=$1 LIMIT 1`,
		email,
	)

	if err == sql.ErrNoRows {
		return nil, repositories.ErrUserNotFound
	}

	return &user, err
}

func (r *userRepo) Save(user *models.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (id, name, email, password, created_at) VALUES ($1, $2, $3, $4, $5)",
		user.ID, user.Name, user.Email, user.Password, user.CreatedAt)

	return err
}

func (r *userRepo) ExistsEmail(email string) (bool, error) {
	var exists bool

	// TODO: Make this work with sqlx better
	err := r.db.QueryRow("SELECT exists (SELECT FROM users WHERE email = $1)", email).Scan(&exists)

	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	return exists, nil
}
