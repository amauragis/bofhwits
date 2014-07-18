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
    "github.com/ChimeraCoder/anaconda"
)

var con *irc.Connection
var roomName string
var configs Config

// setup -c flag to pass a configuration file to the bot.  I suppose if you want multiple bots, you can use multiple configuration files
var config_file = flag.String("c", "config/bofhwits.yaml", "The path to the configuration file to use (default config/bofhwits.yaml)")

func tweet(msg string) {
    anaconda.SetConsumerKey(configs.Twitter.Appapi)
    anaconda.SetConsumerSecret(configs.Twitter.Appsecret)
    api := anaconda.NewTwitterApi(configs.Twitter.Accountapi, configs.Twitter.Accountsecret)
    _ , err := api.PostTweet(msg, nil)
    if err != nil{
        fmt.Println(err)
        con.Privmsg(roomName, "Could not tweet for some reason...")
    } else {
        
       con.Privmsg(roomName, "OK! Tweeted: " + msg)
    }
}

func faketweet(msg string) {
    anaconda.SetConsumerKey(configs.Twitter.Appapi)
    anaconda.SetConsumerSecret(configs.Twitter.Appsecret)
    api := anaconda.NewTwitterApi(configs.Twitter.Accountapi, configs.Twitter.Accountsecret)
    ok, err := api.VerifyCredentials()
    
    if err != nil{
        fmt.Println(err)

    
    } else {
        fmt.Printf("Success: %#v \n", ok)    
        con.Privmsg(roomName, "Would have tweeted: " + msg)
    }
}



func HandleMessageEvent(e* irc.Event) {
    
    // list of valid commands
//     commandslist := [...]string{"!tweet","!buttes"}
    msg := e.Message()
    
    // tokenize the read string, splitting it off after the first space
    token_msg := strings.SplitN(msg, " ", 2)
    cmd := token_msg[0]
    var params string
    if len(token_msg) > 1 {
        params = token_msg[1]
        params = strings.TrimSpace(params)
    } else{
        params = ""
    }
    // if the first letter is !, it's a real command
    if cmd[0] == '!'{
        
        switch cmd {
            case "!tweet":
            if params != "" {
                tweet(e.Nick +": " + params)
            }
            case "!tweettest":
            if params != "" {
                faketweet(e.Nick +": " + params)
            }
            case "!buttes":
                con.Privmsg(roomName, "Donges.")
            default:
            
        }
    }   
    
}

func main() {
    
    flag.Parse()
    
    // populate the configuration struct from the yaml conf file 
    configs = LoadConfig(*config_file)
    // this could also say configs := LoadConfig(*config_file)
    
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