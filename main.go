package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/vsabirov/genbot/genbot"
)

type GenbotHandlers struct {
	genbot.CommandMessageHandlers
}

func guessCommandHandler(message genbot.Message, chat genbot.MessageBodyChat, arguments []string) {
	guess, err := strconv.Atoi(arguments[1])
	if err != nil {
		message.Respond("Please, enter a valid number.", chat.Game)

		return
	}

	actual := rand.Intn(100)

	success := "CORRECT!"
	if guess != actual {
		success = "incorrect."
	}

	message.Respond(fmt.Sprintf("The number is %d. Your guess was %s", actual, success), chat.Game)
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

	handlers.Commands["guess"] = guessCommandHandler

	err = genbot.ListenAndServe(&bot, handlers)
	if err != nil {
		log.Panicln("Genbot caught a critical error: ", err)
	}

	log.Println("Genbot is shutting down.")
}
