package main

import (
	"fmt"

	"github.com/vsabirov/genbot/genbot"
)

type GenbotHandlers struct {
	genbot.DefaultMessageHandlers
}

func (handlers GenbotHandlers) OnChat(message genbot.Message) {
	chatMessage := genbot.ParseMessageBodyChat(message.Data)

	response := genbot.Message{
		Header: genbot.MessageHeader{
			CRC:   0,
			Magic: genbot.MessageDefaultMagic,

			Type: genbot.MessageChat,

			Username: message.BotInfo.Username,
		},

		Data: genbot.CreateMessageBodyChat(genbot.MessageBodyChat{
			Game:   chatMessage.Game,
			Type:   genbot.ChatNormal,
			Buffer: chatMessage.Buffer,
		}),
	}

	genbot.SendMessage(response, message.Connection, message.Sender)

	fmt.Println(string(message.Header.Username), " says ", string(chatMessage.Buffer), " from ", string(chatMessage.Game))
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

	err = genbot.ListenAndServe(bot, GenbotHandlers{})
	if err != nil {
		fmt.Println("Genbot caught a critical error: ", err)
	}

	fmt.Println("Genbot is shutting down.")
}
