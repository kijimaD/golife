package server

import (
	// "github.com/labstack/echo/middleware"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"golife/config"
	"golife/world"
	"net/http"
)

const (
	PORT = ":8888"
)

func Run() {
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
func createWorld(con echo.Context) error {
	h := &world.History{}
	c := config.Load()
	initialWorld := world.Load(c)
	h.Worlds = h.CreateHistory(*initialWorld, c)

	jsonData, err := json.Marshal(h.Worlds)
	if err != nil {
		fmt.Println(err)
	}

	return con.String(http.StatusOK, fmt.Sprintf("%s", jsonData))
}
