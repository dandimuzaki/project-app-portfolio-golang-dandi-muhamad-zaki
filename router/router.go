package router

import (
	"net/http"

	"github.com/dandimuzaki/project-app-portfolio-golang/handler"
	"github.com/dandimuzaki/project-app-portfolio-golang/middleware"
	"github.com/dandimuzaki/project-app-portfolio-golang/service"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func NewRouter(service service.Service, handler handler.Handler, log *zap.Logger) *chi.Mux {
	r := chi.NewRouter()

	mw := middleware.NewMiddlewareCustome(service, log)
	
	r.Mount("/", ApiV1(handler, mw, log))

	fs := http.FileServer(http.Dir("public"))
	r.Handle("/public/*", http.StripPrefix("/public/", fs))

	return r
}

func ApiV1(handler handler.Handler, mw middleware.MiddlewareCostume, log *zap.Logger) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Use(mw.OptionalAuthMiddleware())

		r.Get("/home", handler.ProjectHandler.ViewHome)

		// Authentication
		r.Get("/login", handler.AuthHandler.ViewLogin)
		r.Get("/register", handler.AuthHandler.ViewRegister)
		r.Get("/logout", handler.AuthHandler.ViewLogout)
		r.Post("/login", handler.AuthHandler.Login)
		r.Post("/register", handler.AuthHandler.Register)
		r.Post("/logout", handler.AuthHandler.Logout)
	})

	r.Route("/explore", func(r chi.Router) {
		r.Use(mw.OptionalAuthMiddleware())
		r.Get("/", handler.ProfileHandler.ViewProfile)
		r.Get("/details", handler.ProjectHandler.ViewExploreDetails)
	})

	r.Route("/profile", func(r chi.Router) {
		r.Use(mw.OptionalAuthMiddleware())
		r.Use(mw.RequireAuthMiddleware())
		r.Get("/edit", handler.ProfileHandler.ViewEditProfile)
		r.Post("/edit", handler.ProfileHandler.UpdateProfile)
	})

	r.Route("/portfolio", func(r chi.Router) {
		r.Use(mw.OptionalAuthMiddleware())
		r.Use(mw.RequireAuthMiddleware())

		// Project
		r.Get("/", handler.ProjectHandler.ViewMyPortfolio)
		r.Get("/create", handler.ProjectHandler.ViewCreatePortfolio)
		r.Post("/create", handler.ProjectHandler.CreatePortfolio)
		r.Get("/edit", handler.ProjectHandler.ViewEditPortfolio)
		r.Post("/edit", handler.ProjectHandler.UpdatePortfolio)
		r.Get("/delete", handler.ProjectHandler.ViewDeletePortfolio)
		r.Post("/delete", handler.ProjectHandler.DeletePortfolio)
	})

	return r
}