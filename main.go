package main

import (
	"fmt"

	"github.com/Lucif3r-in/discord-bot/bot"
	"github.com/Lucif3r-in/discord-bot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println((err.Error()))
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
