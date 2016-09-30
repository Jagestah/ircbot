package main

import "github.com/jmoiron/sqlx"
import "fmt"
import _"database/sql"
import _"github.com/bmizerany/pq"

func main() {
	var users []User
	twitchNick := "rednax"
	db, err := sqlx.Connect("postgres", "host=localhost user=postgres dbname=twitchPoints password=mccork sslmode=disable parseTime=true")
		if err != nil {
			checkErr(err)
		} else {
			fmt.Println("successfully connected to DB\n")
		}
	userGet := db.Select(&users, `SELECT username, points FROM points WHERE username = $1`, twitchNick)
//	stmtUser, err := sqlUser.Exec(twitchNick)
		if userGet != nil {
			checkErr(err)
			fmt.Println(len(users))
		}
		fmt.Println(len(users))
		var user User
		if len(users) == 0 {
			fmt.Println("Creating new User: " + twitchNick)
			user.Username = twitchNick
			user.Points = 1
			insertUser, err := db.NamedExec(`INSERT INTO points (username, points) VALUES (:username, :points)`, user)
			if err != nil {
				fmt.Println("Error at InsertUser", insertUser)
				checkErr(err)
			}
		}
//		} else {
//			fmt.Println("Made it to Else")
//			user = users[0]
//		}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
//Creates a 1 to 1 association between fields and the database
type User struct {
	Username	string	`db:"username"`
	Points		int		`db:"points"`
}

