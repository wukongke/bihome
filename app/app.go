package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"work-codes/bihome/app/common"
	"work-codes/bihome/app/routes"
)

func StartServer() error {
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	routes.InitRoute(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	if os.Getenv("APP_MODE") == "dev" || os.Getenv("APP_MODE") == "test" {
		app.Debug = true
	}
	err := app.Start(":" + port)
	app.Logger.Fatal(err)
	return err
}

func main() {
	defer common.MgoClose()

	err := StartServer()
	if err != nil {
		fmt.Println(err)
	}
}
