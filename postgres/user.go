package postgres

import (
	"database/sql"

	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/repositories"
)

type userRepo struct {
	db *sql.DB
}

func NewRepositoryUser(db *sql.DB) repositories.User {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) FindByEmail(email string) (*models.User, error) {
	user := models.User{}

	err := r.db.QueryRow(
		`SELECT id, email, password FROM users WHERE email=$1 LIMIT 1`,
		email,
	).
		Scan(&user.ID, &user.Email, &user.Password)

	if err == sql.ErrNoRows {
		return nil, repositories.ErrUserNotFound
	}

	return &user, err
}

func (r *userRepo) Create(user *models.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (id, email, password, created_at) VALUES ($1, $2, $3, $4)",
		user.ID, user.Email, user.Password, user.CreatedAt)

	return err
}

func (r *userRepo) ExistsEmail(email string) (bool, error) {
	var exists bool

	err := r.db.QueryRow("SELECT exists (SELECT FROM users WHERE email = $1)", email).Scan(&exists)

	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	return exists, nil
}
