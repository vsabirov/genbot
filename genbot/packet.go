package genbot

import (
	"encoding/binary"
)

// This code is responsible for sending packets over wire to other clients in the network.

type WirePacketHeader struct {
	crc   uint32 // CRC Checksum of this packet.
	magic uint16 // Magic value for this packet.
}

type WirePacket struct {
	header WirePacketHeader // Header of this packet.
	data   [1024]byte       // Data of this packet.
	length int32            // Length of data of this packet.
	addr   uint32           // Receiver address.
	port   uint16           // Receiver port.
}

func messageToWirePacket(message Message, addr uint32, port uint16) WirePacket {
	four := make([]byte, 4)
	binary.LittleEndian.PutUint32(four, uint32(message.Header.Type))

	var data []byte
	data = append(data, four...)

	data = append(data, utf16ToByteSequence(message.Header.Username)...)

	binary.LittleEndian.PutUint32(four, uint32(message.Header.flag))
	data = append(data, four...)

	data = append(data, message.Data...)

	size := len(data)
	for i := 0; i < 1024-size; i++ {
		data = append(data, 0x0)
	}

	return WirePacket{
		header: WirePacketHeader{
			crc:   message.Header.CRC,
			magic: message.Header.Magic,
		},

		data:   [1024]byte(data),
		length: int32(len(data)),

		addr: addr,
		port: port,
	}
}

func wirePacketToBlob(packet WirePacket) []byte {
	var result []byte

	four := make([]byte, 4)
	two := make([]byte, 2)

	binary.LittleEndian.PutUint32(four, packet.header.crc)
	result = append(result, four...)

	binary.LittleEndian.PutUint16(two, packet.header.magic)
	result = append(result, two...)

	result = append(result, packet.data[:]...)

	binary.LittleEndian.PutUint32(four, uint32(packet.length))
	result = append(result, four...)

	binary.LittleEndian.PutUint32(four, packet.addr)
	result = append(result, four...)

	binary.LittleEndian.PutUint16(two, packet.port)
	result = append(result, two...)

	size := len(result)

	for i := 0; i < 1044-size; i++ {
		result = append(result, 0x0)
	}

	return result
}
