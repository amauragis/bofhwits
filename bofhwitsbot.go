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
    "strings"
)

var con *irc.Connection
var roomName string

// setup -c flag to pass a configuration file to the bot.  I suppose if you want multiple bots, you can use multiple configuration files
var config_file = flag.String("c", "config/bofhwits.yaml", "The path to the configuration file to use (default config/bofhwits.yaml)")

func HandleMessageEvent(e* irc.Event) {
    
    // list of valid commands
//     commandslist := [...]string{"!tweet","!buttes"}
    msg := e.Message()
    
    // tokenize the read string, splitting it off after the first space
    token_msg := strings.SplitN(msg, " ", 2)
    cmd := token_msg[0]
    params := token_msg[1]
    
    // if the first letter is !, it's a real command
    if cmd[0] == '!'{
        
        switch cmd {
            case "!tweet":
                //tweet()
            case "!buttes":
                con.Privmsg(roomName, "Donges.")
            default:
            
        }
    }   
    
}

func main() {
    
    flag.Parse()
    
    // populate the configuration struct from the yaml conf file 
    var configs Config = LoadConfig(*config_file)
    
    // connect to IRC
    con = irc.IRC(configs.Nick, configs.Username)
    err := con.Connect(configs.Address)
    if err != nil {
        fmt.Println("Failed to connect")
        return
    }
    
    roomName = configs.Channel
    
    // Connected to server callback
    con.AddCallback("001", func (e *irc.Event) {
        con.Join(roomName)
    })
    
    // Join a channel callback
//     con.AddCallback("JOIN", func (e *irc.Event) {
//         con.Privmsg(roomName, "Hello!")
//     })
    
    // get a message callback
    con.AddCallback("PRIVMSG", HandleMessageEvent)
    
    // necessary for ircevent.  Processing loop to handle all events
    con.Loop()
}