package routers

import (
	"github.com/labstack/echo"
	"github.com/rdyc/go-echo/driver"
	"github.com/rdyc/go-echo/handlers"
)

// UserRouter ...
func UserRouter(g *echo.Group, db *driver.DB) {
	handler := handlers.NewUserHandler(db)

	v1 := g.Group("/v1")

	v1.GET("/users", handler.GetUserAll)
	v1.GET("/users/:id", handler.GetUserByID)
	v1.POST("/users", handler.AddUser)
}
