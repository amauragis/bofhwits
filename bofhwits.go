package main

import (
    "github.com/amauragis/bofhwits/bofhwitsbot"
    "fmt"
    "flag"
)

// setup -c flag to pass a configuration file to the bot.  I suppose if you want multiple bots, you can use multiple configuration files
var config_file = flag.String("c", "config/bofhwits.yaml", "The path to the configuration file to use (default config/bofhwits.yaml)")


func main() {

	flag.Parse()

	fmt.Printf("Type of config: %T\n",config_file)

	bofhwitsbot.RunBot(config_file)

}