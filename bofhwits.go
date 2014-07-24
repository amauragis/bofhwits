package main

import (
	"flag"
	"github.com/amauragis/bofhwits/bofhwitsbot"
)

// setup -c flag to pass a configuration file to the bot.  I suppose if you want multiple bots, you can use multiple configuration files
var config_file = flag.String("c", "config/bofhwits.yaml", "The path to the configuration file to use (default config/bofhwits.yaml)")

func main() {

	flag.Parse()

	bot := bofhwitsbot.BofhwitsBot{ConfigFilePath: *config_file}
	bot.LoadConfig()
	bot.RunBot()

}
