package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/rdyc/go-echo/driver"
	repository "github.com/rdyc/go-echo/repository"
	user "github.com/rdyc/go-echo/repository/user"
)

// NewUserHandler ...
func NewUserHandler(db *driver.DB) *UserHandler {
	return &UserHandler{
		repo: user.NewSQLPostRepo(db.SQL),
	}
}

// UserHandler ...
type UserHandler struct {
	repo repository.UserRepo
}

// GetUserAll ...
func (u *UserHandler) GetUserAll(c echo.Context) error {
	users, err := u.repo.Fetch(c.Request().Context(), 10)

	// users := []models.UserResponse{
	// 	{
	// 		Name:  "Dr. Bruce Banner",
	// 		Email: "hulk@avenger.net",
	// 	},
	// 	{
	// 		Name:  "Tony Stark",
	// 		Email: "ironman@avenger.net",
	// 	},
	// 	{
	// 		Name:  "Thor Odinson",
	// 		Email: "thor@avenger.net",
	// 	},
	// }

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}

// // GetUserByID ...
// func (u *UserHandler) GetUserByID(c echo.Context) error {
// 	id := c.Param("id")

// 	return c.String(http.StatusOK, id)
// }

// // AddUser new user
// func (u *UserHandler) AddUser(c echo.Context) error {
// 	u := new(models.UserRequest)
// 	if err := c.Bind(u); err != nil {
// 		return err
// 	}

// 	return c.JSON(http.StatusCreated, u)
// 	// or
// 	// return c.XML(http.StatusCreated, u)
// }
