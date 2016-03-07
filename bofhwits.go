package main

import (
	"flag"
	"log"
	"os"

	"github.com/amauragis/bofhwits/bofhwitsbot"
)

func main() {

	// setup -c flag to pass a configuration file to the bot.
	// I suppose if you want multiple bots, you can use multiple configuration files
	configFile := flag.String("c", "config/bofhwits.yaml",
		"The path to the configuration file to use")
	logFile := flag.String("l", "",
		"The path to the log file to use. If not provided, uses stdout")

	flag.Parse()

	var logger = (*log.Logger)(nil)
	if *logFile == "" {
		logger = log.New(os.Stdout, "BOFH: ", log.Ldate|log.Ltime)
	} else {
		file, err := os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
		if err != nil {
			log.Fatal(err)
		}
		logger = log.New(file, "BOFH: ", log.Ldate|log.Ltime)
		logger.Println("-----")
		logger.Println("Log opened")
	}

	bot := bofhwitsbot.BofhwitsBot{ConfigFilePath: *configFile, Log: logger}
	err := bot.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = bot.Setup()
	bot.RunBot()

}
