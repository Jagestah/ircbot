package main

import irc "github.com/thoj/go-ircevent"
import "fmt"
//import "time"
import "io/ioutil"
import "strings"

var IRCserver = "irc.twitch.tv:6667"
func main() {
	con := irc.IRC("rivalrybot", "rivlarybot")	//The name of the bot's twitch account
	joined := func(event *irc.Event) {
		fmt.Println(event.Nick, "has joined the channel")	//prints user and text to terminal when someone joins
		}
	initCon := func(event *irc.Event) {
		con.SendRaw("CAP REQ :twitch.tv/membership")	//requests JOIN and other info from Twitch servers
		con.Join("#jagestah  ")	//joins chat in jagestah channel
		con.Join("#rivalrybot  ")	//joins chat in rivalrybot channel - necessary for stability
		return
		}
	printtt := func(event *irc.Event) {
		fmt.Println(event.Nick,":", event.Message())	//prints user and their message in chat into the terminal
		return
		}
	con.Connect(IRCserver)
        con.Password = "oauth:t4wle77yx4fzfzhtmo0bfvccp01y8g" //password for rivalrybot
        con.AddCallback("JOIN", joined)	//event for when a user joins
	con.AddCallback("001", initCon)	//event for when receiving welcome message from Twitch
	con.AddCallback("PRIVMSG", printtt)	//prints twitch chat to terminal
////////////////////////////

//	var text string
//	fmt.Print("Enter text: ")
//	fmt.Scanln(&text)
//	
//	b, err := ioutil.ReadFile("points")
//		if err != nil {
//			panic(err)
//		}
//	s := string(b)
//	fmt.Println(strings.Contains(s, text))
//
//	pointsTimer := func AfterFunc(d 60 * Second, f func(
//		event.Nick
//	con.AddCallback("JOIN", pointsTimer

////////////////////////////
	con.Loop()
	<-make(chan struct{})
return
}
