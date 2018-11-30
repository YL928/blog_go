package db

import (
	"fmt"
	_ "fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Hero struct
type Hero struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hero")
	checkErr(err)
	db = database
}

// create new hero
func Addheroes(name string) Hero {
	row, err := db.Exec("INSERT INTO heroes (name) Values (?)", name)
	checkErr(err)
	id, err := row.LastInsertId()
	checkErr(err)
	return Hero{ID: int(id), Name: name}
}

func DeleteHeroes(id int) int {
	res, err := db.Exec("DELETE FROM heroes WHERE id = ?", id)
	checkErr(err)
	row, err := res.RowsAffected()
	checkErr(err)
	return int(row)
}

// update list of heroes
func UpdateHeroes(heroes []Hero) []Hero {
	newheroes := []Hero{}
	for _, h := range heroes {
		fmt.Println(heroes)
		fmt.Println(h)
		fmt.Println("=======")
		row, err := db.Exec("UPDATE heroes set name = ? where id = ?", h.Name, h.ID)
		checkErr(err)
		changeRow, err := row.RowsAffected()
		checkErr(err)
		if changeRow == 1 {
			newheroes = append(newheroes, Hero{h.ID, h.Name})
		}
	}
	return newheroes
}

// get heroes or get particular one
func Getheroes(id int) []Hero {
	heroes := []Hero{}
	if id == 0 {
		rows, err := db.Query("SELECT * FROM heroes ORDER BY id DESC LIMIT 10")
		checkErr(err)
		for rows.Next() {
			var uid int
			var name string
			err = rows.Scan(&uid, &name)
			checkErr(err)
			hero := Hero{ID: uid, Name: name}
			heroes = append(heroes, hero)
		}
		defer rows.Close()

	} else {
		rows, err := db.Query("SELECT * FROM heroes WHERE id = ?", id)
		checkErr(err)
		for rows.Next() {
			var uid int
			var name string
			err = rows.Scan(&uid, &name)
			checkErr(err)
			hero := Hero{ID: uid, Name: name}
			heroes = append(heroes, hero)
		}
		defer rows.Close()
	}
	return heroes
}

// asdasdasdasdas
func Gethero(q map) []Hero {
	fmt.println(q)
	heroes := []Hero{}
	hero := Hero{1, "hhh"}
	heroes = append(heroes, hero)
	return heroes
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
