package genbot

import (
	"encoding/binary"
)

// Message magic should be equal to this if the message is valid.
const MessageDefaultMagic uint16 = 0xF00D

// Enumeration of all known message types.
type MessageType uint32

const (
	// TODO: Research this one.
	MessageRequestLocations MessageType = iota

	// Announcements.
	MessageGameAnnounce
	MessageLobbyAnnounce

	// Player join.
	MessageRequestJoin
	MessageJoinAccept
	MessageJoinDeny

	// Player leave.
	MessageRequestGameLeave
	MessageRequestLobbyLeave

	// Player sets the acceptance flag.
	MessageSetAccept

	// Announce to other players whether someone has the current map installed or not.
	MessageMapAvailability

	// Chat.
	MessageChat

	// Game state (in lobby).
	MessageGameStart
	MessageGameTimer
	MessageGameOptions

	// TODO: Research this one.
	MessageSetActive

	// Players gather game info upon joining.
	MessageRequestGameInfo
)

type MessageHeader struct {
	crc   uint32 // CRC Sum of the entire message (including the header).
	magic uint16 // Magic number to make sure the message has been decoded successfully.

	mtype MessageType // Type of this message.

	username []rune // Name of the player that sent this message.
	flag     uint32 // TODO: Research this one.
}

type Message struct {
	header MessageHeader // Message metadata.
	data   []byte        // Message payload.
}

// Transform a decoded packet into a structurized message.
func BuildMessage(packet []byte) Message {
	var cursor int = 0

	crc := binary.LittleEndian.Uint32(packet[cursor : cursor+4])
	cursor += 4

	magic := binary.LittleEndian.Uint16(packet[cursor : cursor+2])
	cursor += 2

	mtype := binary.LittleEndian.Uint32(packet[cursor : cursor+4])
	cursor += 4

	username := ByteSequenceToUTF16(packet[cursor : cursor+26])
	cursor += 26

	flag := binary.LittleEndian.Uint32(packet[cursor : cursor+4])
	cursor += 4

	data := packet[cursor : cursor+(len(packet)-cursor)]

	return Message{
		header: MessageHeader{
			crc:   crc,
			magic: magic,

			mtype: MessageType(mtype),

			username: username,
			flag:     flag,
		},

		data: data,
	}
}

// Message type to string conversion function.
func (mtype MessageType) String() string {
	switch mtype {
	case MessageRequestLocations:
		return "Request locations"
	case MessageGameAnnounce:
		return "Game announce"
	case MessageLobbyAnnounce:
		return "Lobby announce"
	case MessageRequestJoin:
		return "Request join"
	case MessageJoinAccept:
		return "Join accept"
	case MessageJoinDeny:
		return "Join deny"
	case MessageRequestGameLeave:
		return "Request game leave"
	case MessageRequestLobbyLeave:
		return "Request lobby leave"
	case MessageSetAccept:
		return "Set accept"
	case MessageMapAvailability:
		return "Map availability"
	case MessageChat:
		return "Chat"
	case MessageGameStart:
		return "Game start"
	case MessageGameTimer:
		return "Game timer"
	case MessageGameOptions:
		return "Game options"
	case MessageSetActive:
		return "Set active"
	case MessageRequestGameInfo:
		return "Request game info"
	}

	return "<unknown message>"
}
