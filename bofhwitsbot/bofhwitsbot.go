package bofhwitsbot

import (
	// "fmt"
	"database/sql"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/amauragis/sanitize"
	"github.com/thoj/go-ircevent"
)

// holds connection pointer, config file path, and contents of config
type BofhwitsBot struct {
	con *irc.Connection

	dbOpen func(*BofhwitsBot) (*sql.DB, error)

	// configuration and logger
	// This struct is populated by YAML
	Configs struct {
		Address  string
		Username string
		Nick     string
		Channel  string

		UseTwitter bool
		DbType     string

		Twitter struct {
			AppAPI        string
			AppSecret     string
			AccountAPI    string
			AccountSecret string
		}

		Mysql struct {
			Host string
			DB   string
			User string
			Pass string
		}

		Sqlite struct {
			File string
		}
	}

	ConfigFilePath string
	Log            *log.Logger
}

// populate a config struct from a yaml file.
func (bot *BofhwitsBot) LoadConfig() error {

	if bot.Log == nil {
		bot.Log = log.New(os.Stdout, "BOFH: ", log.Ldate|log.Ltime)
	}

	source, err := ioutil.ReadFile(bot.ConfigFilePath)
	if err != nil {
		bot.Log.Printf("Read file failure: %v\n", err)
		return err
	}

	// literal wizard magic
	err = yaml.Unmarshal(source, &(bot.Configs))
	if err != nil {
		bot.Log.Printf("Unmarshal YAML failure: %v\n", err)
		return err
	}

	// DEBUG: remove me
	bot.Log.Printf("here is the configs:\n%+v\n", bot.Configs)

	return nil
}

// setup appropriate things
func (bot *BofhwitsBot) Setup() error {
	// setup database things based on different configs
	switch bot.Configs.DbType {
	case "mysql":
		bot.mysqlInit()
	case "sqlite":
		bot.sqliteInit()
	case "none":
		// TODO: disable database interaction
	default:
		bot.Log.Fatal("unsupported database")
		return errors.New("Bofh Setup: Invalid database type")
	}

	return nil
}

func (bot *BofhwitsBot) handleMessageEvent(e *irc.Event) {

	msg := strings.TrimSpace(e.Message())

	// tokenize the read string, splitting it off after the first space
	tokenMsg := strings.SplitN(msg, " ", 2)

	cmd := strings.TrimSpace(tokenMsg[0])
	var params string

	// if we only have 1 msg, we failed to split it into two words,
	// so set params to an empty string.  We also want to trim any junk
	// space off of the params string
	if len(tokenMsg) > 1 {
		params = tokenMsg[1]
		params = strings.TrimSpace(params)
	} else {
		params = ""
	}

	if len(cmd) > 1 {
		// if the first letter is !, it's a real command
		if cmd[0] == '!' {

			// command definitions.  For readability, they are broken into
			// helper functions
			switch cmd {
			// case "!tweet":
			// 	if params != "" {
			// 		bot.tweet(e.Nick + ": " + params)
			// 	}
			// case "!tweettest":
			// 	if params != "" {
			// 		bot.faketweet(e.Nick + ": " + params)
			// 	}
			case "!buttes":
				bot.con.Privmsg(bot.Configs.Channel, "Donges.")
				bot.Log.Println("Donged " + e.Nick)

			// case "!dbtest":
			// 	bot.postSql("ryzic", "test")

			// case "!unparse":
			// 	if params != "" {
			// 		user, msg := separateUsername(params)
			// 		bot.con.Privmsg(bot.Configs.Channel, "Nick: "+user)
			// 		bot.con.Privmsg(bot.Configs.Channel, "Msg: "+msg)

			// 	}

			case "!info":
				bot.Log.Println("Info requested by " + e.Nick)
				bot.con.Privmsg(bot.Configs.Channel, "bofhwits created by ryzic and comradephate // feed: https://bofh.wtf // twitter: https://twitter.com/bofhwits // use !bofhwitsdie to kill")
			case "!bofh":
				if params == "" {
					bot.con.Privmsg(bot.Configs.Channel, "Usage: !bofh <message>")
				} else if testSubmissionValidity(params) {
					bot.Log.Println("BOFH requested by " + e.Nick)
					bot.Log.Println("Msg " + params)
					requestor := e.Nick
					user, msg := separateUsername(params)
					bot.postSQL(user, sanitize.HTML(msg), requestor)
					// bot.tweet(params + " BOFH'd by " + requestor)
					bot.con.Privmsg(bot.Configs.Channel, "Okay "+e.Nick+", I posted your shitpost.")
				} else {
					bot.con.Privmsg(bot.Configs.Channel, "Hey "+e.Nick+", stop trying to break the bot (or delimit usernames better).")
					bot.Log.Printf("Delimit Failure:\n\tMsg: %v\n\tReq'd: %v\n", e.Message(), e.Nick)
				}
			case "!bofhwitsdie":
				log.Fatal("Killed by " + e.Nick)

			default:
				// no match, pretend nothing happened
			}
		}
	}

}

// main entry point function for starting the bot.
func (bot *BofhwitsBot) RunBot() {

	if bot.Log == nil {
		bot.Log = log.New(os.Stdout, "BOFH: ", log.Ldate|log.Ltime)
	}
	// connect to IRC
	bot.con = irc.IRC(bot.Configs.Nick, bot.Configs.Username)
	err := bot.con.Connect(bot.Configs.Address)
	if err != nil {
		bot.Log.Println("Failed to connect to " + bot.Configs.Address)
		return
	}

	bot.Log.Println("Connected to " + bot.Configs.Address)

	// Join our specified channel when we connect
	bot.con.AddCallback("001", func(e *irc.Event) {
		bot.con.Join(bot.Configs.Channel)
	})

	// // If we get kicked, assume it was for a good reason
	// bot.con.AddCallback("KICK", func(e *irc.Event) {
	// 	bot.Log.Fatal("Kicked!")
	// })

	// get a message callback
	bot.con.AddCallback("PRIVMSG", bot.handleMessageEvent)

	// Processing loop to handle all events
	bot.con.Loop()
}
