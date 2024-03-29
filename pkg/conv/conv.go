package conv

func ByteToBits(byteValue byte) []bool {
	bits := make([]bool, 0)
	for byteValue > 0 {
		remainder := byteValue % 2
		bits = append([]bool{remainder == 1}, bits...)
		byteValue = byteValue / 2
	}
	complementaryBits := make([]bool, 0)
	for i := 0; i < 8-len(bits); i++ {
		complementaryBits = append(complementaryBits, false)
	}
	bits = append(complementaryBits, bits...)
	return bits
}

func BytesToBits(bytes []byte) []bool {
	bits := make([]bool, 0)
	for _, byteValue := range bytes {
		bitsSlice := ByteToBits(byteValue)
		bits = append(bits, bitsSlice...)
	}
	return bits
}

func BitsToBytes(bits []bool) []byte {
	byteCount := (len(bits) + 7) / 8
	bytes := make([]byte, byteCount)
	for i := 0; i < len(bits); i += 8 {
		var byteValue byte
		for j := 0; j < 8; j++ {
			bitIndex := i + j
			if bitIndex < len(bits) && bits[bitIndex] {
				byteValue |= 1 << (7 - j)
			}
		}
		bytes[i/8] = byteValue
	}
	return bytes
}
