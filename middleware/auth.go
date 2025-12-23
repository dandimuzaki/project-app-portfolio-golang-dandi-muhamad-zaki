package middleware

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/dandimuzaki/project-app-portfolio-golang/model"
	"go.uber.org/zap"
)

func (mw *MiddlewareCostume) OptionalAuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("session")
			if err != nil {
				// user not logged in → continue as guest
				next.ServeHTTP(w, r)
				return
			}

			userIDStr := strings.TrimPrefix(c.Value, "user-")
			userID, err := strconv.Atoi(userIDStr)
			if err != nil {
				mw.Log.Error("Error convert cookie value: ", zap.Error(err))
			}
			// var user model.User
			u, err := mw.Service.UserService.GetUserByID(userID)
			if err != nil {
				mw.Log.Error("Error get user by id on middleware: ", zap.Error(err))
			}

			if err != nil || u == nil {
				// invalid session → continue as guest
				next.ServeHTTP(w, r)
				return
			}

			// user = *u
			ctx := context.WithValue(r.Context(), "user", u)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (mw *MiddlewareCostume) RequireAuthMiddleware() func(http.Handler) http.Handler {
  return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, ok := r.Context().Value("user").(*model.User)
			if !ok || user == nil {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
