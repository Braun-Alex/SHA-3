package sha3

const b int = 1600
const rate int = 576
const rounds int = 24
const l int = 6
const w int = 64
const d int = 512

func Sum512(data []byte) [d / 8]byte {
	countOfBits := 8 * len(data)
	bits := make([]bool, 0)
	for _, value := range data {
		bitsSlice := ByteToBits(value)
		for i := 0; i < 8-len(bitsSlice); i++ {
			bits = append(bits, false)
		}
		bits = append(bits, bitsSlice...)
	}
	lastBlockSize := countOfBits % rate
	difference := rate - 2
	if lastBlockSize == 0 {
		bits = append(bits, true)
		for i := 0; i < difference; i++ {
			bits = append(bits, false)
		}
		bits = append(bits, true)
	} else {
		if lastBlockSize <= difference {
			bits = append(bits, true)
			countOfFalseValues := difference - lastBlockSize
			for i := 0; i < countOfFalseValues; i++ {
				bits = append(bits, false)
			}
			bits = append(bits, true)
		} else {
			bits = append(bits, true)
			for i := 0; i < difference; i++ {
				bits = append(bits, false)
			}
			bits = append(bits, false)
			bits = append(bits, true)
		}
	}
	capacity := b - rate
	countOfBits = len(bits)
	countOfBlocks := countOfBits / rate
	var state [b]bool
	for i := range state {
		state[i] = false
	}
	for i := 0; i < countOfBlocks+1; i++ {
		iterationBlock := bits[i*rate : (i+1)*rate]
		for j := 0; j < capacity; j++ {
			iterationBlock = append(iterationBlock, false)
		}
		for j := 0; j < b; j++ {
			state[j] = xor(state[j], iterationBlock[j])
		}
		state = keccakPermutation(state)
	}
	var Z []bool
	for d > len(Z) {
		Z = append(Z, state[:rate]...)
		if d <= len(Z) {
			break
		} else {
			state = keccakPermutation(state)
		}
	}
	Z = Z[:d]
	var bytes [d / 8]byte
	copy(bytes[:], bitsToBytes(Z))
	return bytes
}
