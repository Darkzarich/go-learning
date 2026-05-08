package repository

import (
	"database/sql"
	"time"

	"users-service/internal/model"
	"users-service/pkg/app_error"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

// Creates initial database and tables if they don't exist.
func (r *UserRepo) InitDatabase() error {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE,
        active INTEGER NOT NULL DEFAULT 1,
        last_login DATETIME,
        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
    );`

	_, err := r.db.Exec(query)

	return err
}

func (r *UserRepo) Create(user *model.User) error {
	result, err := r.db.Exec(
		"INSERT INTO users (name, email, active, last_login) VALUES (?, ?, ?, ?)",
		user.Name, user.Email, user.Active, user.LastLogin,
	)
	if err != nil {
		return app_error.NewInternal(err)
	}

	id, _ := result.LastInsertId()
	user.ID = id

	return nil
}

func (r *UserRepo) FindByID(id int64) (*model.User, error) {
	user := &model.User{}

	err := r.db.QueryRow(
		"SELECT id, name, email, active, last_login, created_at FROM users WHERE id = ?", id,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Active, &user.LastLogin, &user.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, app_error.NewNotFound("user not found")
	} else if err != nil {
		return nil, app_error.NewInternal(err)
	}

	return user, nil
}

func (r *UserRepo) Update(user *model.User) error {
	_, err := r.db.Exec(
		"UPDATE users SET name = ?, email = ?, active = ?, last_login = ? WHERE id = ?",
		user.Name, user.Email, user.Active, user.LastLogin, user.ID,
	)
	if err != nil {
		return app_error.NewInternal(err)
	}

	return nil
}

func (r *UserRepo) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}

// DeleteInactiveBefore deletes users who haven't logged in since 'before'.
func (r *UserRepo) DeleteInactiveBefore(before time.Time) (int64, error) {
	res, err := r.db.Exec("DELETE FROM users WHERE last_login < ?", before)
	if err != nil {
		return 0, app_error.NewInternal(err)
	}

	return res.RowsAffected()
}
