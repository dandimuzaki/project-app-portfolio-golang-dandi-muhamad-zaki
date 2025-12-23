package model

type User struct {
	Model
	Name        *string `json:"name"`
	Email       *string `json:"email"`
	Password    *string `json:"password"`
	Avatar      *string `json:"avatar"`
	Description *string `json:"description"`
	Github      *string `json:"github"`
	CV          *string `json:"cv"`
	LinkedIn    *string `json:"linkedin"`
	PhoneNumber *string `json:"phone_number"`
}