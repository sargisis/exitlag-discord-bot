package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var Session *discordgo.Session

func Start(token string) {
	var err error
	Session, err = discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	Session.AddHandler(InteractionHandler)
	Session.AddHandler(MessageCreateHandler)

	// Identify intents
	Session.Identify.Intents = discordgo.IntentsGuilds |
		discordgo.IntentsGuildMessages |
		discordgo.IntentsMessageContent

	err = Session.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	log.Println("Bot is running...")
}
