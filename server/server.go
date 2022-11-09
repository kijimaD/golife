package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golife/config"
	"golife/util"
	"golife/world"
	"net/http"
	"os"
)

const (
	DEFAULT_PORT = "8888"
	INPUT_FILE   = "world.txt"
)

func Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	// e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {}))

	e.GET("/", root)
	e.POST("/world/create", createWorld)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = DEFAULT_PORT
	}
	port = fmt.Sprintf(":%s", port)

	e.Logger.Fatal(e.Start(port))
}

func root(c echo.Context) error {
	return c.String(http.StatusOK, "Hi, this is golife API server. https://github.com/kijimaD/golife")
}

// curl -X POST -d $'Debug=true&GenCap=0&InitialWorld=●○○\n○○○\n○○○' http://localhost:8888/world/create
// curl -X POST -d $'Debug=true&GenCap=1&InitialWorld=●●○\n○○○\n○○○' http://kd-golife.herokuapp.com/world/create
func createWorld(con echo.Context) error {
	var params util.CreateWorldParams
	err := con.Bind(&params)
	if err != nil {
		return con.String(http.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		return con.String(http.StatusBadRequest, err.Error())
	}

	c := config.ServerLoad(params)
	h := &world.History{Configs: c}

	initialWorld := world.Load(c, params.InitialWorld)
	h.Worlds = h.CreateHistory(*initialWorld)

	json, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
	}

	return con.String(http.StatusOK, fmt.Sprintf("%s", json))
}
