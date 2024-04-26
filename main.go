package main

import (
	"github.com/gtuk/discordwebhook"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// webhook URL in .env
	webhook := os.Getenv("WEBHOOK_URL")

	//set the username and avatar of the bot
	username := "🍽️ CROUSBLOUNG"
	avatar := os.Getenv("AVATAR_URL")

	// Prepare information for the CROUS meal
	crousRestaurantId := os.Getenv("CROUS_RESTAURANT_ID")

	embeds := []discordwebhook.Embed{
		getMenuEmbed(crousRestaurantId),
	}
	sendMessage(webhook, username, avatar, embeds)
}

func sendMessage(webhook string, username string, avatar string, embed []discordwebhook.Embed) {
	message := discordwebhook.Message{
		Username:  &username,
		AvatarUrl: &avatar,
		Embeds:    &embed,
	}

	err := discordwebhook.SendMessage(webhook, message)
	if err != nil {
		log.Fatal(err)
	}
}
