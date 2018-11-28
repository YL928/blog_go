package db

import (
	"fmt"
	_ "fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

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

// update list of heroes
func UpdateHeroes(heroes []Hero) []Hero {
	newheroes := []Hero{}
	for h := range heroes {
		fmt.Println(heroes)
		fmt.Println(h)
		fmt.Println("=======")
		// row, err := db.Exec("UPDATE heroes set name = ? where id = ?", h.Name, h.ID)
		// checkErr(err)
		// row, err := res.RowsAffected()
		// checkErr(err)
		// if row.RowsAffected() == 1 {
		// 	newheroes = append(newheroes, Hero{h.ID, h.Name})
		// }

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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
