package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/adisatr7/cosette_go/internal/handlers/responses"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file!")
	}

	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("No token found in .env file!")
	}

	bot, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal("Error creating Discord session: ", err)
	}

	bot.AddHandler(func(session *discordgo.Session, message *discordgo.MessageCreate) {
		if message.Author.ID == session.State.User.ID {
			return
		}

		if message.Content == "ping" {
			session.ChannelMessageSend(message.ChannelID, "pong")
		}
	})

	bot.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = bot.Open()
	if err != nil {
		log.Fatal("Error opening Discord session: ", err)
	}
	defer bot.Close()

	fmt.Println(responses.AtMentioned())

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}