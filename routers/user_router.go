package routers

import (
	"github.com/labstack/echo"
	"github.com/rdyc/go-echo/driver"
	"github.com/rdyc/go-echo/handlers"
)

// UserRouter ...
func UserRouter(g *echo.Group, db *driver.DB) {
	handler := handlers.NewUserHandler(db)

	g.GET("/users", handler.GetUserAll)
	// g.GET("/users/:id", pHandler.GetUserByID)
	// g.POST("/users", pHandler.AddUser)
}
