package main

import (
    "github.com/thoj/go-ircevent"
    "fmt"
)

const nick = "bofhwits"
const srv = "whatever.tld:6667"
func main() {
    con := irc.IRC(nick, nick)
    err := con.Connect(srv)
    if err != nil {
        fmt.Println("Failed to connect")
        return
    }
    
    var roomName = "#ryzic_dumb_junk"
    
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