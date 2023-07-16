package sha3

import (
	"github.com/Braun-Alex/SHA-3/pkg/cmp"
	"github.com/Braun-Alex/SHA-3/pkg/conv"
	"github.com/Braun-Alex/SHA-3/pkg/keccak"
)

const b int = 1600
const rate int = 576
const d int = 512

func Sum512(data []byte) [d / 8]byte {
	countOfBits := 8 * len(data)
	bits := make([]bool, 0)
	for _, value := range data {
		bitsSlice := conv.ByteToBits(value)
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
	for i := 0; i < countOfBlocks; i++ {
		iterationBlock := bits[i*rate : (i+1)*rate]
		for j := 0; j < capacity; j++ {
			iterationBlock = append(iterationBlock, false)
		}
		for j := 0; j < b; j++ {
			state[j] = cmp.Xor(state[j], iterationBlock[j])
		}
		state = keccak.Permute(state)
	}
	var Z []bool
	for d > len(Z) {
		Z = append(Z, state[:rate]...)
		if d <= len(Z) {
			break
		} else {
			state = keccak.Permute(state)
		}
	}
	Z = Z[:d]
	var bytes [d / 8]byte
	copy(bytes[:], conv.BitsToBytes(Z))
	return bytes
}
