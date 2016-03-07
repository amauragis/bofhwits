package bofhwitsbot

import (
	// "fmt"

	"github.com/ChimeraCoder/anaconda"
)

func (bot *BofhwitsBot) tweet(msg string) {

	anaconda.SetConsumerKey(bot.Configs.Twitter.AppAPI)
	anaconda.SetConsumerSecret(bot.Configs.Twitter.AppSecret)
	api := anaconda.NewTwitterApi(bot.Configs.Twitter.AccountAPI, bot.Configs.Twitter.AccountSecret)
	_, err := api.PostTweet(msg, nil)
	if err != nil {
		bot.con.Privmsg(bot.Configs.Channel, "Could not tweet for some reason...")
		bot.Log.Printf("Tweet Failure: %v\n", err)
	} else {

		bot.Log.Println("Tweet Success: " + msg)
	}

}

func (bot *BofhwitsBot) faketweet(msg string) {

	anaconda.SetConsumerKey(bot.Configs.Twitter.AppAPI)
	anaconda.SetConsumerSecret(bot.Configs.Twitter.AppSecret)
	api := anaconda.NewTwitterApi(bot.Configs.Twitter.AccountAPI, bot.Configs.Twitter.AccountSecret)
	_, err := api.VerifyCredentials()

	if err != nil {
		bot.Log.Printf("Fake Tweet Failure: %v\n", err)

	} else {
		bot.con.Privmsg(bot.Configs.Channel, "Would have tweeted: "+msg)
		bot.Log.Println("Fake Tweet Success: " + msg)
	}

}
