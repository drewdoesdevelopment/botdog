package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	loginBot()
}

func loginBot() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file in main.go!")
	}

	botApiKey := os.Getenv("BOT_API_SECRET")
    fmt.Println(botApiKey)

	discord, discLibErr := discordgo.New("Bot " + botApiKey)

    if discLibErr != nil {
        log.Fatal(discLibErr)
    }

    connectionErr := discord.Open()
    if connectionErr != nil {
        log.Fatal(connectionErr)
    }
    defer discord.Close()

    fmt.Println("Botdog running...")
    ch := make(chan os.Signal, 1)
    signal.Notify(ch, os.Interrupt)
    <- ch
}
