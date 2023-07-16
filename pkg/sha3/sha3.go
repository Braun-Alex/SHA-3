package sha3

import (
	"github.com/Braun-Alex/SHA-3/pkg/cmp"
	"github.com/Braun-Alex/SHA-3/pkg/conv"
	"github.com/Braun-Alex/SHA-3/pkg/keccak"
)

const B int = 1600
const Rate int = 576
const D int = 512

func Sum512(data []byte) [D / 8]byte {
	countOfBits := 8 * len(data)
	bits := conv.BytesToBits(data)
	lastBlockSize := countOfBits % Rate
	difference := Rate - 2
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
	capacity := B - Rate
	countOfBits = len(bits)
	countOfBlocks := countOfBits / Rate
	var state [B]bool
	for i := 0; i < countOfBlocks; i++ {
		iterationBlock := bits[i*Rate : (i+1)*Rate]
		for j := 0; j < capacity; j++ {
			iterationBlock = append(iterationBlock, false)
		}
		for j := 0; j < B; j++ {
			state[j] = cmp.Xor(state[j], iterationBlock[j])
		}
		state = keccak.Permute(state)
	}
	var Z []bool
	for D > len(Z) {
		Z = append(Z, state[:Rate]...)
		if D <= len(Z) {
			break
		} else {
			state = keccak.Permute(state)
		}
	}
	Z = Z[:D]
	var bytes [D / 8]byte
	copy(bytes[:], conv.BitsToBytes(Z))
	return bytes
}
