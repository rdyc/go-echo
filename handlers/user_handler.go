package handlers

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/rdyc/go-echo/driver"
	"github.com/rdyc/go-echo/entities"
	repository "github.com/rdyc/go-echo/repository"
	user "github.com/rdyc/go-echo/repository/user"
)

// NewUserHandler ...
func NewUserHandler(db *driver.DB) *UserHandler {
	return &UserHandler{
		repo: user.NewSQLUserRepo(db.SQL),
	}
}

// UserHandler ...
type UserHandler struct {
	repo repository.UserRepo
}

// GetUserAll ...
func (u *UserHandler) GetUserAll(c echo.Context) error {
	users, err := u.repo.Fetch(c.Request().Context(), 10)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}

// GetUserByID ...
func (u *UserHandler) GetUserByID(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	user, err := u.repo.GetByID(c.Request().Context(), id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}

// AddUser new user
func (u *UserHandler) AddUser(c echo.Context) error {
	payload := new(entities.User)
	if err := c.Bind(payload); err != nil {
		return err
	}

	payload.Id = uuid.New()

	user, err := u.repo.Create(c.Request().Context(), payload)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, user)
}

// UpdateUser ...
func (u *UserHandler) UpdateUser(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	payload := new(entities.User)
	if err := c.Bind(payload); err != nil {
		return err
	}

	payload.Id = id

	user, err := u.repo.Update(c.Request().Context(), payload)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusAccepted, user)
}

// DeleteUser ...
func (u *UserHandler) DeleteUser(c echo.Context) error {
	id, _ := uuid.Parse(c.Param("id"))

	_, err := u.repo.Delete(c.Request().Context(), id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusAccepted, nil)
}
