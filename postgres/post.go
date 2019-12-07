package postgres

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/repositories"
)

type postRepo struct {
	db *sqlx.DB
}

const postSelect = `
SELECT 
	posts.id "id", 
	posts.title "title", 
	posts.data "data", 
	posts.created_at "created_at", 
	author.id "author.id", 
	author.name "author.name", 
	author.created_at "author.created_at"
FROM
	posts JOIN users AS author ON posts.author_id = author.id
`

func NewPostRepository(db *sqlx.DB) repositories.Post {
	return &postRepo{
		db: db,
	}
}

func (r *postRepo) Find(id string) (*models.Post, error) {
	post := models.Post{}

	err := r.db.Get(&post,
		postSelect+`
		WHERE 
			posts.id=$1`,
		id)

	if err == sql.ErrNoRows {
		return nil, repositories.ErrPostNotFound
	}

	return &post, err
}

func (r *postRepo) Save(post *models.Post) error {
	_, err := r.db.NamedExec(
		`INSERT INTO posts (id, author_id, title, data, created_at) 
		 VALUES (:id, :author.id, :title, :data, :created_at)`,
		post,
	)

	return err
}

func (r *postRepo) Latest(limit int, after *time.Time) ([]*models.Post, error) {
	// TODO: This could be faster by using make with limit as the size but this
	//       causes nulls in the output
	posts := []*models.Post{}

	var err error
	if after != nil {
		err = r.db.Select(
			&posts,
			postSelect+`
			WHERE
				posts.created_at > $2
			ORDER BY
				created_at LIMIT $1`,
			limit,
			after,
		)
	} else {
		err = r.db.Select(
			&posts,
			postSelect+`
			ORDER BY
				created_at LIMIT $1`,
			limit,
		)
	}
	if err != nil {
		return nil, err
	}

	return posts, nil
}
