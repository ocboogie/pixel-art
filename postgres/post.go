package postgres

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/repositories"
)

type postRepo struct {
	db *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) repositories.Post {
	return &postRepo{
		db: db,
	}
}

func (r *postRepo) Find(id string) (*models.Post, error) {
	post := models.Post{}

	err := r.db.Get(&post, "SELECT * FROM posts WHERE id=$1 LIMIT 1", id)

	if err == sql.ErrNoRows {
		return nil, repositories.ErrPostNotFound
	}

	return &post, err
}

func (r *postRepo) Save(post *models.Post) error {
	_, err := r.db.NamedExec(
		`INSERT INTO posts (id, user_id, title, data, created_at) 
		 VALUES (:id, :user_id, :title, :data, :created_at)`,
		post,
	)

	return err
}

func (r *postRepo) Latest(limit int) ([]*models.Post, error) {
	posts := []*models.Post{}
	err := r.db.Select(
		&posts,
		`SELECT * FROM posts ORDER BY created_at LIMIT $1`,
		limit,
	)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
