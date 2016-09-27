package main

import irc "github.com/thoj/go-ircevent"
import "fmt"

func main() {
	jagestah := irc.IRC("rivalrybot", "rivalrybot")
	joined := func(event *irc.Event) {
		jagestah.SendRaw("A user has joined")
		fmt.Println("A user has joined the channel")
		}
	left := func(event *irc.Event) {
		jagestah.SendRaw("A user has left")
		fmt.Println("A user has left the channel")
		}
	printtt := func(event *irc.Event) {
		fmt.Println(event.Message())
		jagestah.SendRaw("CAP REQ :twitch.tv/membership")
		jagestah.Join("#jagestah  ")
		jagestah.Join("#rivalrybot  ")
		jagestah.SendRaw("All your base R belong to me")
		return
		}
	jagestah.Connect("irc.twitch.tv:6667")
        jagestah.Password = "oauth:t4wle77yx4fzfzhtmo0bfvccp01y8g"
//	jagestah.Join("#rivalrybot  ")
//	jagestah.Join("#jagestah  ")
        jagestah.AddCallback("JOIN", joined)
	jagestah.AddCallback("PART", left)
	jagestah.AddCallback("001", printtt)
	jagestah.AddCallback("002", printtt)
	jagestah.AddCallback("PRIVMSG", printtt)
	jagestah.Loop()
	<-make(chan struct{})
return
}
