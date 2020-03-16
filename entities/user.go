package entities

import "github.com/google/uuid"

// User ...
type User struct {
	Id       uuid.UUID `json:"_id"`
	UserName string    `json:"user_name"`
	Email    string    `json:"email"`
}
