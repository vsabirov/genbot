package genbot

import (
	"encoding/binary"
	"unicode/utf16"
)

// String operations. The game mainly uses UTF-16.

func byteSequenceToUTF16(sequence []byte) []rune {
	size := len(sequence)

	var words []uint16
	for w := 0; w < size; w += 2 {
		words = append(words, binary.LittleEndian.Uint16(sequence[w:w+2]))
	}

	return utf16.Decode(words)
}

func utf16ToByteSequence(text []rune) []byte {
	words := utf16.Encode(text)

	two := make([]byte, 2)

	var sequence []byte
	for w := 0; w < len(words); w++ {
		binary.LittleEndian.PutUint16(two, words[w])

		sequence = append(sequence, two...)
	}

	return sequence
}

func findNullTerminator(sequence []byte) int {
	var prev byte = 0xFF

	for pos, cur := range sequence {
		if prev == 0x00 && cur == 0x00 {
			return pos
		}

		prev = cur
	}

	return len(sequence)
}
