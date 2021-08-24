package postgres

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/ocboogie/pixel-art/models"
	"github.com/ocboogie/pixel-art/repositories"
)

type userRepo struct {
	db *sqlx.DB
	sb sq.StatementBuilderType
}

func NewRepositoryUser(db *sqlx.DB, sb sq.StatementBuilderType) repositories.User {
	return &userRepo{
		db: db,
		sb: sb,
	}
}

func userBaseSelect(sb sq.StatementBuilderType, includes repositories.UserIncludes) sq.SelectBuilder {
	stmt := sb.Select(`*`).
		From("users")

	if includes.IsFollowing != "" {
		stmt = stmt.
			Column("EXISTS(SELECT 1 FROM follows WHERE follows.followed_id = users.id AND follows.follower_id = ?) AS is_following", includes.IsFollowing)
	}

	if includes.Followers {
		// TODO: Could to be better to use a left join idk
		stmt = stmt.
			Column("(SELECT count(*) FROM follows WHERE follows.followed_id = users.id) followers")
	}
	if includes.FollowingCount {
		stmt = stmt.
			Column("(SELECT count(*) FROM follows WHERE follows.follower_id = users.id) following_count")
	}

	return stmt
}

func (r *userRepo) Find(id string, includes repositories.UserIncludes) (*models.User, error) {
	user := models.User{}

	query, args, err := userBaseSelect(r.sb, includes).
		Where("users.id=?", id).
		Limit(1).
		ToSql()
	if err != nil {
		panic(err)
	}

	err = r.db.Get(&user, query, args...)

	if err == sql.ErrNoRows {
		return nil, repositories.ErrUserNotFound
	}

	return &user, err
}

func (r *userRepo) All(includes repositories.UserIncludes) ([]*models.User, error) {
	users := make([]*models.User, 0)

	query, args, err := userBaseSelect(r.sb, includes).ToSql()
	if err != nil {
		panic(err)
	}

	err = r.db.Select(&users, query, args...)

	return users, nil
}

func (r *userRepo) FindByEmail(email string, includes repositories.UserIncludes) (*models.User, error) {
	user := models.User{}

	query, args, err := userBaseSelect(r.sb, includes).
		Where("users.email=?", email).
		Limit(1).
		ToSql()
	if err != nil {
		panic(err)
	}

	err = r.db.Get(&user, query, args...)

	if err == sql.ErrNoRows {
		return nil, repositories.ErrUserNotFound
	}

	return &user, err
}

func (r *userRepo) Update(user *models.User) error {
	_, err := r.db.NamedExec(`UPDATE users SET name=:name, avatar=:avatar, email=:email, password=:password WHERE id=:id`, user)

	return err
}

func (r *userRepo) Save(user *models.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (id, name, avatar, email, password, created_at) VALUES ($1, $2, $3, $4, $5, $6)",
		user.ID, user.Name, user.Avatar, user.Email, user.Password, user.CreatedAt)

	return err
}

func (r *userRepo) ExistsEmail(email string) (bool, error) {
	var exists bool

	// TODO: Make this work with sqlx better
	err := r.db.QueryRow("SELECT exists (SELECT FROM users WHERE email = $1)", email).Scan(&exists)

	if err != nil && err != sql.ErrNoRows {
		return false, err
	}

	return exists, nil
}
