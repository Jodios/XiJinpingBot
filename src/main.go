package main

import (
	"log"
	"os"

	"xi_jinping_bot/config"
	my_util "xi_jinping_bot/util"

	"github.com/Goscord/goscord"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/gateway"
)

var client *gateway.Session
var context config.Configuration

func main() {
	config.InitContext(&context)
	client = goscord.New(&gateway.Options{
		Token: os.Getenv("DISCORD_TOKEN"),
	})

	client.On("ready", on_ready)
	client.Login()
	select {}
}

func on_ready() {
	config.PrintBanner()
	log.Println("Logged in as " + client.Me().Tag())
	client.On("message", on_message)
}

func on_message(msg *discord.Message) {
	if my_util.Contains(&context.No_No_Words, msg.Content) {
		client.Channel.Send(msg.ChannelId, ":angry:")
	}
}
