package repository

import (
	"context"

	"github.com/dandimuzaki/project-app-portfolio-golang/database"
	"github.com/dandimuzaki/project-app-portfolio-golang/model"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type UserRepository interface {
	Create(user model.User) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	GetUserByID(id int) (*model.User, error)
	UpdateUser(userID int, u *model.User) error
}

type userRepository struct {
	db database.PgxIface
	Logger *zap.Logger
}

func NewUserRepository(db database.PgxIface, log *zap.Logger) UserRepository {
	return &userRepository{
		db: db,
		Logger: log,
	}
}

func (r *userRepository) Create(user model.User) (*model.User, error) {
	query := `
		INSERT INTO users (name, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING id
	`

	err := r.db.QueryRow(context.Background(), query, user.Name, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		r.Logger.Error("Error query create user: ", zap.Error(err))
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	query := `
		SELECT id, email, password
		FROM users
		WHERE email = $1 AND deleted_at IS NULL
	`
	var user model.User
	err := r.db.QueryRow(context.Background(), query, email).Scan(
		&user.ID, &user.Email, &user.Password,
	)

	if err == pgx.ErrNoRows {
		r.Logger.Error("User not found on given email: ", zap.Error(err))
		return nil, nil
	}

	if err != nil {
		r.Logger.Error("Error query find user by email: ", zap.Error(err))
		return nil, nil
	}

	return &user, err
}

func (r *userRepository) GetUserByID(id int) (*model.User, error) {
	query := "SELECT id, name, email, avatar, description, github, linkedin, cv_link, phone_number FROM users WHERE id = $1"

	var user model.User
	err := r.db.QueryRow(context.Background(), query, id).Scan(&user.ID, &user.Name, &user.Email, 
		&user.Avatar, &user.Description, &user.Github, &user.LinkedIn, &user.CV, &user.PhoneNumber)
	if err == pgx.ErrNoRows {
		r.Logger.Error("User not found on given id: ", zap.Error(err))
		return nil, err
	}

	if err != nil {
		r.Logger.Error("Error query get user by id: ", zap.Error(err))
		return nil, nil
	}

	return &user, nil
}

func (r *userRepository) UpdateUser(userID int, u *model.User) error {
	query := `
		UPDATE users
		SET name = COALESCE($1, name), email = COALESCE($2, email),
		description = COALESCE($3, description), avatar = COALESCE($4, avatar),
		github = COALESCE($5, github), linkedin = COALESCE($6, linkedin),
		cv_link = COALESCE($7, cv_link), phone_number = COALESCE($8, phone_number),
		updated_at = NOW()
		WHERE id = $9 AND deleted_at IS NULL
	`

	result, err := r.db.Exec(context.Background(), query,
		&u.Name, &u.Email, &u.Description, &u.Avatar, &u.Github,
		&u.LinkedIn, &u.CV, &u.PhoneNumber, userID,
	)

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		r.Logger.Error("Error not found user: ", zap.Error(err))
		return nil
	}

	if err != nil {
		r.Logger.Error("Error query update user: ", zap.Error(err))
	}

	return err
}
