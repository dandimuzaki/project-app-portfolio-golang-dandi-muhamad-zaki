package model

type Project struct {
	Model
	UserID      int       `json:"user_id"`
	Owner       string    `json:"owner"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	URL         *string   `json:"url"`
	Image       *string   `json:"image"`
	TechStack   *[]string `json:"tech_stack"`
	IsPublished bool      `json:"is_published"`
}