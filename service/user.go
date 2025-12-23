package service

import (
	"fmt"
	"io"
	"os"

	"github.com/dandimuzaki/project-app-portfolio-golang/dto"
	"github.com/dandimuzaki/project-app-portfolio-golang/model"
	"github.com/dandimuzaki/project-app-portfolio-golang/repository"
	"go.uber.org/zap"
)

type UserService interface {
	GetUserByID(id int) (*model.User, error)
	UpdateUser(userID int, data *dto.UpdateUserRequest) error
}

type userService struct {
	Repo repository.Repository
	Logger *zap.Logger
}

func NewUserService(repo repository.Repository, log *zap.Logger) UserService {
	return &userService{
		Repo: repo,
		Logger: log,
	}
}

func (s *userService) GetUserByID(id int) (*model.User, error) {
	user, err := s.Repo.UserRepo.GetUserByID(id)
	if err != nil {
		s.Logger.Error("Error get user by id service: ", zap.Error(err))
	}
	return user, err
}

func (s *userService) UpdateUser(userID int, data *dto.UpdateUserRequest) error {
	uploadDir := "public/uploads"
	os.MkdirAll(uploadDir, os.ModePerm)

	// Save image file to disk
	var avatarName, avatarPath, avatarURL string
	if data.AvatarFile != nil {
		avatarName = fmt.Sprintf("%d_%s", userID, data.AvatarHeaderFile.Filename)
		avatarPath = fmt.Sprintf("%s/%s", uploadDir, avatarName)
		avatarURL = fmt.Sprintf("http://localhost:8080/%s/%s", uploadDir, avatarName)

		dst, err := os.Create(avatarPath)
		if err != nil {
			s.Logger.Error("Error create avatar filepath: ", zap.Error(err))
			return err
		}
		defer dst.Close()

		// Copy from src to dst
		_, err = io.Copy(dst, data.AvatarFile)
		if err != nil {
			s.Logger.Error("Error copy avatar file to dst: ", zap.Error(err))
			return err
		}
	}

	if avatarURL != "" {
		data.Avatar = &avatarURL
	}

	// Save image file to disk
	var CVName, CVPath, CVURL string
	if data.CVFile != nil {
		CVName = fmt.Sprintf("%d_%s", userID, data.CVHeaderFile.Filename)
		CVPath = fmt.Sprintf("%s/%s", uploadDir, CVName)
		CVURL = fmt.Sprintf("http://localhost:8080/%s/%s", uploadDir, CVName)

		dst, err := os.Create(CVPath)
		if err != nil {
			s.Logger.Error("Error create CV filepath: ", zap.Error(err))
			return err
		}
		defer dst.Close()

		// Copy from src to dst
		_, err = io.Copy(dst, data.CVFile)
		if err != nil {
			s.Logger.Error("Error copy CV file to dst: ", zap.Error(err))
			return err
		}
	}

	if CVURL != "" {
		data.CV = &CVURL
	}

	user := model.User{
		Name: data.Name,
		Email: data.Email,
		Avatar: data.Avatar,
		Description: data.Description,
		Github: data.Github,
		LinkedIn: data.LinkedIn,
		CV: data.CV,
		PhoneNumber: data.PhoneNumber,
	}
	
	err := s.Repo.UserRepo.UpdateUser(userID, &user)
	if err != nil {
		s.Logger.Error("Error update user service: ", zap.Error(err))
	}
	return err
}