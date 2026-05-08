package service

import (
	"strings"
	"time"

	"users-service/internal/model"
	"users-service/internal/repository"
	"users-service/pkg/app_error"
)

type UserService struct {
	repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create(name, email string) (*model.User, error) {
	// Simple validation / business rules
	if strings.TrimSpace(name) == "" {
		return nil, app_error.NewInvalidInput("name is required")
	}

	if !strings.Contains(email, "@") {
		return nil, app_error.NewInvalidInput("invalid email format")
	}

	user := &model.User{
		Name:      name,
		Email:     email,
		Active:    true,
		LastLogin: time.Now(),
		CreatedAt: time.Now(),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetByID(id int64) (*model.User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) DeactivateUser(id int64) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	user.Active = false

	return s.repo.Update(user)
}

/*
Clears users who haven't logged in since 'days' days ago.
This logic runs with a cron job
*/
func (s *UserService) CleanupInactive(days int) (deleted int64, err error) {
	cutoff := time.Now().AddDate(0, 0, -days)

	return s.repo.DeleteInactiveBefore(cutoff)
}
