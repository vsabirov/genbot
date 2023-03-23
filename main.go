package main

import (
	"fmt"

	"github.com/vsabirov/genbot/genbot"
)

type GenbotMessageHandlers struct {
	genbot.DefaultMessageHandlers
}

func (handlers GenbotMessageHandlers) OnChat(message genbot.Message) {
	chatMessage := genbot.ParseMessageBodyChat(message.Data)

	fmt.Println(string(message.Header.Username), " says ", string(chatMessage.Buffer))
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

	err = genbot.ListenAndServe(address, port, GenbotMessageHandlers{})
	if err != nil {
		fmt.Println("Genbot caught a critical error: ", err)
	}

	fmt.Println("Genbot is shutting down.")
}
