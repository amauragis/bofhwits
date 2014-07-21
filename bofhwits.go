package main

import (
    "github.com/amauragis/bofhwits/bofhwitsbot"
    "flag"
    "fmt"
)

// setup -c flag to pass a configuration file to the bot.  I suppose if you want multiple bots, you can use multiple configuration files
var config_file = flag.String("c", "config/bofhwits.yaml", "The path to the configuration file to use (default config/bofhwits.yaml)")


func main() {

	flag.Parse()

	var bot bofhwitsbot.BofhwitsBot

	bot.ConfigFilePath = *config_file
	fmt.Printf("before config:\n%v\n",bot)
	bot.LoadConfig()
	fmt.Printf("after config:\n%v\n",bot)
	bot.RunBot()

}
