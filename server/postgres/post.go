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

// BaseSelect is used to fetch post(s) because the query for posts is pretty
// complex
func BaseSelect(sb sq.StatementBuilderType) sq.SelectBuilder {
	return sb.Select(`posts.id "id",
			posts.title "title",
			posts.data "data",
			posts.created_at "created_at",
			count(likes.post_id) "likes",
			author.id "author.id",
			author.name "author.name",
			author.avatar "author.avatar",
			author.created_at "author.created_at"`).
		From("posts").
		Join("users AS author ON posts.author_id = author.id").
		JoinClause("LEFT OUTER JOIN likes ON posts.id = likes.post_id").
		// Not sure if the use of author.id in the GROUP BY will break anything but
		// postgres complains otherwise
		GroupBy("posts.id", "author.id")
}

func NewPostRepository(db *sqlx.DB, sb sq.StatementBuilderType) repositories.Post {
	return &postRepo{
		db: db,
		sb: sb,
	}
}

func (r *postRepo) Find(id string) (*models.Post, error) {
	post := models.Post{}

	query, args, _ := BaseSelect(r.sb).
		Having("posts.id=?", id).
		ToSql()

	err := r.db.Get(&post, query, args...)

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

	stmt := BaseSelect(r.sb)

	if after != nil {
		stmt = stmt.Having("posts.created_at > ?", after)
	}

	query, args, _ := stmt.
		OrderBy("created_at").
		Limit(uint64(limit)).
		ToSql()

	if err := r.db.Select(&posts, query, args...); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *postRepo) PostsByUser(userID string, limit int, after *time.Time) ([]*models.Post, error) {
	// TODO: DRY: Latest
	posts := []*models.Post{}

	stmt := BaseSelect(r.sb)

	if after != nil {
		stmt = stmt.Having("posts.created_at > ?", after)
	}

	query, args, _ := stmt.
		Having("posts.author_id = ?", userID).
		OrderBy("created_at").
		Limit(uint64(limit)).
		ToSql()

	if err := r.db.Select(&posts, query, args...); err != nil {
		return nil, err
	}

	return posts, nil
}
