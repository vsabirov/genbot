package genbot

// This package provides code to work with lobby packets.

// Body of the lobby announce message.
type MessageBodyLobbyAnnounce struct {
	PCName []rune // User PC name.
}

func ParseMessageBodyLobbyAnnounce(data []byte) MessageBodyLobbyAnnounce {
	var cursor int = 0

	pcname := byteSequenceToUTF16(data[cursor : cursor+16])
	cursor += 16

	return MessageBodyLobbyAnnounce{
		PCName: pcname,
	}
}

func CreateMessageBodyLobbyAnnounce(body MessageBodyLobbyAnnounce) []byte {
	var result []byte

	result = append(result, padBytes(utf16ToByteSequence(body.PCName), 16)...)

	return result
}
