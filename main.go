package main

import (
	"fmt"

	"github.com/vsabirov/genbot/genbot"
)

type GenbotHandlers struct {
	genbot.CommandMessageHandlers
}

func echoCommandHandler(message genbot.Message, chat genbot.MessageBodyChat) {
	message.Respond("Echo", chat.Game)
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

	fmt.Println("Genbot is starting.")

	bot := genbot.GenbotInfo{}
	bot.Address = address
	bot.Port = port
	bot.PCName = []rune("Lobby")
	bot.Username = []rune("Genbot")

	handlers := GenbotHandlers{}
	handlers.Prefix = "!"
	handlers.Commands = make(genbot.CommandRegistry)

	handlers.Commands["echo"] = echoCommandHandler

	err = genbot.ListenAndServe(&bot, handlers)
	if err != nil {
		fmt.Println("Genbot caught a critical error: ", err)
	}

	fmt.Println("Genbot is shutting down.")
}
