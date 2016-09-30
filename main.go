package main

import irc "github.com/thoj/go-ircevent"
import "fmt"
import "github.com/fatih/color"
import "strings"
import "github.com/jmoiron/sqlx"
import _"database/sql"
import _"github.com/bmizerany/pq"
import "math/rand"


//Creates a 1 to 1 association between fields and the database
type User struct {
	Username	string	`db:"username"`
	Points		int		`db:"points"`
}

type Quote struct {
	Qid			int 	`db:"qid"`
	Addedby		string	`db:"addedby"`
	Channel		string 	`db:"channel"`
	Quote 		string 	`db:"quote"`
}

var IRCserver = "irc.twitch.tv:6667"
var users []User
var con = irc.IRC("tyrannicalbot", "tyrannicalbot")	//The name of the bot's twitch account


func main() {
	con.Connect(IRCserver)
        con.Password = "oauth:1a8oxng9fiar2vq9rojhnj3fus14dx" //password for rivalrybot
    con.AddCallback("JOIN", joined)	//event for when a user joins
	con.AddCallback("001", initCon)	//event for when receiving welcome message from Twitch
	con.AddCallback("PRIVMSG", printtt,)	//prints twitch chat to terminal
	con.AddCallback("PRIVMSG", cmdCheck)
	con.Loop()
	<-make(chan struct{})
return
}

func initCon(event *irc.Event) {
	con.SendRaw("CAP REQ :twitch.tv/membership")	//requests JOIN and other info from Twitch servers
	con.Join("#jagestah  ")	//joins chat in jagestah channel
	con.Join("#rivalrybot  ")
	con.Join("#tyrannicalbot  ")	//joins chat in rivalrybot channel - necessary for stability
	return
}
func cmdCheck(event *irc.Event) {
	var channel = event.Arguments[0]
	var quotes []Quote
	var quote Quote
	db, err := sqlx.Connect("postgres", "host=localhost user=postgres dbname=twitchPoints password=mccork sslmode=disable parseTime=true")
	eventMessage := string(event.Message())
	args := strings.Split(eventMessage, " ")

	if strings.Contains(eventMessage, "!wut") == true { //checks chat messages for !wut command
		con.Privmsg(channel, "You dun did it.")
		fmt.Println("Sending command for !wut")
		return
	} else if strings.Contains(eventMessage, "!quote") == true {
		var channel = event.Arguments[0]
		if len(args) == 1 {
				var count int
				quoteCount := db.Get(&count, `SELECT count(*) FROM quotes`)
				_ = quoteCount
					if err != nil {
						checkErr(err)
					}
				count = rand.Intn(count) - 1
				fmt.Println((count) + 1)
		} else if len(args) == 2 {
 			if strings.ContainsAny(args[1], "1 & 2 & 3 & 4 & 5 & 6 & 7 & 8 & 9")  == true {
				quoteGet := db.Select(&quotes, `SELECT quote FROM quotes WHERE qid = $1 AND channel = $2`, args[1], channel)
					if len(quotes) > 0 {
					fmt.Sprint(quoteGet)
					quoteOut := quotes[0].Quote
					fmt.Println("Sending,",quoteOut,"to channel for !quote cmd")
					con.Privmsg(channel, quoteOut)
					quote = quotes[0]
					return
					} else {
					fmt.Println("Quote not found")
					con.Privmsg(channel, "Quote not found.")
					return
					}
				}  else if strings.Contains(args[1], "?") {
				fmt.Println("Sending !quote syntax")
				con.Privmsg(channel, `For quotes use !quote <#> or !quote "add"`)
				return
			}
		}
		if len(args) > 2 {

				if strings.Contains(args[1], "add") {		
						if err != nil {
						checkErr(err)
						return
						}
					quoteSplit := strings.Split(eventMessage, "!quote add ")
					quote.Addedby = event.Nick
					quote.Channel = channel
					quote.Quote = quoteSplit[1]
					insertQuote, err := db.NamedExec(`INSERT INTO quotes (addedby, channel, quote) VALUES (:addedby, :channel, :quote)`, quote)
					_ = insertQuote
						if err != nil {
						fmt.Println("Failed to add quote.\n")
						fmt.Println(err)
						return
					} else if len(quotes) > 0 {
							getQid := db.Select(&quotes, `SELECT qid FROM quotes WHERE quote = $1`, quoteSplit[1])
							addReply := fmt.Sprint("Added quote '",quoteSplit[1],"' with quote ID ",quotes[0].Qid)
								if err != nil {
									checkErr(err)
								}
							fmt.Println(addReply)
							con.Privmsg(channel, addReply)
							fmt.Println(getQid)
							return
							}	
					}
				}
			}
	}

func printtt(event *irc.Event) {
	color.Set(color.FgYellow)
	var channel = event.Arguments[0]
	fmt.Println(channel,".",event.Nick,":", event.Message())	//prints user and their message in chat into the terminal
	color.Unset()
	return
}


func joined(event *irc.Event) {
	color.Set(color.FgCyan)
	var channel = event.Arguments[0]
	var users []User
	fmt.Println(event.Nick, "has joined", channel)	//prints user and text to terminal when someone joins
	db, err := sqlx.Connect("postgres", "host=localhost user=postgres dbname=twitchPoints password=mccork sslmode=disable parseTime=true")
	if err != nil {
		checkErr(err)
	}
	userGet := db.Select(&users, `SELECT username, points FROM points WHERE username = $1`, event.Nick)
	if userGet != nil {
		checkErr(err)
		fmt.Println("Error at userGet")
	}
	var user User
	if len(users) == 0 {
		fmt.Println("Creating new User: " + event.Nick)
		user.Username = event.Nick
		user.Points = 1
		insertUser, err := db.NamedExec(`INSERT INTO points (username, points) VALUES (:username, :points)`, user)
		if err != nil {
			fmt.Println("Error at InsertUser", insertUser)
			checkErr(err)
		}
	} else {
//		fmt.Println("User already exists in DB")
		user = users[0]
	}
	color.Unset()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}