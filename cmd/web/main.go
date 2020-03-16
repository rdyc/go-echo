package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/rdyc/go-echo/driver"
	"github.com/rdyc/go-echo/routers"
)

// Config is configuration for Server
type Config struct {
	// BasePath is base path for public folders
	BasePath string
}

func main() {
	// get configuration
	var cfg Config
	flag.StringVar(&cfg.BasePath, "base-path", "", "base path for public folders")
	flag.Parse()

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

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

	// web resources
	e.File("/", cfg.BasePath+"public/index.html")
	e.Static("/assets", cfg.BasePath+"public/assets")

	// api group
	v1 := e.Group("/api")

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
		port = "8080"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	} else {
		fmt.Println("INFO: Heroku environment port detected to " + port)
	}
	return ":" + port
}
