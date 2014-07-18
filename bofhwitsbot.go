package main

// TODO:
// - implement commands
//    - !twat (or whatever)
//    - ???
// - hook to a suitable microblogging platform (probably also twitter)
// - Write own IRC backend, libraries are for chumps

import (
    "github.com/thoj/go-ircevent"
    "fmt"
    "flag"
)

// setup -c flag to pass a configuration file to the bot.  I suppose if you want multiple bots, you can use multiple configuration files
var config_file = flag.String("c", "config/bofhwits.yaml", "The path to the configuration file to use (default config/bofhwits.yaml)")

func main() {
    
    flag.Parse()
    
    // populate the configuration struct from the yaml conf file 
    var configs Config = LoadConfig(*config_file)
    
    // connect to IRC
    con := irc.IRC(configs.Nick, configs.Username)
    err := con.Connect(configs.Address)
    if err != nil {
        fmt.Println("Failed to connect")
        return
    }
    
    var roomName = configs.Channel
    
    // Connected to server callback
    con.AddCallback("001", func (e *irc.Event) {
        con.Join(roomName)
    })
    
    // Join a channel callback
    con.AddCallback("JOIN", func (e *irc.Event) {
        con.Privmsg(roomName, "Hello!")
    })
    
    // get a message callback
    con.AddCallback("PRIVMSG", func (e *irc.Event) {
        con.Privmsg(roomName, e.Message())
    })
    
    // necessary for ircevent.  Processing loop to handle all events
    con.Loop()
}