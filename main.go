package main

import (
	"fmt"
	"math"
)

const b int = 1600
const rate int = 576
const rounds int = 24
const l = 6
const w = 64

func xor(x, y bool) bool {
	return (x && !y) || (!x && y)
}

func sha3(data []byte) []byte {
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
			state[j] = xor(state[j], iterationBlock[j])
		}
	}
}

func keccakPermutation(state [b]bool) [b]bool {
	var A [5][5][w]bool
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < w; z++ {
				A[x][y][z] = state[w*(5*y+x)+z]
			}
		}
	}
}

func theta(A [5][5][w]bool) [5][5][w]bool {
	var C [5][w]bool
	for x := 0; x < 5; x++ {
		for z := 0; z < w; z++ {
			C[x][z] = xor(A[x][0][z], A[x][1][z])
			for y := 2; y < 5; y++ {
				C[x][z] = xor(C[x][z], A[x][y][z])
			}
		}
	}
	var D [5][w]bool
	for x := 0; x < 5; x++ {
		for z := 0; z < w; z++ {
			D[x][z] = xor(C[(x-1)%5][z], C[(x+1)%5][(z-1)%w])
		}
	}
	var R [5][5][w]bool
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < w; z++ {
				R[x][y][z] = xor(A[x][y][z], D[x][z])
			}
		}
	}
	return R
}

func rho(A [5][5][w]bool) [5][5][w]bool {
	var R [5][5][w]bool
	for z := 0; z < w; z++ {
		R[0][0][z] = A[0][0][z]
	}
	x, y := 1, 0
	for t := 0; t < 24; t++ {
		for z := 0; z < w; z++ {
			R[x][y][z] = A[x][y][(z-(t+1)*(t+2)/2)%w]
		}
		x, y = y, (2*x+3*y)%5
	}
	return R
}

func pi(A [5][5][w]bool) [5][5][w]bool {
	var R [5][5][w]bool
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < w; z++ {
				R[x][y][z] = A[(x+3*y)%5][x][z]
			}
		}
	}
	return R
}

func chi(A [5][5][w]bool) [5][5][w]bool {
	var R [5][5][w]bool
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < w; z++ {
				R[x][y][z] = xor(A[x][y][z], !A[(x+1)%5][y][z] && A[(x+2)%5][y][z])
			}
		}
	}
	return R
}

func rc(t int) bool {
	if t%255 == 0 {
		return true
	}
	var R []bool
	R = append(R, true)
	for i := 0; i < 7; i++ {
		R = append(R, false)
	}
	for i := 1; i <= t%255; i++ {
		R = append([]bool{false}, R...)
		R[0] = xor(R[0], R[8])
		R[4] = xor(R[4], R[8])
		R[5] = xor(R[5], R[8])
		R[6] = xor(R[6], R[8])
		R = R[:8]
	}
	return R[0]
}

func iota(A [5][5][w]bool, round int) [5][5][w]bool {
	var R [5][5][w]bool
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < w; z++ {
				R[x][y][z] = A[x][y][z]
			}
		}
	}
	var RC [w]bool
	for j := 0; j <= l; j++ {
		RC[int(math.Pow(2, float64(j)))-1] = rc(j + 7*round)
	}
	for z := 0; z < w; z++ {
		R[0][0][z] = xor(R[0][0][z], RC[z])
	}
	return R
}

func main() {
	fmt.Print("AAA")
}
