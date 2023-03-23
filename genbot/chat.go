package genbot

import (
	"encoding/binary"
)

// This package provides code to work with chat message packets.

// Chat message types.
type ChatType int32

const (
	ChatNormal ChatType = iota
	ChatEmote
	ChatSystem
)

// Body of the chat message packet.
type MessageBodyChat struct {
	Game   []rune   // Source game name.
	Type   ChatType // Chat message type.
	Buffer []rune   // Buffer with the chat message contents.
}

func ParseMessageBodyChat(data []byte) MessageBodyChat {
	var cursor int = 0

	game := byteSequenceToUTF16(data[cursor : cursor+34])
	cursor += 34

	ctype := ChatType(binary.LittleEndian.Uint32(data[cursor : cursor+4]))
	cursor += 4

	buffer := byteSequenceToUTF16(data[cursor : cursor+202])
	cursor += 202

	return MessageBodyChat{
		Game:   game,
		Type:   ctype,
		Buffer: buffer,
	}
}

func createMessageBodyChat(body MessageBodyChat) []byte {
	var result []byte

	four := make([]byte, 4)
	binary.LittleEndian.PutUint32(four, uint32(body.Type))

	result = append(result, utf16ToByteSequence(body.Game)...)
	result = append(result, four...)
	result = append(result, utf16ToByteSequence(body.Buffer)...)

	return result
}
