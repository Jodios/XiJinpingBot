package main

import (
	"log"

	"xi_jinping_bot/config"

	"github.com/Goscord/goscord"
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/gateway"
)

var client *gateway.Session
var context config.Configuration

func main() {
	log.Println("Starting bot...")
	config.InitContext(&context)

	client = goscord.New(&gateway.Options{
		Token: context.Token,
	})
	client.On("ready", func() {
		log.Println("Logged in as " + client.Me().Tag())
	})

	client.On("message", func(msg *discord.Message) {
		if msg.Content == "ping" {
			client.Channel.Send(msg.ChannelId, "Pong ! 🏓")
		}
	})
	client.Login()
	select {}
}
