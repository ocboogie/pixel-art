package postgres

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/repositories"
)

type postRepo struct {
	db *sqlx.DB
	sb sq.StatementBuilderType
}

// postBaseSelect is used to fetch post(s) because the query for posts is pretty
// complex
func postBaseSelect(sb sq.StatementBuilderType, includes repositories.PostIncludes) sq.SelectBuilder {
	stmt := sb.Select(`posts.*`).
		From("posts")

	if includes.Author {
		stmt = stmt.
			Join("users AS author ON posts.author_id = author.id").
			// Ugh https://github.com/jmoiron/sqlx/issues/131
			Columns(`author.id "author.id",
				author.name "author.name",
				author.avatar "author.avatar"`)
	}
	if includes.Likes {
		stmt = stmt.
			LeftJoin("likes ON posts.id = likes.post_id").
			Column(`COUNT(likes.post_id) likes`).
			GroupBy("posts.id")

		// See https://stackoverflow.com/q/19601948/4910911
		if includes.Author {
			stmt = stmt.GroupBy("author.id")
		}
	}
	if includes.Liked != "" {
		stmt = stmt.
			Column("EXISTS(SELECT 1 FROM likes WHERE likes.post_id = posts.id AND likes.user_id = ?) liked", includes.Liked)
	}
	return stmt
}

func NewPostRepository(db *sqlx.DB, sb sq.StatementBuilderType) repositories.Post {
	return &postRepo{
		db: db,
		sb: sb,
	}
}

func (r *postRepo) Find(id string, includes repositories.PostIncludes) (*models.Post, error) {
	post := models.Post{}

	query, args, err := postBaseSelect(r.sb, includes).
		Where("posts.id=?", id).
		ToSql()
	if err != nil {
		panic(err)
	}

	err = r.db.Get(&post, query, args...)

	if err == sql.ErrNoRows {
		return nil, repositories.ErrPostNotFound
	}

	return &post, err
}

func (r *postRepo) Delete(id string) error {
	res, err := r.db.Exec(
		`DELETE FROM posts
		 WHERE id=$1`,
		id,
	)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if count != 1 {
		return repositories.ErrPostNotFound
	}

	return nil
}

func (r *postRepo) Save(post *models.Post) error {
	_, err := r.db.NamedExec(
		`INSERT INTO posts (id, author_id, title, art, created_at) 
		 VALUES (:id, :author_id, :title, :art, :created_at)`,
		post,
	)

	return err
}

func (r *postRepo) Latest(limit int, after *time.Time, includes repositories.PostIncludes) ([]*models.Post, error) {
	// TODO: This could be faster by using make with limit as the size but this
	//       causes nulls in the output
	posts := []*models.Post{}

	stmt := postBaseSelect(r.sb, includes)

	if after != nil {
		stmt = stmt.Having("posts.created_at > ?", after)
	}

	query, args, err := stmt.
		OrderBy("posts.created_at").
		Limit(uint64(limit)).
		ToSql()
	if err != nil {
		panic(err)
	}

	if err = r.db.Select(&posts, query, args...); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *postRepo) PostsByUser(userID string, limit int, after *time.Time, includes repositories.PostIncludes) ([]*models.Post, error) {
	// TODO: DRY: Latest
	posts := []*models.Post{}

	stmt := postBaseSelect(r.sb, includes)

	if after != nil {
		stmt = stmt.Having("posts.created_at > ?", after)
	}

	query, args, err := stmt.
		Where("posts.author_id = ?", userID).
		OrderBy("posts.created_at").
		Limit(uint64(limit)).
		ToSql()
	if err != nil {
		panic(err)
	}

	if err := r.db.Select(&posts, query, args...); err != nil {
		return nil, err
	}

	return posts, nil
}
