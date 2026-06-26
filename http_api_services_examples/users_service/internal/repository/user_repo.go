package repository

import (
	"database/sql"
	"time"

	"users-service/internal/model"
	"users-service/pkg/apperror"

	"github.com/mattn/go-sqlite3"
)

type PostgreSQLUserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *PostgreSQLUserRepo {
	return &PostgreSQLUserRepo{db: db}
}

// Creates initial database and tables if they don't exist.
func (r *PostgreSQLUserRepo) InitDatabase() error {
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

func (r *PostgreSQLUserRepo) FindAll() ([]*model.User, error) {
	rows, err := r.db.Query("SELECT id, name, email, active, last_login, created_at FROM users")
	if err != nil {
		return nil, apperror.NewInternal(err)
	}
	defer rows.Close()

	var users []*model.User

	for rows.Next() {
		user := &model.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Active, &user.LastLogin, &user.CreatedAt)
		if err != nil {
			return nil, apperror.NewInternal(err)
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *PostgreSQLUserRepo) Create(user *model.User) error {
	result, err := r.db.Exec(
		"INSERT INTO users (name, email, active, last_login) VALUES (?, ?, ?, ?)",
		user.Name, user.Email, user.Active, user.LastLogin,
	)

	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
				return apperror.NewAlreadyExists("user already exists")
			}
		} else {
			return apperror.NewInternal(err)
		}
	}

	id, _ := result.LastInsertId()
	user.ID = id

	return nil
}

func (r *PostgreSQLUserRepo) FindByID(id int64) (*model.User, error) {
	user := &model.User{}

	err := r.db.QueryRow(
		"SELECT id, name, email, active, last_login, created_at FROM users WHERE id = ?", id,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Active, &user.LastLogin, &user.CreatedAt)

	if err == sql.ErrNoRows {
		return nil, apperror.NewNotFound("user not found")
	} else if err != nil {
		return nil, apperror.NewInternal(err)
	}

	return user, nil
}

func (r *PostgreSQLUserRepo) Update(user *model.User) (*model.User, error) {
	result, err := r.db.Exec(
		"UPDATE users SET name = ?, email = ?, active = ?, last_login = ? WHERE id = ?",
		user.Name, user.Email, user.Active, user.LastLogin, user.ID,
	)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok {
			if sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
				return nil, apperror.NewAlreadyExists("user already exists")
			}
		} else {
			return nil, apperror.NewInternal(err)
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, apperror.NewInternal(err)
	} else if rowsAffected == 0 {
		return nil, apperror.NewNotFound("user not found")
	}

	return user, nil
}

func (r *PostgreSQLUserRepo) Delete(id int64) error {
	result, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return apperror.NewInternal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return apperror.NewInternal(err)
	} else if rowsAffected == 0 {
		return apperror.NewNotFound("user not found")
	}

	return nil
}

// DeleteInactiveBefore deletes users who haven't logged in since 'before'.
func (r *PostgreSQLUserRepo) DeleteInactiveBefore(before time.Time) (int64, error) {
	result, err := r.db.Exec("DELETE FROM users WHERE last_login < ?", before)
	if err != nil {
		return 0, apperror.NewInternal(err)
	}

	return result.RowsAffected()
}
