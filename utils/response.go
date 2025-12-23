package utils

import (
	"encoding/json"
	"net/http"

	"github.com/dandimuzaki/project-app-portfolio-golang/dto"
)

func ResponseWithPagination(w http.ResponseWriter, code int, message string, data any, pagination *dto.Pagination) {
	response := map[string]interface{}{
		"status":     true,
		"message":    message,
		"data":       data,
		"pagination": &pagination,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}