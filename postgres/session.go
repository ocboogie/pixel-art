package postgres

import (
	"database/sql"

	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/repositories"
)

type sessionRepo struct {
	db *sql.DB
}

func NewSessionRepository(db *sql.DB) repositories.Session {
	return &sessionRepo{
		db: db,
	}
}

func (r *sessionRepo) Create(session *models.Session) error {
	println(session.ID)

	_, err := r.db.Exec(
		"INSERT INTO sessions (id, user_id, expires_at) VALUES ($1, $2, $3)",
		session.ID, session.UserID, session.ExpiresAt,
	)

	return err
}

func (r *sessionRepo) Find(id string) (*models.Session, error) {
	session := models.Session{}

	err := r.db.QueryRow(
		`SELECT id, user_id, expires_at FROM sessions WHERE id=$1 LIMIT 1`,
		id,
	).
		Scan(&session.ID, &session.UserID, &session.ExpiresAt)

	if err == sql.ErrNoRows {
		return nil, repositories.ErrSessionNotFound
	}

	return &session, err
}

func (r *sessionRepo) Delete(id string) error {
	_, err := r.db.Exec(
		`DELETE FROM sessions WHERE id=$1`,
		id,
	)

	if err == sql.ErrNoRows {
		return repositories.ErrSessionNotFound
	}

	return err
}
