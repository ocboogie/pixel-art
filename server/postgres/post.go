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

// Not sure if the use of author.id in the GROUP BY will break anything but
// postgres complains otherwise
const postSelect = `
SELECT
	posts.id "id",
	posts.title "title",
	posts.data "data",
	posts.created_at "created_at",
	count(likes.post_id) "likes",
	author.id "author.id",
	author.name "author.name",
	author.avatar "author.avatar",
	author.created_at "author.created_at"
FROM
	posts
JOIN
	users AS author ON posts.author_id = author.id
LEFT OUTER JOIN
	likes ON posts.id = likes.post_id
GROUP BY 
	posts.id, author.id
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
		HAVING 
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
			HAVING
				posts.created_at > $2
			ORDER BY
				created_at 
			LIMIT
				$1`,
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

func (r *postRepo) PostsByUser(userID string, limit int, after *time.Time) ([]*models.Post, error) {
	// TODO: DRY: Latest
	posts := []*models.Post{}

	var err error
	if after != nil {
		err = r.db.Select(
			&posts,
			postSelect+`
			HAVING
				posts.author_id = $1 AND 
				posts.created_at > $2
			ORDER BY
				created_at 
			LIMIT
				$3`,
			userID,
			after,
			limit,
		)
	} else {
		err = r.db.Select(
			&posts,
			postSelect+`
			HAVING
				posts.author_id = $1
			ORDER BY
				created_at 
			LIMIT 
				$2`,
			userID,
			limit,
		)
	}
	if err != nil {
		return nil, err
	}

	return posts, nil
}
