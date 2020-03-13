package entities

// User ...
type User struct {
	Id       string `db:"Id"`
	UserName string
	Email    string
}
