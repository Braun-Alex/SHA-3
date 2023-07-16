package keccak

import "github.com/Braun-Alex/SHA-3/pkg/cmp"

const b int = 1600
const rounds int = 24
const l int = 6
const w int = 64

func Permute(state [b]bool) [b]bool {
	var A [5][5][w]bool
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < w; z++ {
				A[x][y][z] = state[w*(5*y+x)+z]
			}
		}
	}
	for round := 0; round < rounds; round++ {
		A = theta(A)
		A = rho(A)
		A = pi(A)
		A = chi(A)
		A = iota(A, round)
	}
	bitIndex := 0
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			for z := 0; z < w; z++ {
				state[bitIndex] = A[x][y][z]
				bitIndex++
			}
		}
	}
	return state
}

func theta(A [5][5][w]bool) [5][5][w]bool {
	var C [5][w]bool
	for x := 0; x < 5; x++ {
		for z := 0; z < w; z++ {
			C[x][z] = cmp.Xor(A[x][0][z], A[x][1][z])
			for y := 2; y < 5; y++ {
				C[x][z] = cmp.Xor(C[x][z], A[x][y][z])
			}
		}
	}
	var D [5][w]bool
	for x := 0; x < 5; x++ {
		for z := 0; z < w; z++ {
			D[x][z] = cmp.Xor(C[cmp.Mod(x-1, 5)][z], C[cmp.Mod(x+1, 5)][cmp.Mod(z-1, w)])
		}
	}
	var R [5][5][w]bool
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < w; z++ {
				R[x][y][z] = cmp.Xor(A[x][y][z], D[x][z])
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
			R[x][y][z] = A[x][y][cmp.Mod(z-(t+1)*(t+2)/2, w)]
		}
		x, y = y, cmp.Mod(2*x+3*y, 5)
	}
	return R
}

func pi(A [5][5][w]bool) [5][5][w]bool {
	var R [5][5][w]bool
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < w; z++ {
				R[x][y][z] = A[cmp.Mod(x+3*y, 5)][x][z]
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
				R[x][y][z] = cmp.Xor(A[x][y][z],
					!A[cmp.Mod(x+1, 5)][y][z] && A[cmp.Mod(x+2, 5)][y][z])
			}
		}
	}
	return R
}

func rc(t int) bool {
	if cmp.Mod(t, 255) == 0 {
		return true
	}
	var R []bool
	R = append(R, true)
	for i := 0; i < 7; i++ {
		R = append(R, false)
	}
	for i := 1; i <= cmp.Mod(t, 255); i++ {
		R = append([]bool{false}, R...)
		R[0] = cmp.Xor(R[0], R[8])
		R[4] = cmp.Xor(R[4], R[8])
		R[5] = cmp.Xor(R[5], R[8])
		R[6] = cmp.Xor(R[6], R[8])
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
		RC[1<<j-1] = rc(j + 7*round)
	}
	for z := 0; z < w; z++ {
		R[0][0][z] = cmp.Xor(R[0][0][z], RC[z])
	}
	return R
}
