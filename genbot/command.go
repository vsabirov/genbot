package genbot

import (
	"fmt"
)

// Bot command API.

type CommandRegistry map[string]func(message Message, chat MessageBodyChat)

type CommandMessageHandlers struct {
	DefaultMessageHandlers

	Prefix   string
	Commands CommandRegistry
}

func (handlers CommandMessageHandlers) OnChat(message Message) {
	chatMessage := ParseMessageBodyChat(message.Data)
	payload := string(chatMessage.Buffer)

	prefixLength := len(handlers.Prefix)
	if payload[:prefixLength] != handlers.Prefix {
		return
	}

	command := payload[prefixLength:]
	if handlers.Commands[command] == nil {
		fmt.Printf("Command '%s' not found\n", command)

		return
	}

	handlers.Commands[command](message, chatMessage)
}

func (message Message) Respond(responsePayload string, game []rune) {
	response := Message{
		Header: MessageHeader{
			CRC:   0,
			Magic: MessageDefaultMagic,

			Type: MessageChat,

			Username: message.BotInfo.Username,
		},

		Data: CreateMessageBodyChat(MessageBodyChat{
			Game:   game,
			Type:   ChatSystem,
			Buffer: []rune(responsePayload),
		}),
	}

	response.BotInfo = message.BotInfo

	BroadcastMessage(response, message.Connection)
}
