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
	data   []byte           // Data of this packet.
}

func padBytes(sequence []byte, desiredSize int) []byte {
	size := len(sequence)
	if size >= desiredSize {
		return sequence
	}

	for p := 0; p < desiredSize-size; p++ {
		sequence = append(sequence, 0x0)
	}

	return sequence
}

func messageToWirePacket(message Message) WirePacket {
	four := make([]byte, 4)
	binary.LittleEndian.PutUint32(four, uint32(message.Header.Type))

	var data []byte
	data = append(data, four...)

	data = append(data, padBytes(utf16ToByteSequence(message.Header.Username), 26)...)

	binary.LittleEndian.PutUint32(four, uint32(message.Header.flag))
	data = append(data, four...)

	data = append(data, message.Data...)

	return WirePacket{
		header: WirePacketHeader{
			crc:   message.Header.CRC,
			magic: message.Header.Magic,
		},

		data: []byte(data),
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

	return result
}

func fillCRCForWirePacket(packet WirePacket) WirePacket {
	blob := wirePacketToBlob(packet)

	packet.header.crc = computeCRC(blob[4:])

	return packet
}
