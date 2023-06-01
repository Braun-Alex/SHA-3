package main

import "fmt"

const b int = 1600
const rate int = 576
const rounds int = 24
const w = 64

func sha3(data []byte) {
	countOfBits := 8 * len(data)
	bits := make([]bool, countOfBits)
	for i, byteValue := range data {
		for j := 0; j < 8; j++ {
			bits[8*i+j] = (byteValue & (1 << byte(7-j))) == 1
		}
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
	for i := 0; i < countOfBlocks-1; i++ {
		iterationBlock := bits[i : i+rate]
		for j := 0; j < b-rate; j++ {
			iterationBlock = append(iterationBlock, false)
		}
		for j := 0; j < b; j++ {
			state[j] = (state[j] && !iterationBlock[j]) || (!state[j] && iterationBlock[j])
		}
	}
}

func keccakPermutation(state [b]bool) {
	var A [5][5][w]bool
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < w; z++ {
				A[x][y][z] = state[w*(5*y+x)+z]
			}
		}
	}
}

func theta(A [5][5][w]bool) {
	var C [5][w]bool
	for x := 0; x < 5; x++ {
		for z := 0; z < w; z++ {
			C[x][z] = (A[x][0][z] && !A[x][1][z]) || (!A[x][0][z] && A[x][1][z])
			for y := 2; y < 5; y++ {
				C[x][z] = (C[x][z] && !A[x][y][z]) || (C[x][z] && A[x][y][z])
			}
		}
	}
	var D [5][w]bool
	for x := 0; x < 5; x++ {
		for z := 0; z < w; z++ {
			D[x][z] = (C[(x-1)%5][z] && !C[(x+1)%5][(z-1)%w]) ||
				(!C[(x-1)%5][z] && C[(x+1)%5][(z-1)%w])
		}
	}
	var R [5][5][w]bool
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < w; z++ {
				R[x][y][z] = (A[x][y][z] && !D[x][z]) || (A[x][y][z] && D[x][z])
			}
		}
	}
}

func main() {
	fmt.Print("AAA")
}
