package main

import bot "github.com/rhaigtz/go-discord-bot/bot"

func main() {
	bot.BotToken = "DISCORD TOKEN ID"
	bot.Run() // call the run function of bot/bot.go
}
