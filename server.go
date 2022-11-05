package main

import (
	// "github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
	"net/http"
)

const (
	PORT = ":8888"
)

func main() {
	e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	e.GET("/health", health)
	e.POST("/world/create", createWorld)

	e.Logger.Fatal(e.Start(PORT))
}

func health(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// curl -X POST http://localhost:8888/world/create
func createWorld(c echo.Context) error {
	return c.String(http.StatusOK, "未実装")
}
