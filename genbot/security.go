package genbot

import (
	"encoding/binary"
)

/*
	Packet obfuscation/deobfuscation.

	Packet data can be obfuscated in sequential blocks, each 4 bytes in size.
	This obfuscation is a simple XOR operation that is done for each individual block of data in the packet.
	The initial value for the XOR operation is set to a specific key. In a generic scenario,
	the value of the key will be equal to 0xFADE.

	Each block (4 bytes) XORed, the value for the XOR operation (mixMagic) is incremented by 0x321.
*/

func Obfuscate(deobfuscated *[]byte, key uint32) {
	size := len(*deobfuscated)

	var amountOfBlocks int = size / 4
	var currentBlock uint32 = binary.LittleEndian.Uint32((*deobfuscated)[:4])
	var mixMagic uint32 = key

	currentBlock ^= uint32(mixMagic)
	binary.BigEndian.PutUint32((*deobfuscated)[:4], currentBlock)

	mixMagic += 0x321

	for b := 1; b < amountOfBlocks; b++ {
		currentBlock = binary.BigEndian.Uint32((*deobfuscated)[b*4 : b*4+4])

		currentBlock ^= uint32(mixMagic)
		binary.BigEndian.PutUint32((*deobfuscated)[b*4:b*4+4], currentBlock)

		mixMagic += 0x321
	}
}

func Deobfuscate(obfuscated *[]byte, key uint32) {
	size := len(*obfuscated)

	var amountOfBlocks int = size / 4
	var currentBlock uint32 = binary.BigEndian.Uint32((*obfuscated)[:4])
	var mixMagic uint32 = key

	currentBlock ^= uint32(mixMagic)
	binary.LittleEndian.PutUint32((*obfuscated)[:4], currentBlock)

	mixMagic += 0x321

	for b := 1; b < amountOfBlocks; b++ {
		currentBlock = binary.BigEndian.Uint32((*obfuscated)[b*4 : b*4+4])

		currentBlock ^= uint32(mixMagic)
		binary.LittleEndian.PutUint32((*obfuscated)[b*4:b*4+4], currentBlock)

		mixMagic += 0x321
	}
}
