package bofhwitsbot

// TODO:
// - implement commands
//    - !tweet (or whatever)
//    - ???
// - hook to a suitable microblogging platform (probably also twitter)
// - Write own IRC backend, libraries are for chumps

import (
    "github.com/thoj/go-ircevent"
    "fmt"
    "strings"
    "github.com/ChimeraCoder/anaconda"
    "log"
    "gopkg.in/yaml.v1"
    "io/ioutil"
)

// holds connection pointer, config file path, and contents of config
// populated via yaml
type BofhwitsBot struct {

    con *irc.Connection    

    Configs struct {
        Address string
        Username string
        Nick string
        Channel string
        Twitter struct {
            AppApi string
            AppSecret string
            AccountApi string
            AccountSecret string
        }
    }
    ConfigFilePath string
    Log *log.Logger
}


// populate a config struct from a yaml file.
func (bot *BofhwitsBot) LoadConfig() {

    source, err := ioutil.ReadFile(bot.ConfigFilePath)
     if err != nil {
        log.Fatal(err)
    }
    
    // literal wizard magic
    err = yaml.Unmarshal(source, &(bot.Configs))
    if err != nil {
        log.Fatal(err)
    }

}


func (bot *BofhwitsBot) tweet(msg string) {

    anaconda.SetConsumerKey(bot.Configs.Twitter.AppApi)
    anaconda.SetConsumerSecret(bot.Configs.Twitter.AppSecret)
    api := anaconda.NewTwitterApi(bot.Configs.Twitter.AccountApi, bot.Configs.Twitter.AccountSecret)
    _ , err := api.PostTweet(msg, nil)
    if err != nil{
        fmt.Println(err)
        bot.con.Privmsg(bot.Configs.Channel, "Could not tweet for some reason...")
    } else {
        
       bot.con.Privmsg(bot.Configs.Channel, "OK! Tweeted: " + msg)
    }

}

func (bot *BofhwitsBot) faketweet(msg string) {

    anaconda.SetConsumerKey(bot.Configs.Twitter.AppApi)
    anaconda.SetConsumerSecret(bot.Configs.Twitter.AppSecret)
    api := anaconda.NewTwitterApi(bot.Configs.Twitter.AccountApi, bot.Configs.Twitter.AccountSecret)
    _, err := api.VerifyCredentials()
    
    if err != nil{
        fmt.Println(err)
    
    } else {
        bot.con.Privmsg(bot.Configs.Channel, "Would have tweeted: " + msg)
    }

}



func (bot *BofhwitsBot) handleMessageEvent(e* irc.Event) {


    // list of valid commands
    msg := e.Message()
    
    // tokenize the read string, splitting it off after the first space
    token_msg := strings.SplitN(msg, " ", 2)
    cmd := token_msg[0]
    var params string

    // if we only have 1 msg, we failed to split it into two words,
    // so set params to an empty string.  We also want to trim any junk
    // space off of the params string
    if len(token_msg) > 1 {
        params = token_msg[1]
        params = strings.TrimSpace(params)
    } else{
        params = ""
    }

    // if the first letter is !, it's a real command
    if cmd[0] == '!'{
        
        // command definitions.  For readability, they are broken into
        // helper functions
        switch cmd {
            case "!tweet":
            if params != "" {
                bot.tweet(e.Nick +": " + params)
            }
            case "!tweettest":
            if params != "" {
                bot.faketweet(e.Nick +": " + params)
            }
            case "!buttes":
                bot.con.Privmsg(bot.Configs.Channel, "Donges.")
            default:
            
        }
    }   

}

// main entry point function for starting the bot.
func (bot *BofhwitsBot) RunBot() {  
    // fmt.Printf("Running with Configs:\n%v\n",bot.Configs)
    // connect to IRC
    bot.con = irc.IRC(bot.Configs.Nick, bot.Configs.Username)
    err := bot.con.Connect(bot.Configs.Address)
    if err != nil {
        fmt.Println("Failed to connect")
        return
    }
    
    // Connected to server callback
    bot.con.AddCallback("001", func (e *irc.Event) {
        bot.con.Join(bot.Configs.Channel)
    })
    
    // Join a channel callback
//     con.AddCallback("JOIN", func (e *irc.Event) {
//         con.Privmsg(roomName, "Hello!")
//     })
    
    // get a message callback
    bot.con.AddCallback("PRIVMSG", bot.handleMessageEvent)
    
    // necessary for ircevent.  Processing loop to handle all events
    bot.con.Loop()
}