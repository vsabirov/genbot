package commands

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/vsabirov/genbot/genbot"
)

var currentNumber int = rand.Intn(100)

// Guess a number with binary search (tell if greater/lesser than the guessed number).
func Guess(message genbot.Message, chat genbot.MessageBodyChat, arguments []string) {
	guess, err := strconv.Atoi(arguments[1])
	if err != nil {
		// Argument is NaN, unacceptable.

		return
	}

	if guess == currentNumber {
		message.Respond(fmt.Sprintf(
			"%s IS VICTORIOUS! The number was %d.", string(message.Header.Username), currentNumber),
			chat.Game)

		currentNumber = rand.Intn(100)

		return
	}

	// Give a greater/lesser hint if the guess was unsuccessful.
	var hint string
	if currentNumber > guess {
		hint = fmt.Sprintf("The number is GREATER than %d.", guess)
	} else if currentNumber < guess {
		hint = fmt.Sprintf("The number is LESS than %d.", guess)
	}

	message.Respond(fmt.Sprintf("Incorrect. %s", hint), chat.Game)
}
