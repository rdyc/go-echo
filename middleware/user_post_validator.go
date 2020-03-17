package middleware

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	"github.com/rdyc/go-echo/models"
)

// UserPostValidator ...
func UserPostValidator() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			payload := new(models.UserRequest)
			if err := c.Bind(payload); err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}

			_, err := govalidator.ValidateStruct(payload)
			if err != nil {
				return c.JSON(http.StatusBadRequest, govalidator.ErrorsByField(err))
			}

			return next(c)
		}
	}
}
