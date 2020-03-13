package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/rdyc/go-echo/driver"
	"github.com/rdyc/go-echo/routers"
)

const (
	host   = "localhost"
	port   = "5432"
	user   = "appuser"
	pass   = "P@ssw0rd!"
	dbname = "IdentityServer4Admin"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	if dbHost == "" {
		dbHost = host
	}
	if dbPort == "" {
		dbPort = port
	}
	if dbName == "" {
		dbName = dbname
	}
	if dbUser == "" {
		dbUser = user
	}
	if dbPass == "" {
		dbPass = pass
	}

	db, err := driver.ConnectSQL(dbHost, dbPort, dbUser, dbPass, dbName)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	// create new echo instance
	e := echo.New()

	// register middlewares
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// web router
	e.File("/", "public/index.html")
	e.Static("/assets", "public/assets")

	// api router group
	v1 := e.Group("/v1")

	// register routers with group
	routers.UserRouter(v1, db)

	// error handler on startup
	e.Logger.Fatal(e.Start(getPort()))
}

// getPort ...
// The Port from the environment so we can run on Heroku
func getPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4000"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	} else {
		fmt.Println("INFO: Heroku environment port detected to " + port)
	}
	return ":" + port
}
