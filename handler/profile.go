package handler

import (
	"html/template"
	"mime/multipart"
	"net/http"
	"strconv"

	"github.com/dandimuzaki/project-app-portfolio-golang/dto"
	"github.com/dandimuzaki/project-app-portfolio-golang/model"
	"github.com/dandimuzaki/project-app-portfolio-golang/service"
	"go.uber.org/zap"
)

type ProfileHandler struct {
	Service   service.Service
	Templates *template.Template
	Logger    *zap.Logger
}

func NewProfileHandler(service service.Service, templates *template.Template, log *zap.Logger) ProfileHandler {
	return ProfileHandler{
		Service:   service,
		Templates: templates,
		Logger:    log,
	}
}

func (h *ProfileHandler) ViewProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Get logged user data
	var userLogged *model.User
	u := r.Context().Value("user")
	if u != nil {
		userLogged = u.(*model.User)
	}

	// Get user id
	userIDStr := r.URL.Query().Get("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		h.Logger.Error("Error converting project id to int: ", zap.Error(err))
		return
	}
	user, err := h.Service.UserService.GetUserByID(userID)
	projects, err := h.Service.ProjectService.GetPersonalProjects(userID)
	if err != nil {
		h.Logger.Error("Error handling get projects by user id: ", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	embbed := display {
		User: user,
		Data: projects,
	}

	display := display{
		User: userLogged,
		Data: embbed,
	}

	// Display profile page
	if err := h.Templates.ExecuteTemplate(w, "profile", display); err != nil {
		h.Logger.Error("Error executing profile template: ", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ProfileHandler) ViewEditProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Get user
	var user *model.User
	u := r.Context().Value("user")
	if u != nil {
		user = u.(*model.User)
	}
	
	u, err := h.Service.UserService.GetUserByID(user.ID)
	if err != nil {
		h.Logger.Error("Error handling get user by id", zap.Error(err))
		return
	}

	display := display{
		User: user,
		Data: u,
	}

	// Display edit profile form
	if err := h.Templates.ExecuteTemplate(w, "profile-edit", display); err != nil {
		h.Logger.Error("Error executing profile-edit template: ", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	
	// Get user
	user := r.Context().Value("user").(*model.User)

	// Check form method
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/profile/edit", http.StatusSeeOther)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		h.Logger.Error("Error parsing multipart form: ", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get form value
	name := r.FormValue("name")
	email := r.FormValue("email")
	description := r.FormValue("description")
	github := r.FormValue("github")
	linkedin := r.FormValue("linkedin")
	phoneNumber := r.FormValue("phone")

	// Get avatar file
	var avatarFile multipart.File
	var avatarFileHeader *multipart.FileHeader
	avatarFile, avatarFileHeader, err = r.FormFile("avatar")

	if err != nil {
		if err != http.ErrMissingFile {
			h.Logger.Error("Error retrieving avatar file: ", zap.Error(err))
			http.Error(w, "Error retrieving avatar file: "+err.Error(), http.StatusBadRequest)
			return
		}
	}

	// Get cv file
	var cvFile multipart.File
	var cvFileHeader *multipart.FileHeader

	cvFile, cvFileHeader, err = r.FormFile("cv")

	if err != nil {
		if err != http.ErrMissingFile {
			h.Logger.Error("Error retrieving CV file: ", zap.Error(err))
			http.Error(w, "Error retrieving CV file: "+err.Error(), http.StatusBadRequest)
			return
		}
	}

	u := dto.UpdateUserRequest{
		Name: &name,
		Email: &email,
		Description: &description,
		AvatarFile: avatarFile,
		AvatarHeaderFile: avatarFileHeader,
		CVFile: cvFile,
		CVHeaderFile: cvFileHeader,
		Github: &github,
		LinkedIn: &linkedin,
		PhoneNumber: &phoneNumber,
	}

	err = h.Service.UserService.UpdateUser(user.ID, &u)
	if err != nil {
		h.Logger.Error("Error handling update user: ", zap.Error(err))
		http.Redirect(w, r, "/profile/edit", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/portfolio", http.StatusSeeOther)
}
