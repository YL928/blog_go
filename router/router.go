package router

import (
	"fmt"
	"net/http"
	"strconv"

	"blog/db"

	"github.com/labstack/echo"
)

// Getheroes: To get heroes list or get specific hero
func Getheroes(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	heroes := db.Getheroes(id)
	return c.JSON(http.StatusOK, heroes)
}

// Getheroes: To get heroes list or get specific hero
func Createheroes(c echo.Context) error {
	m := echo.Map{}
	if err := c.Bind(&m); err != nil {
		return err
	}
	
	return c.JSON(200, m)
}
