package service

import (
	"github.com/dandimuzaki/project-app-portfolio-golang/repository"
	"go.uber.org/zap"
)

type Service struct {
	UserService UserService
	ProjectService ProjectService
	AuthService AuthService
}

func NewService(repo repository.Repository, log *zap.Logger) Service {
	return Service{
		UserService: NewUserService(repo, log),
		ProjectService: NewProjectService(repo, log),
		AuthService: NewAuthService(repo, log),
	}
}