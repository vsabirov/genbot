package genbot

import (
	"fmt"
	"net"
)

// Default key for deobfuscation.
const PacketKey = 0xFADE

// Starts the bot, opens the socket and listens for messages.
func ListenAndServe(address string, port uint16) error {
	listener, err := net.ListenUDP("udp", &net.UDPAddr{Port: int(port), IP: net.ParseIP(address)})
	if err != nil {
		return err
	}

	return serve(listener)
}

// Receive incoming messages and process each one in a different goroutine.
func serve(connection *net.UDPConn) error {
	defer connection.Close()

	packet := make([]byte, 1024)
	for {
		len, addr, err := connection.ReadFrom(packet)
		if err != nil {
			fmt.Println("Genbot client sent bad packet, dropping connection. ", err)
		}

		go handle(packet[:len], addr)
	}
}

// Read the message & respond if needed.
func handle(packet []byte, sender net.Addr) {
	sanitize(&packet)

	message := buildMessage(packet)
	if message.header.magic != MessageDefaultMagic {
		// Ignore invalid messages.
		return
	}

	if message.header.mtype == MessageChat {
		body := parseMessageBodyChat(message.data)

		fmt.Println(string(message.header.username), " says ", string(body.buffer))
	}
}

// Prepare packet for structurization.
func sanitize(packet *[]byte) {
	deobfuscate(packet, PacketKey)
}
