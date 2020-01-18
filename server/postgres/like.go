package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/ocboogie/pixel-art/repositories"
)

type likeRepo struct {
	db *sqlx.DB
}

func NewLikeRepository(db *sqlx.DB) repositories.Like {
	return &likeRepo{
		db: db,
	}
}

func (r *likeRepo) Save(userID string, postID string) error {
	_, err := r.db.Exec(
		`INSERT INTO likes (user_id, post_id) 
		 VALUES ($1, $2)`,
		userID,
		postID,
	)

	if err, ok := err.(*pq.Error); ok {
		switch err.Code.Name() {
		case "unique_violation":
			return repositories.ErrLikeAlreadyExists
		case "foreign_key_violation":
			if err.Constraint == "likes_post_id_fkey" {
				return repositories.ErrPostNotFound
			} else if err.Constraint == "likes_user_id_fkey" {
				return repositories.ErrUserNotFound
			}
		}
	}

	return err
}

func (r *likeRepo) Delete(userID string, postID string) error {
	result, err := r.db.Exec(
		`DELETE FROM likes WHERE likes.user_id = $1 AND likes.post_id = $2`,
		userID,
		postID,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// This should never happen if we're using the postgres driver
		panic("Not using a sql driver that supports RowsAffected")
	}

	if rowsAffected == 0 {
		return repositories.ErrLikeNotFound
	}

	return nil
}
