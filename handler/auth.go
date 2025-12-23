package handler

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/dandimuzaki/project-app-portfolio-golang/model"
	"github.com/dandimuzaki/project-app-portfolio-golang/service"
	"go.uber.org/zap"
)

type AuthHandler struct {
	Service service.Service
	Templates   *template.Template
	Logger      *zap.Logger
}

func NewAuthHandler(service service.Service, templates *template.Template, log *zap.Logger) AuthHandler {
	return AuthHandler{
		Service: service,
		Templates: templates,
		Logger: log,
	}
}

func (h *AuthHandler) ViewRegister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Display register page
	if err := h.Templates.ExecuteTemplate(w, "register", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Method should be post
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	// Get input field from form
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	var data model.User
	data.Name = &name
	data.Email = &email
	data.Password = &password

	// Execute sign in
	user, err := h.Service.AuthService.SignIn(data)
	if err != nil {
		h.Templates.ExecuteTemplate(w, "register", map[string]interface{}{
			"Error": err.Error(),
		})
		return
	}

	// Set cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "user-" + strconv.Itoa(user.ID),
		Path:     "/",
		HttpOnly: true,
	})

	// Redirect to home page
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func (h *AuthHandler) ViewLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Display login page
	if err := h.Templates.ExecuteTemplate(w, "login", nil); err != nil {
		h.Logger.Error("Error executing login template: ", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Method should be post
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Get input field from form
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Execute login
	user, err := h.Service.AuthService.Login(email, password)
	if err != nil {
		h.Logger.Error("Error handling login user: ", zap.Error(err))
		h.Templates.ExecuteTemplate(w, "login", map[string]interface{}{
			"Error": err.Error(),
		})
		return
	}

	// Set cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "user-" + strconv.Itoa(user.ID),
		Path:     "/",
		HttpOnly: true,
	})

	// Redirect to home page
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func (h *AuthHandler) ViewLogout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Display logout confirmation
	if err := h.Templates.ExecuteTemplate(w, "logout", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Reset cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	// Redirect to login page
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}
