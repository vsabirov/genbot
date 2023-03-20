package genbot

// Cyclic redundancy check calculations.

func addCRC(crc uint32, adder byte) uint32 {
	return uint32(adder) + (crc >> 31) + 2*crc
}

func computeCRC(data []byte) uint32 {
	var result uint32 = 0

	for _, adder := range data {
		result = addCRC(result, adder)
	}

	return result
}
