package genbot

import (
	"fmt"
	"net"
)

// Map which helps to search message handler by its type.
type MessageHandlerMap map[MessageType]func(message Message)

// Default key for deobfuscation.
const PacketKey = 0xFADE

// Starts the bot, opens the socket and listens for messages.
func ListenAndServe(address string, port uint16, handlers MessageHandlers) error {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{Port: int(port), IP: net.ParseIP(address)})
	if err != nil {
		return err
	}

	return serve(listener, createHandlerMap(handlers))
}

// Transform handlers interface to message resolver map.
func createHandlerMap(handlers MessageHandlers) MessageHandlerMap {
	result := make(MessageHandlerMap)

	result[MessageRequestLocations] = handlers.OnRequestLocations

	result[MessageGameAnnounce] = handlers.OnGameAnnounce
	result[MessageLobbyAnnounce] = handlers.OnLobbyAnnounce

	result[MessageRequestJoin] = handlers.OnRequestJoin
	result[MessageJoinAccept] = handlers.OnJoinAccept
	result[MessageJoinDeny] = handlers.OnJoinDeny

	result[MessageRequestGameLeave] = handlers.OnRequestGameLeave
	result[MessageRequestLobbyLeave] = handlers.OnRequestLobbyLeave

	result[MessageSetAccept] = handlers.OnSetAccept

	result[MessageMapAvailability] = handlers.OnMapAvailability

	result[MessageChat] = handlers.OnChat

	result[MessageGameStart] = handlers.OnGameStart
	result[MessageGameTimer] = handlers.OnGameTimer
	result[MessageGameOptions] = handlers.OnGameOptions

	result[MessageSetActive] = handlers.OnSetActive

	result[MessageRequestGameInfo] = handlers.OnRequestGameInfo

	return result
}

// Receive incoming messages and process each one in a different goroutine.
func serve(connection *net.UDPConn, handlerMap MessageHandlerMap) error {
	defer connection.Close()

	packet := make([]byte, 1024)
	for {
		len, addr, err := connection.ReadFrom(packet)
		if err != nil {
			fmt.Println("Genbot client sent bad packet, dropping connection. ", err)
		}

		go handle(packet[:len], addr, connection, handlerMap)
	}
}

// Read the message & respond if needed.
func handle(packet []byte, sender net.Addr, connection *net.UDPConn, handlerMap MessageHandlerMap) {
	sanitize(&packet)

	message := buildMessage(packet)
	if message.Header.Magic != MessageDefaultMagic {
		// Ignore invalid message magics.
		return
	}

	crc := computeCRC(packet[4:])
	if crc != message.Header.CRC {
		// Ignore invalid message checksums.
		return
	}

	if !(message.Header.Type >= MessageRequestLocations && message.Header.Type <= MessageRequestGameInfo) {
		// Ignore invalid message types.

		return
	}

	handlerMap[message.Header.Type](message)
}

// Prepare packet for structurization.
func sanitize(packet *[]byte) {
	deobfuscate(packet, PacketKey)
}
