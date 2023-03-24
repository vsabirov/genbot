package commands

import (
	"strings"

	"github.com/vsabirov/genbot/genbot"
)

// Can be called from any game chat to announce a message in the main lobby.
func Announce(message genbot.Message, chat genbot.MessageBodyChat, arguments []string) {
	if len(arguments) < 2 {
		return
	}

	if len(strings.TrimSpace(string(chat.Game))) == 0 {
		// Can't announce in the main lobby.

		return
	}

	announcementPayload := " +++ FROM " + string(message.Header.Username) + " +++ " + strings.Join(arguments, " ")

	// Respond to the command in the main lobby.
	message.Respond(announcementPayload, []rune(""))
}
