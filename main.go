package main

import (
	"fmt"
	"log"

	"github.com/vsabirov/genbot/commands"
	"github.com/vsabirov/genbot/genbot"
)

type GenbotHandlers struct {
	genbot.CommandMessageHandlers
}

func main() {
	var (
		address string
		port    uint16
	)

	fmt.Println("Enter address and port to listen to (format: 'address port'): ")

	_, err := fmt.Scanln(&address, &port)
	if err != nil {
		fmt.Println("Please, enter valid address and port. ", err)

		return
	}

	log.Println("Genbot is starting.")

	bot := genbot.GenbotInfo{}
	bot.Address = address
	bot.Port = port
	bot.PCName = []rune("Lobby")
	bot.Username = []rune("Genbot")

	handlers := GenbotHandlers{}
	handlers.Prefix = "!"
	handlers.Commands = make(genbot.CommandRegistry)

	handlers.Commands["guess"] = commands.Guess
	handlers.Commands["announce"] = commands.Announce

	err = genbot.ListenAndServe(&bot, handlers)
	if err != nil {
		log.Panicln("Genbot caught a critical error: ", err)
	}

	log.Println("Genbot is shutting down.")
}
