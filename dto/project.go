package dto

import "time"

type ProjectRequest struct {
	Title       string   `json:"title" validate:"required,min=10,max=200"`
	Description string   `json:"description" validate:"min=10"`
	URL         string   `json:"url"`
	Image       string   `json:"image"`
	TechStack   []string `json:"tech_stack"`
}

type ProjectResponse struct {
	ID          int        `json:"id"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
	Owner       string     `json:"owner"`
	Title       string     `json:"title"`
	Description *string     `json:"description"`
	URL         *string     `json:"url"`
	Image       *string     `json:"image"`
	TechStack   *string   `json:"tech_stack"`
}