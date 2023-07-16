package conv

func ByteToBits(byteValue byte) []bool {
	bits := make([]bool, 0)
	for byteValue > 0 {
		remainder := byteValue % 2
		bits = append([]bool{remainder == 1}, bits...)
		byteValue = byteValue / 2
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
