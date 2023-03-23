package genbot

import "strings"

// Bot command API.

type CommandRegistry map[string]func(message Message, chat MessageBodyChat, arguments []string)

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

	arguments := strings.Split(payload, " ")

	command := arguments[0][prefixLength:]
	if handlers.Commands[command] == nil {
		return
	}

	go handlers.Commands[command](message, chatMessage, arguments)
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
			Type:   ChatEmote,
			Buffer: []rune(" > " + responsePayload),
		}),
	}

	response.BotInfo = message.BotInfo

	BroadcastMessage(response, message.Connection)
}
