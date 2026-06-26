package service

import (
	"strings"
	"time"

	"users-service/internal/model"
	"users-service/internal/domain"
	"users-service/pkg/apperror"
)



type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAll() ([]*model.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) Create(name, email string) (*model.User, error) {
	// Simple validation / business rules
	if strings.HasSuffix(email, "@mailinator.com") {
		return nil, apperror.NewInvalidInput("disposable emails are not allowed")
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

func (s *UserService) Update(id int64, name, email string) (*model.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	user.Name = name
	user.Email = email

	return s.repo.Update(user)
}

func (s *UserService) DeleteByID(id int64) error {
	return s.repo.Delete(id)
}

/*
Clears users who haven't logged in since 'days' days ago.
This logic runs with a cron job
*/
func (s *UserService) CleanupInactive(days int) (deleted int64, err error) {
	cutoff := time.Now().AddDate(0, 0, -days)

	return s.repo.DeleteInactiveBefore(cutoff)
}
