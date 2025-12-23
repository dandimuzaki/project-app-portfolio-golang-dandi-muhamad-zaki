package handler

import (
	"html/template"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/dandimuzaki/project-app-portfolio-golang/dto"
	"github.com/dandimuzaki/project-app-portfolio-golang/model"
	"github.com/dandimuzaki/project-app-portfolio-golang/service"
	"github.com/dandimuzaki/project-app-portfolio-golang/utils"
	"go.uber.org/zap"
)

type ProjectHandler struct {
	Service   service.Service
	Templates *template.Template
	Logger    *zap.Logger
}

func NewProjectHandler(service service.Service, templates *template.Template, log *zap.Logger) ProjectHandler {
	return ProjectHandler{
		Service: service,
		Templates: templates,
		Logger: log,
	}
}

type display struct {
	User any
	Data any
}

func (h *ProjectHandler) ViewHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Get user and projects data
	var user *model.User
	u := r.Context().Value("user")
	if u != nil {
		user = u.(*model.User)
	}
	projects, _, err := h.Service.ProjectService.GetAllProjects(-1, -1, true)
	if err != nil {
		h.Logger.Error("Error handling get all projects: ", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	display := display{
		User: user,
		Data: projects,
	}

	// Display home page
	if err := h.Templates.ExecuteTemplate(w, "home", display); err != nil {
		h.Logger.Error("Error executing home template: ", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ProjectHandler) ViewMyPortfolio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Get user and projects data
	var user *model.User
	u := r.Context().Value("user")
	if u != nil {
		user = u.(*model.User)
	}
	projects, err := h.Service.ProjectService.GetPersonalProjects(user.ID)
	if err != nil {
		h.Logger.Error("Error handling get projects by user id: ", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	display := display{
		User: user,
		Data: projects,
	}

	// Display my portfolio page
	if err := h.Templates.ExecuteTemplate(w, "my-portfolio", display); err != nil {
		h.Logger.Error("Error executing my-portfolio template: ", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ProjectHandler) ViewCreatePortfolio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Get user
	var user *model.User
	u := r.Context().Value("user")
	if u != nil {
		user = u.(*model.User)
	}

	display := display{
		User: user,
		Data: nil,
	}

	// Display create portfolio form
	if err := h.Templates.ExecuteTemplate(w, "portfolio-create", display); err != nil {
		h.Logger.Error("Error executing portfolio-create template: ", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ProjectHandler) CreatePortfolio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Get user
	var user *model.User
	u := r.Context().Value("user")
	if u != nil {
		user = u.(*model.User)
	}

	// Check form method
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/portfolio/create", http.StatusSeeOther)
		return
	}

	// Get form value
	title := r.FormValue("title")
	description := r.FormValue("description")
	url := r.FormValue("url")
	techStack := utils.StrToSlice(r.FormValue("tech_stack"))

	// Get image file
	file, fileHeader, err := r.FormFile("image")
	if err != nil {
		h.Logger.Error("Error retrieving image file: ", zap.Error(err))
		http.Error(w, "error retrieving image file", http.StatusBadRequest)
		return
	}

	project := model.Project{
		UserID: user.ID,
		Title: title,
		Description: &description,
		URL: &url,
		TechStack: &techStack,
	}

	err = h.Service.ProjectService.CreateProject(&project, file, fileHeader)
	if err != nil {
		h.Logger.Error("Error handling create project: ", zap.Error(err))
		http.Redirect(w, r, "/portfolio/create", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/portfolio", http.StatusSeeOther)
}

func (h *ProjectHandler) ViewEditPortfolio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Get user
	var user *model.User
	u := r.Context().Value("user")
	if u != nil {
		user = u.(*model.User)
	}

	// Get project by ID
	projectIDStr := r.URL.Query().Get("id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		h.Logger.Error("Error converting project id to int: ", zap.Error(err))
		http.Redirect(w, r, "/portfolio", http.StatusSeeOther)
		return
	}
	p, err := h.Service.ProjectService.GetProjectByID(projectID)
	
	var techStack string
	if p.TechStack != nil {
		techStack = strings.Join(*p.TechStack, ", ")
	}
	
	project := dto.ProjectResponse{
		ID: p.ID,
		Title: p.Title,
		Description: p.Description,
		URL: p.URL,
		Image: p.Image,
		TechStack: &techStack,
	}
	display := display{
		User: user,
		Data: project,
	}

	// Display edit portfolio form
	if err := h.Templates.ExecuteTemplate(w, "portfolio-edit", display); err != nil {
		h.Logger.Error("Error executing portfolio-edit template: ", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ProjectHandler) UpdatePortfolio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Get user
	var user *model.User
	u := r.Context().Value("user")
	if u != nil {
		user = u.(*model.User)
	}

	// Check form method
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/portfolio/edit", http.StatusSeeOther)
		return
	}

	projectIDStr := r.FormValue("project_id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		h.Logger.Error("Error converting project id to int: ", zap.Error(err))
		http.Redirect(w, r, "/portfolio/edit", http.StatusSeeOther)
		return
	}

	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		h.Logger.Error("Error parsing multipart form: ", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get form value
	title := r.FormValue("title")
	description := r.FormValue("description")
	url := r.FormValue("url")
	techStack := utils.StrToSlice(r.FormValue("tech_stack"))

	// Get image file
	var file multipart.File
	var fileHeader *multipart.FileHeader
	
	file, fileHeader, err = r.FormFile("image")

	if err != nil {
		if err != http.ErrMissingFile {
		h.Logger.Error("Error retrieving image file: ", zap.Error(err))
		http.Error(w, "Error retrieving file: " + err.Error(), http.StatusBadRequest)
		return
		}
	}

	project := model.Project{
		UserID: user.ID,
		Title: title,
		Description: &description,
		URL: &url,
		TechStack: &techStack,
	}

	err = h.Service.ProjectService.UpdateProject(projectID, &project, file, fileHeader)
	if err != nil {
		h.Logger.Error("Error handling update project: ", zap.Error(err))
		http.Redirect(w, r, "/portfolio/edit", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/portfolio", http.StatusSeeOther)
}

func (h *ProjectHandler) ViewDeletePortfolio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Get user
	var user *model.User
	u := r.Context().Value("user")
	if u != nil {
		user = u.(*model.User)
	}

	// Get project by ID
	projectIDStr := r.URL.Query().Get("id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		h.Logger.Error("Error converting project id to int: ", zap.Error(err))
		return
	}
	project, err := h.Service.ProjectService.GetProjectByID(projectID)
	display := display{
		User: user,
		Data: project,
	}

	// Display edit portfolio form
	if err := h.Templates.ExecuteTemplate(w, "portfolio-delete", display); err != nil {
		h.Logger.Error("Error executing portfolio-delete template: ", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *ProjectHandler) DeletePortfolio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Check form method
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/portfolio/delete", http.StatusSeeOther)
		return
	}

	projectIDStr := r.FormValue("project_id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		h.Logger.Error("Error converting project id to int: ", zap.Error(err))
		http.Redirect(w, r, "/portfolio/delete", http.StatusSeeOther)
		return
	}
	
	err = h.Service.ProjectService.DeleteProject(projectID)
	if err != nil {
		h.Logger.Error("Error handling delete project: ", zap.Error(err))
		http.Redirect(w, r, "/portfolio/delete", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/portfolio", http.StatusSeeOther)
}

func (h *ProjectHandler) ViewExploreDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Get user
	var user *model.User
	u := r.Context().Value("user")
	if u != nil {
		user = u.(*model.User)
	}

	// Get project by ID
	projectIDStr := r.URL.Query().Get("id")
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		h.Logger.Error("Error converting project id to int: ", zap.Error(err))
		return
	}
	project, err := h.Service.ProjectService.GetProjectByID(projectID)
	
	display := display{
		User: user,
		Data: project,
	}

	// Display edit portfolio form
	if err := h.Templates.ExecuteTemplate(w, "explore-details", display); err != nil {
		h.Logger.Error("Error executing explore-details template: ", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
