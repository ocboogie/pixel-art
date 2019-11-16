package postgres

import (
	"database/sql"

	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/repositories"
)

type postRepo struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) repositories.Post {
	return &postRepo{
		db: db,
	}
}

func (r *postRepo) Find(id string) (*models.Post, error) {
	post := models.Post{}

	err := r.db.QueryRow(
		`SELECT id, user_id, title, data, create_at FROM posts WHERE id=$1 LIMIT 1`,
		id,
	).
		Scan(&post.ID, &post.UserID, &post.Title, &post.Data, &post.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, repositories.ErrPostNotFound
	}

	return &post, err
}

func (r *postRepo) Save(post *models.Post) error {
	_, err := r.db.Exec(
		"INSERT INTO posts (id, user_id, title, data, created_at) VALUES ($1, $2, $3, $4, $5)",
		post.ID, post.UserID, post.Title, post.Data, post.CreatedAt,
	)

	return err
}

func (r *postRepo) Latest(limit uint) ([]*models.Post, error) {
	rows, err := r.db.Query(
		`SELECT id, user_id, title, data, created_at FROM posts ORDER BY created_at LIMIT $1`,
		limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := make([]*models.Post, limit)
	// FIXME: This is just weird
	i := 0
	for rows.Next() {
		rows.Scan(
			&posts[i].ID,
			&posts[i].UserID,
			&posts[i].Title,
			&posts[i].Data,
			&posts[i].CreatedAt,
		)

		i++
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
