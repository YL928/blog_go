package db

import (
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

// dbgetheroes asdadasdhasdjasdjansdasd
func Getheroes(id int) []Hero {
	heroes := []Hero{}
	if id == 0 {
		rows, err := db.Query("SELECT * FROM heroes ORDER BY id LIMIT 10")
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
