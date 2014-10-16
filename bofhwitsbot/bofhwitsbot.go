package bofhwitsbot

import (
	// "fmt"
	"database/sql"
	"github.com/ChimeraCoder/anaconda"
	"github.com/amauragis/sanitize"
	_ "github.com/go-sql-driver/mysql"
	"github.com/thoj/go-ircevent"
	"gopkg.in/yaml.v1"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// holds connection pointer, config file path, and contents of config
// populated via yaml
type BofhwitsBot struct {
	con *irc.Connection

	Configs struct {
		Address  string
		Username string
		Nick     string
		Channel  string
		Twitter  struct {
			AppApi        string
			AppSecret     string
			AccountApi    string
			AccountSecret string
		}
		Mysql struct {
			Host string
			DB   string
			User string
			Pass string
		}
	}
	ConfigFilePath string
	Log            *log.Logger
}

// populate a config struct from a yaml file.
func (bot *BofhwitsBot) LoadConfig() error {

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

	return nil

}

// format the string (probably a tweet request) for a single line
func formatTextOneLine(s string) string {
	return strings.Replace(s, "\n", " | ", -1)
}

// format the string (log request for site) preserving line endings
func formatTextMultiLine(s string) string {
	return s
}

func separateUsername(s string) (user string, msg string) {

	// search for a < and > pair and take everything between them
	if strings.ContainsRune(s, '<') {
		userStart := strings.IndexRune(s, '<') + 1
		userEnd := strings.IndexRune(s, '>')

		user = strings.TrimSpace(s[userStart:userEnd])
		msg = strings.TrimSpace(s[userEnd+1:])

	} else {
		// handle single ended delimiters
		delims := []rune{'>', ':', ','}
		var matchRune rune
		for _, elem := range delims {

			if strings.ContainsRune(s, elem) {
				matchRune = elem
			}
		}

		quotedRuneStr := strconv.QuoteRuneToASCII(matchRune)
		runeStr := quotedRuneStr[1 : len(quotedRuneStr)-1]
		splitstr := strings.Split(s, runeStr)

		user = strings.TrimSpace(splitstr[0])
		msg = strings.TrimSpace(splitstr[1])
	}

	return
}

func (bot *BofhwitsBot) dbInit() {
	sqlcon, oerr := sql.Open("mysql", bot.Configs.Mysql.User+":"+bot.Configs.Mysql.Pass+"@tcp("+bot.Configs.Mysql.Host+":3306)/"+bot.Configs.Mysql.DB)
	if oerr != nil {
		bot.Log.Printf("DB Failure: %v\n", oerr)
	}
	defer sqlcon.Close()

	_, execErr :=
		sqlcon.Exec("CREATE TABLE IF NOT EXISTS bofhwits_posts ( post_id int PRIMARY KEY AUTO_INCREMENT, user VARCHAR(50), post VARCHAR(1000), requestor VARCHAR(50), ts TIMESTAMP);")

	if execErr != nil {
		bot.Log.Printf("Failed to init database: %v\n", execErr)
	}
}

func (bot *BofhwitsBot) postSql(user string, msg string, requestor string) {
	sqlcon, oerr := sql.Open("mysql", bot.Configs.Mysql.User+":"+bot.Configs.Mysql.Pass+"@tcp("+bot.Configs.Mysql.Host+":3306)/"+bot.Configs.Mysql.DB)
	if oerr != nil {
		bot.con.Privmsg(bot.Configs.Channel, "Could not connect db for some reason...")
		bot.Log.Printf("DB Failure: %v\n", oerr)
	}
	defer sqlcon.Close()

	bot.Log.Printf("DB: INSERT INTO bofhwits_posts (user, post, requestor) VALUES (%s, %s, %s)", user, msg, requestor)

	stmt, err := sqlcon.Prepare("INSERT INTO bofhwits_posts (user, post, requestor) VALUES (?, ?, ?)")
	stmt.Exec(user, msg, requestor)

	if err != nil {
		bot.con.Privmsg(bot.Configs.Channel, "Could not db for some reason...")
		bot.Log.Printf("DB Failure: %v\n", err)
	}
}

func (bot *BofhwitsBot) tweet(msg string) {

	anaconda.SetConsumerKey(bot.Configs.Twitter.AppApi)
	anaconda.SetConsumerSecret(bot.Configs.Twitter.AppSecret)
	api := anaconda.NewTwitterApi(bot.Configs.Twitter.AccountApi, bot.Configs.Twitter.AccountSecret)
	_, err := api.PostTweet(msg, nil)
	if err != nil {
		bot.con.Privmsg(bot.Configs.Channel, "Could not tweet for some reason...")
		bot.Log.Printf("Tweet Failure: %v\n", err)
	} else {

		// bot.con.Privmsg(bot.Configs.Channel, "OK! Tweeted: "+msg)
		bot.Log.Println("Tweet Success: " + msg)
	}

}

func (bot *BofhwitsBot) faketweet(msg string) {

	anaconda.SetConsumerKey(bot.Configs.Twitter.AppApi)
	anaconda.SetConsumerSecret(bot.Configs.Twitter.AppSecret)
	api := anaconda.NewTwitterApi(bot.Configs.Twitter.AccountApi, bot.Configs.Twitter.AccountSecret)
	_, err := api.VerifyCredentials()

	if err != nil {
		bot.Log.Printf("Fake Tweet Failure: %v\n", err)

	} else {
		bot.con.Privmsg(bot.Configs.Channel, "Would have tweeted: "+msg)
		bot.Log.Println("Fake Tweet Success: " + msg)
	}

}

func testSubmissionValidity(s string) bool {
	delims := []rune{'>', ':', ','}
	exists := false
	for _, elem := range delims {
		if strings.ContainsRune(s, elem) {
			exists = true
		}
	}
	return exists
}

func (bot *BofhwitsBot) handleMessageEvent(e *irc.Event) {

	msg := strings.TrimSpace(e.Message())

	// tokenize the read string, splitting it off after the first space
	token_msg := strings.SplitN(msg, " ", 2)

	cmd := strings.TrimSpace(token_msg[0])
	var params string

	// if we only have 1 msg, we failed to split it into two words,
	// so set params to an empty string.  We also want to trim any junk
	// space off of the params string
	if len(token_msg) > 1 {
		params = token_msg[1]
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
					bot.postSql(user, sanitize.HTML(msg), requestor)
					bot.tweet(params + " BOFH'd by " + requestor)
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

	// bot.dbInit()

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
