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
	game   []rune   // Source game name.
	ctype  ChatType // Chat message type.
	buffer []rune   // Buffer with the chat message contents.
}

func ParseMessageBodyChat(data []byte) MessageBodyChat {
	var cursor int = 0

	game := ByteSequenceToUTF16(data[cursor : cursor+34])
	cursor += 34

	ctype := ChatType(binary.LittleEndian.Uint32(data[cursor : cursor+4]))
	cursor += 4

	buffer := ByteSequenceToUTF16(data[cursor : cursor+202])
	cursor += 202

	return MessageBodyChat{
		game:   game,
		ctype:  ctype,
		buffer: buffer,
	}
}
