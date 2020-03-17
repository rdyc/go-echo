package models

// UserRequest model
type UserRequest struct {
	Name  *string `json:"name" xml:"name" form:"name" query:"name" valid:"required,alpha~alpha a-z or A-Z only"`
	Email *string `json:"email" xml:"email" form:"email" query:"email" valid:"required,email"`
}
