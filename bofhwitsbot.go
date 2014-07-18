package main

import (
    "github.com/thoj/go-ircevent"
    "fmt"

)

// const nick = "bofhwits"
// const srv = "whatever.tld:6667"
func main() {
    
    var configs Config = LoadConfig()
    
    con := irc.IRC(configs.Nick, configs.Username)
    err := con.Connect(configs.Address)
    if err != nil {
        fmt.Println("Failed to connect")
        return
    }
    
    var roomName = configs.Channel
    
    con.AddCallback("001", func (e *irc.Event) {
        con.Join(roomName)
    })
    
    con.AddCallback("JOIN", func (e *irc.Event) {
        con.Privmsg(roomName, "Hello!")
    })
    
    con.AddCallback("PRIVMSG", func (e *irc.Event) {
        con.Privmsg(roomName, e.Message())
    })
    con.Loop()
}