package router

import (
	"fmt"
	"net/http"
	"strconv"

	"blog/db"
	"reflect"

	"github.com/labstack/echo"
)

// Getheroes: To get heroes list or get specific hero
func Getheroes(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	heroes := db.Getheroes(id)
	fmt.Println(reflect.TypeOf(heroes))
	return c.JSON(http.StatusOK, heroes)
}

// Getheroes: To get heroes list or get specific hero
func Addheroes(c echo.Context) error {
	m := db.Hero{}
	if err := c.Bind(&m); err != nil {
		return err
	}
	newHero := db.Addheroes(m.Name)
	fmt.Println(newHero)
	return c.JSON(200, newHero)
}

// update list of heroes
func Putheroes(c echo.Context) error {
	h := []db.Hero{}
	if err := c.Bind(&h); err != nil {
		return err
	}
	newHero := db.UpdateHeroes(h)
	return c.JSON(200, newHero)
}
