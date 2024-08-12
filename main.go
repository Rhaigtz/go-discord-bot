package main

import bot "github.com/Rhaigtz/godiscordbot/bot"

func main() {
	bot.BotToken = "DISCORD TOKEN ID"
	bot.Run() // call the run function of bot/bot.go
}
