package genbot

import (
	"fmt"
	"net"
)

// General info about a bot.
type GenbotInfo struct {
	Address string
	Port    uint16

	Username []rune
	PCName   []rune

	Players map[string]net.Addr // Internal.
}

// Map which helps to search message handler by its type.
type MessageHandlerMap map[MessageType]func(message Message)

// Default key for deobfuscation.
const PacketKey = 0xFADE

// Starts the bot, opens the socket and listens for messages.
func ListenAndServe(bot *GenbotInfo, handlers MessageHandlers) error {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{Port: int(bot.Port), IP: net.ParseIP(bot.Address)})
	if err != nil {
		return err
	}

	fmt.Println("Connection opened, announcing self as ", string(bot.Username))
	bot.Players = make(map[string]net.Addr)

	return serve(listener, createHandlerMap(handlers), bot)
}

// Prepare message packet for transport.
func prepare(message Message) []byte {
	wire := messageToWirePacket(message)
	wire = fillCRCForWirePacket(wire)

	blob := wirePacketToBlob(wire)

	obfuscate(&blob, PacketKey)

	return blob
}

// Sends a message down a connection to a receiver.
func SendMessage(message Message, connection *net.UDPConn, receiver net.Addr) {
	blob := prepare(message)

	connection.WriteTo(blob, receiver)
}

// Broadcasts a message down a connection to everyone.
func BroadcastMessage(message Message, connection *net.UDPConn) {
	blob := prepare(message)

	for username, address := range message.BotInfo.Players {
		fmt.Println(username, address)

		connection.WriteTo(blob, address)
	}
}

// Announce to a specific player that genbot is connected to the lobby.
func announceSelf(connection *net.UDPConn, receiver net.Addr, bot GenbotInfo) {
	heartbeat := Message{
		Header: MessageHeader{
			CRC:   0,
			Magic: MessageDefaultMagic,

			Type: MessageLobbyAnnounce,

			Username: bot.Username,
		},

		Data: CreateMessageBodyLobbyAnnounce(MessageBodyLobbyAnnounce{
			PCName: bot.PCName,
		}),
	}

	SendMessage(heartbeat, connection, receiver)
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
func serve(connection *net.UDPConn, handlerMap MessageHandlerMap, bot *GenbotInfo) error {
	defer connection.Close()

	packet := make([]byte, 1030)
	for {
		len, addr, err := connection.ReadFrom(packet)
		if err != nil {
			fmt.Println("Genbot client sent bad packet, dropping connection. ", err)
		}

		go handle(packet[:len], addr, connection, handlerMap, bot)
	}
}

// Read the message & respond if needed.
func handle(packet []byte, sender net.Addr, connection *net.UDPConn, handlerMap MessageHandlerMap, bot *GenbotInfo) {
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

	message.Sender = sender
	message.Connection = connection
	message.BotInfo = bot

	go handlerMap[message.Header.Type](message)
}

// Prepare packet for structurization.
func sanitize(packet *[]byte) {
	deobfuscate(packet, PacketKey)
}
