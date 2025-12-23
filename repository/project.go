package repository

import (
	"context"
	"errors"

	"github.com/dandimuzaki/project-app-portfolio-golang/database"
	"github.com/dandimuzaki/project-app-portfolio-golang/model"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type ProjectRepository interface {
	GetAllProjects(page, limit int, all bool) ([]model.Project, int, error)
	GetPersonalProjects(userID int) ([]model.Project, error)
	GetProjectByID(projectID int) (*model.Project, error)
	CreateProject(project *model.Project) error
	UpdateProject(projectID int, data *model.Project) error
	DeleteProject(projectID int) error
}

type projectRepository struct {
	db database.PgxIface
	Logger *zap.Logger
}

func NewProjectRepository(db database.PgxIface, log *zap.Logger) ProjectRepository {
	return &projectRepository{
		db: db,
		Logger: log,
	}
}

func (r *projectRepository) GetAllProjects(page, limit int, all bool) ([]model.Project, int, error) {
	var offset int
	offset = (page - 1) * limit
	
	// Get total data for pagination
	var total int
	countQuery := `SELECT COUNT(*) FROM projects WHERE deleted_at IS NULL`
	err := r.db.QueryRow(context.Background(), countQuery).Scan(&total)
	if err != nil {
		r.Logger.Error("Error query count projects: ", zap.Error(err))
		return nil, 0, err
	}

	// Initiate rows
	var rows pgx.Rows
	
	// Conditional query based on page, limit, and all param
	query := `
		SELECT p.id, u.name AS owner, title, p.description, url, image, tech_stack
		FROM projects p JOIN users u ON p.user_id = u.id WHERE is_published = true AND p.deleted_at IS NULL ORDER BY p.updated_at
	`

	if !all && limit > 0 {
		query += `LIMIT $1 OFFSET $2`
		rows, err = r.db.Query(context.Background(), query, limit, offset)
	} else {
		rows, err = r.db.Query(context.Background(), query)
	}

	if err != nil {
		r.Logger.Error("Error query get all projects: ", zap.Error(err))
		return nil, 0, err
	}
	defer rows.Close()

	var projects []model.Project
	for rows.Next() {
		var p model.Project
		err := rows.Scan(
			&p.ID, &p.Owner, &p.Title,
			&p.Description, &p.URL, &p.Image, &p.TechStack,
		)
		if err != nil {
			r.Logger.Error("Error scanning projects: ", zap.Error(err))
			return nil, 0, err
		}
		projects = append(projects, p)
	}

	return projects, total, nil
}

func (r *projectRepository) GetPersonalProjects(userID int) ([]model.Project, error) {
	query := `
		SELECT id, created_at, updated_at, deleted_at, title, description, url, image, tech_stack
		FROM projects WHERE user_id = $1 AND deleted_at IS NULL ORDER BY created_at
	`

	rows, err := r.db.Query(context.Background(), query, userID)
	if err != nil {
		r.Logger.Error("Error query get projects by user id: ", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var projects []model.Project
	for rows.Next() {
		var p model.Project
		err := rows.Scan(
			&p.ID, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt, &p.Title,
			&p.Description, &p.URL, &p.Image, &p.TechStack,
		)
		if err != nil {
			r.Logger.Error("Error scanning projects: ", zap.Error(err))
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, nil
}

func (r *projectRepository) GetProjectByID(projectID int) (*model.Project, error) {
	query := `
		SELECT id, created_at, updated_at, deleted_at, user_id, title, description, url, image, tech_stack
		FROM projects WHERE id = $1 AND deleted_at IS NULL
	`

	var p model.Project
	err := r.db.QueryRow(context.Background(), query, projectID).Scan(
		&p.ID, &p.CreatedAt, &p.UpdatedAt, &p.DeletedAt, &p.UserID, &p.Title,
		&p.Description, &p.URL, &p.Image, &p.TechStack,
	)
	if err == pgx.ErrNoRows {
		r.Logger.Error("Error project not found: ", zap.Error(err))
		return nil, err
	}
	return &p, err
}

func (r *projectRepository)	CreateProject(p *model.Project) error {
	query := `
		INSERT INTO projects (user_id, title, description, url, image, tech_stack, created_at, updated_at)
		VALUES (
		$1, $2, $3, $4, $5, $6, NOW(), NOW())
		RETURNING id
	`
	err := r.db.QueryRow(context.Background(), query, p.UserID, p.Title,
		p.Description, p.URL, p.Image, p.TechStack,
	).Scan(&p.ID)
	
	if err != nil {
		r.Logger.Error("Error query create project: ", zap.Error(err))
	}

	return err
}

func (r *projectRepository)	UpdateProject(projectID int, p *model.Project) error {
	query := `
		UPDATE projects
		SET title = COALESCE($1, title), description = COALESCE($2, description), 
		url = COALESCE($3, url), image = COALESCE($4, image), 
		tech_stack = COALESCE($5, tech_stack),
		updated_at = NOW()
		WHERE id = $6 AND user_id = $7 AND deleted_at IS NULL
	`

	result, err := r.db.Exec(context.Background(), query,
		p.Title, p.Description, p.URL, p.Image, p.TechStack,
		projectID, p.UserID,
	)

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		r.Logger.Error("Error not found project: ", zap.Error(err))
		return nil
	}

	if err != nil {
		r.Logger.Error("Error query update project: ", zap.Error(err))
	}

	return err
}

func (r *projectRepository)	DeleteProject(projectID int) error {
	query := `
		UPDATE projects SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`

	result, err := r.db.Exec(context.Background(), query, projectID)
	if err != nil {
		r.Logger.Error("Error query delete project: ", zap.Error(err))
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		r.Logger.Error("Error project not found: ", zap.Error(err))
		return errors.New("assignment not found or already deleted")
	}
	return err
}