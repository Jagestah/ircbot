package main

import "github.com/jmoiron/sqlx"
//import "fmt"
import _ "database/sql"
import _ "github.com/bmizerany/pq"

func main() {
	db, err := sqlx.Connect("postgres", "host=localhost user=postgres dbname=twitchPoints password=mccork sslmode=disable parseTime=true")
	checkErr(err)
	defer db.Close()
}
func UserGet() User {
	
username := "rednax"
	var users []User
	//fmt.Println(discordUser.ID)
	err := db.Select(&points, WHERE username = $1)
	if err != nil {
		log.Fatal(err)
	}
	var user User
	if len(points) == 0 {
		fmt.Println("creating user: " + username)
//		createUser(discordUser, db)
//		user = UserGet(discordUser, db)
	} else {
		user = users[0]
	}
	return user
}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
