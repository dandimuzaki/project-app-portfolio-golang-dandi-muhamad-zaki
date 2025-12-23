package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/dandimuzaki/project-app-portfolio-golang/dto"
	"github.com/dandimuzaki/project-app-portfolio-golang/model"
	"github.com/dandimuzaki/project-app-portfolio-golang/repository"
	"github.com/dandimuzaki/project-app-portfolio-golang/utils"
	"go.uber.org/zap"
)

type ProjectService interface {
	GetAllProjects(page, limit int, all bool) ([]model.Project, *dto.Pagination, error)
	GetPersonalProjects(userID int) ([]model.Project, error)
	GetProjectByID(projectID int) (*model.Project, error)
	CreateProject(project *model.Project, file multipart.File, fileHeader *multipart.FileHeader) error
	UpdateProject(projectID int, data *model.Project, file multipart.File, fileHeader *multipart.FileHeader) error
	DeleteProject(projectID int) error
}

type projectService struct {
	Repo repository.Repository
	Logger *zap.Logger
}

func NewProjectService(repo repository.Repository, log *zap.Logger) ProjectService {
	return &projectService{
		Repo: repo,
		Logger: log,
	}
}

func (s *projectService) GetAllProjects(page, limit int, all bool) ([]model.Project, *dto.Pagination, error) {
	// Execute repo to get all projects
	projects, total, err := s.Repo.ProjectRepo.GetAllProjects(page, limit, all)
	if err != nil {
		s.Logger.Error("Error get all projects service: ", zap.Error(err))
		return nil, nil, err
	}

	// Calculate total pages
	var totalPages int
	totalPages = utils.TotalPage(limit, total)

	// Create pagination
	var pagination dto.Pagination

	if all {
		pagination = dto.Pagination{
			TotalRecords: total,
		}
	} else {
		pagination = dto.Pagination{
			CurrentPage:  &page,
			Limit:        &limit,
			TotalPages:   &totalPages,
			TotalRecords: total,
		}
	}

	return projects, &pagination, nil
}

func (s *projectService) GetPersonalProjects(userID int) ([]model.Project, error) {
	projects, err := s.Repo.ProjectRepo.GetPersonalProjects(userID)
	if err != nil {
		s.Logger.Error("Error get projects by user id service: ", zap.Error(err))
	}
	return projects, err
}

func (s *projectService) GetProjectByID(projectID int) (*model.Project, error) {
	project, err := s.Repo.ProjectRepo.GetProjectByID(projectID)
	if err != nil {
		s.Logger.Error("Error get project by id service: ", zap.Error(err))
	}
	return project, err
}

func (s *projectService) CreateProject(project *model.Project, file multipart.File, fileHeader *multipart.FileHeader) error {
	// Save image file to disk
	uploadDir := "public/uploads"
	os.MkdirAll(uploadDir, os.ModePerm)

	var filename, filepath, accessURL string
	if file != nil {
		filename = fmt.Sprintf("%d_%d_%s", project.UserID, project.ID, fileHeader.Filename)
		filepath = fmt.Sprintf("%s/%s", uploadDir, filename)
		accessURL = fmt.Sprintf("http://localhost:8080/%s/%s", uploadDir, filename)
	}

	dst, err := os.Create(filepath)
	if err != nil {
		s.Logger.Error("Error create image filepath: ", zap.Error(err))
		return err
	}
	defer dst.Close()

	// Copy from src to dst
	_, err = io.Copy(dst, file)
	if err != nil {
		s.Logger.Error("Error copy image file to dst: ", zap.Error(err))
		return err
	}

	project.Image = &accessURL
	
	err = s.Repo.ProjectRepo.CreateProject(project)
	if err != nil {
		s.Logger.Error("Error create project service: ", zap.Error(err))
	}
	return err
}

func (s *projectService) UpdateProject(projectID int, data *model.Project, file multipart.File, fileHeader *multipart.FileHeader) error {
	// Save image file to disk
	uploadDir := "public/uploads"
	os.MkdirAll(uploadDir, os.ModePerm)

	var filename, filepath, accessURL string
	if file != nil {
		filename = fmt.Sprintf("%d_%d_%s", data.UserID, projectID, fileHeader.Filename)
		filepath = fmt.Sprintf("%s/%s", uploadDir, filename)
		accessURL = fmt.Sprintf("http://localhost:8080/%s/%s", uploadDir, filename)

		dst, err := os.Create(filepath)
		if err != nil {
			s.Logger.Error("Error create image filepath: ", zap.Error(err))
			return err
		}
		defer dst.Close()

		// Copy from src to dst
		_, err = io.Copy(dst, file)
		if err != nil {
			s.Logger.Error("Error copy image file to dst: ", zap.Error(err))
			return err
		}
	}

	if accessURL != "" {
		data.Image = &accessURL
	}
	
	err := s.Repo.ProjectRepo.UpdateProject(projectID, data)
	if err != nil {
		s.Logger.Error("Error update project service: ", zap.Error(err))
	}
	return err
}

func (s *projectService) DeleteProject(projectID int) error {
	err := s.Repo.ProjectRepo.DeleteProject(projectID)
	if err != nil {
		s.Logger.Error("Error delete project service: ", zap.Error(err))
	}
	return err
}
