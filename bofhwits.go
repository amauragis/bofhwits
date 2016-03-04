package main

import (
	"flag"
	"log"

	"github.com/amauragis/bofhwits/bofhwitsbot"
)

// TODO: handle daemonizing
// TODO: handle log path passed in

// setup -c flag to pass a configuration file to the bot.  I suppose if you want multiple bots, you can use multiple configuration files
var configFile = flag.String("c", "config/bofhwits.yaml", "The path to the configuration file to use (default config/bofhwits.yaml)")

func main() {

	flag.Parse()

	bot := bofhwitsbot.BofhwitsBot{ConfigFilePath: *configFile}
	err := bot.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = bot.Setup()
	bot.RunBot()

}
