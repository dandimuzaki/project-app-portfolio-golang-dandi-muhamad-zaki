package repository

import (
	"github.com/dandimuzaki/project-app-portfolio-golang/database"
	"go.uber.org/zap"
)

type Repository struct {
	UserRepo UserRepository
	ProjectRepo ProjectRepository
}

func NewRepository(db database.PgxIface, log *zap.Logger) Repository {
	return Repository{
		UserRepo: NewUserRepository(db, log),
		ProjectRepo: NewProjectRepository(db, log),
	}
}