package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/ocboogie/pixel-art/repositories"
)

type followRepo struct {
	db *sqlx.DB
}

func NewFollowRepository(db *sqlx.DB) repositories.Follow {
	return &followRepo{
		db: db,
	}
}

func (r *followRepo) Save(followedID string, followerID string) error {
	_, err := r.db.Exec(
		`INSERT INTO follows (followed_id, follower_id) 
		 VALUES ($1, $2)`,
		followedID,
		followerID,
	)

	if err, ok := err.(*pq.Error); ok {
		switch err.Code.Name() {
		case "unique_violation":
			return repositories.ErrFollowAlreadyExists
		case "foreign_key_violation":
			return repositories.ErrUserNotFound
		}
	}

	return err
}

func (r *followRepo) Delete(followedID string, followerID string) error {
	result, err := r.db.Exec(
		`DELETE FROM follows WHERE follows.followed_id = $1 AND follows.follower_id = $2`,
		followedID,
		followerID,
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
		return repositories.ErrFollowNotFound
	}

	return nil
}
