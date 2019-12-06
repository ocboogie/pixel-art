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
	post := struct {
		models.Post
		AuthorID string `db:"author_id"`
	}{}

	err := r.db.Get(&post,
		`SELECT 
			posts.*, 
			author.id "author.id", 
			author.name "author.name", 
			author.created_at "author.created_at"
		FROM 
			posts JOIN users AS author ON posts.author_id = author.id
		WHERE 
			posts.id=$1`,
		id)

	if err == sql.ErrNoRows {
		return nil, repositories.ErrPostNotFound
	}

	return &post.Post, err
}

func (r *postRepo) Save(post *models.Post) error {
	_, err := r.db.NamedExec(
		`INSERT INTO posts (id, author_id, title, data, created_at) 
		 VALUES (:id, :author.id, :title, :data, :created_at)`,
		post,
	)

	return err
}

func (r *postRepo) Latest(limit int) ([]*models.Post, error) {
	posts := []*models.Post{}
	err := r.db.Select(
		&posts,
		`SELECT 
			posts.*, 
			author.id "author.id", 
			author.name "author.name", 
			author.created_at "author.created_at"
		FROM 
			posts JOIN users AS author ON posts.author_id = author.id
		ORDER BY
			created_at LIMIT $1`,
		limit,
	)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
