package handler

import (
	"html/template"

	"github.com/dandimuzaki/project-app-portfolio-golang/service"
	"go.uber.org/zap"
)

type Handler struct {
	AuthHandler    AuthHandler
	ProjectHandler ProjectHandler
	ProfileHandler ProfileHandler
}

func NewHandler(service service.Service, templates *template.Template, log *zap.Logger) Handler {
	return Handler{
		AuthHandler: NewAuthHandler(service, templates, log),
		ProjectHandler: NewProjectHandler(service, templates, log),
		ProfileHandler: NewProfileHandler(service, templates, log),
	}
}