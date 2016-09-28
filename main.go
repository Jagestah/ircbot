package main

import irc "github.com/thoj/go-ircevent"
import "fmt"

var IRCserver = "irc.twitch.tv:6667"
func main() {
	con := irc.IRC("rivalrybot", "rivlarybot")
	joined := func(event *irc.Event) {
		fmt.Println(event.Nick, "has joined the channel")
		}
	initCon := func(event *irc.Event) {
		con.SendRaw("CAP REQ :twitch.tv/membership")
		con.Join("#jagestah  ")
		con.Join("#rivalrybot  ")
		return
		}
	printtt := func(event *irc.Event) {
		fmt.Println(event.Nick,":", event.Message())
		con.SendRaw("CAP REQ :twitch.tv/membership")
		return
		}
	con.Connect(IRCserver)
        con.Password = "oauth:t4wle77yx4fzfzhtmo0bfvccp01y8g"
        con.AddCallback("JOIN", joined)
	con.AddCallback("001", initCon)
	con.AddCallback("PRIVMSG", printtt)
	con.Loop()
	<-make(chan struct{})
return
}
