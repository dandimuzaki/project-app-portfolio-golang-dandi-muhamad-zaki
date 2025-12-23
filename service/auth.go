package service

import (
	"errors"

	"github.com/dandimuzaki/project-app-portfolio-golang/model"
	"github.com/dandimuzaki/project-app-portfolio-golang/repository"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type AuthService interface {
	Login(email, password string) (*model.User, error)
	SignIn(user model.User) (*model.User, error)
}

type authService struct {
	Repo repository.Repository
	Logger *zap.Logger
}

func NewAuthService(repo repository.Repository, log *zap.Logger) AuthService {
	return &authService{
		Repo: repo,
		Logger: log,
	}
}

func (s *authService) Login(email, password string) (*model.User, error) {
	user, err := s.Repo.UserRepo.FindByEmail(email)
	if err == pgx.ErrNoRows {
		s.Logger.Error("User not found: ", zap.Error(err))
		return nil, errors.New("user not found")
	}

	if err != nil {
		s.Logger.Error("Error find user by email: ", zap.Error(err))
		return nil, nil
	}

	if *user.Password != password {
		s.Logger.Error("Incorrect password: ", zap.Error(err))
		return nil, errors.New("incorrect password")
	}

	return user, nil
}

func (s *authService) SignIn(data model.User) (*model.User, error) {
	user, err := s.Repo.UserRepo.FindByEmail(*data.Email)
	if user != nil {
		s.Logger.Error("Cannot register existing user: ", zap.Error(err))
		return user, errors.New("user already registered")
	}

	user, err = s.Repo.UserRepo.Create(data)
	if err != nil {
		s.Logger.Error("Error create user service: ", zap.Error(err))
	}
	return user, nil
}