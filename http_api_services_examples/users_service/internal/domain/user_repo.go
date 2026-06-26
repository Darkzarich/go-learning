package domain

import (
	"time"

	"users-service/internal/model"
)

type UserRepository interface {
	FindAll() ([]*model.User, error)
	InitDatabase() error
	Create(user *model.User) error
	FindByID(id int64) (*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(id int64) error
	DeleteInactiveBefore(before time.Time) (int64, error)
}