package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)
import "blog/router"

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:4200"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.GET("/api/heroes", router.Getheroes)
	e.GET("/api/heroes/:id", router.Getheroes)
	e.POST("/api/heroes", router.Createheroes)
	e.Logger.Fatal(e.Start(":3000"))
}
