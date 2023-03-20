package genbot

import (
	"encoding/binary"
	"unicode/utf16"
)

// String operations. The game mainly uses UTF-16.

func ByteSequenceToUTF16(sequence []byte) []rune {
	size := len(sequence)

	var words []uint16
	for w := 0; w < size; w += 2 {
		words = append(words, binary.LittleEndian.Uint16(sequence[w:w+2]))
	}

	return utf16.Decode(words)
}
