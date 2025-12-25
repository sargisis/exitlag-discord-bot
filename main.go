package main

import (
	"log"
	"net/http"
	"os"

	"exitlag-bot/bot"
	"exitlag-bot/config"
)

func main() {
	cfg := config.LoadConfig()

	// Start Discord Bot
	bot.Start(cfg.DiscordToken)

	// Health check for Render
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ExitLag Bot is running"))
	})

	// Get PORT from environment (for Discloud/Heroku)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting HTTP server on :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("HTTP server failed: %v", err)
	}
}
