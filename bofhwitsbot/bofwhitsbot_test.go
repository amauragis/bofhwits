package bofhwitsbot

import (
	"github.com/ChimeraCoder/anaconda"
	"os"
	"testing"
)

var bot BofhwitsBot
var api *anaconda.TwitterApi

func init() {
	// Initialize so we can actually run tests
	bot = BofhwitsBot{ConfigFilePath: "../config/bofhwits.yaml"}

}

// test that we can configure the default config file
func Test_ConfigParse(t *testing.T) {
	if _, err := os.Stat("../config/bofhwits.yaml"); os.IsNotExist(err) {
		t.Fatalf("../config/bofhwits.yaml not found")
	}

	err := bot.LoadConfig()
	if err != nil {
		t.Errorf("Configuration: %v\n", err)
	}
}

// test that the twitter API is working correctly
func Test_TwitterApi(t *testing.T) {

	anaconda.SetConsumerKey(bot.Configs.Twitter.AppApi)
	anaconda.SetConsumerSecret(bot.Configs.Twitter.AppSecret)
	api = anaconda.NewTwitterApi(bot.Configs.Twitter.AccountApi, bot.Configs.Twitter.AccountSecret)
	_, err := api.VerifyCredentials()
	if err != nil {
		t.Errorf("Twitter Credential Failure: %v\n", err)
	}
}
