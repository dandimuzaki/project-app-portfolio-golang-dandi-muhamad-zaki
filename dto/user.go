package dto

import "mime/multipart"

type UpdateUserRequest struct {
	Name        *string  `json:"name"`
	Email       *string  `json:"email"`
	Password    *string  `json:"password"`
	Avatar      *string `json:"avatar"`
	AvatarFile  multipart.File
	AvatarHeaderFile *multipart.FileHeader
	Description *string `json:"description"`
	Github      *string `json:"github"`
	CV          *string `json:"cv"`
	CVFile multipart.File
	CVHeaderFile *multipart.FileHeader
	LinkedIn    *string `json:"linkedin"`
	PhoneNumber *string `json:"phone_number"`
}