package bofhwitsbot

import (
	"github.com/ChimeraCoder/anaconda"
	"os"
	"testing"
)

// tests don't work.  sorry

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

func Test_FormatTextOneLine(t *testing.T) {
	input1 :=
		`20:23 <Malkar> though it is really quiet in here
	20:23 <Malkar> what the hell are you people doing on a friday night
	20:23 <Malkar> you /should/ be keeping me entertained while I work`

	// XXX: figure out how to access private methods in Test
	output1 := bot.formatTextOneLine(input1)
	exp_output1 := `20:23 <Malkar> though it is really quiet in here | 20:23 <Malkar> what the hell are you people doing on a friday night | 20:23 <Malkar> you /should/ be keeping me entertained while I work`

	if output1 != exp_output1 {
		t.Error("One-line format error: case 1")
	}
}

func Test_FormatTextMultiLine(t *testing.T) {
	input1 :=
		`20:23 <Malkar> though it is really quiet in here
	20:23 <Malkar> what the hell are you people doing on a friday night
	20:23 <Malkar> you /should/ be keeping me entertained while I work`

}
