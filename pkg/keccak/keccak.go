package keccak

import "github.com/Braun-Alex/SHA-3/pkg/cmp"

const B int = 1600
const Rounds int = 24
const L int = 6
const W int = 64

func Permute(state [B]bool) [B]bool {
	A := FromState(state)
	for round := 0; round < Rounds; round++ {
		A = Theta(A)
		A = Rho(A)
		A = Pi(A)
		A = Chi(A)
		A = Iota(A, round)
	}
	state = ToState(A)
	return state
}

func FromState(state [B]bool) [5][5][W]bool {
	var A [5][5][W]bool
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < W; z++ {
				A[x][y][z] = state[W*(5*y+x)+z]
			}
		}
	}
	return A
}

func ToState(A [5][5][W]bool) [B]bool {
	bitIndex := 0
	var state [B]bool
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			for z := 0; z < W; z++ {
				state[bitIndex] = A[x][y][z]
				bitIndex++
			}
		}
	}
	return state
}

func Theta(A [5][5][W]bool) [5][5][W]bool {
	var C [5][W]bool
	for x := 0; x < 5; x++ {
		for z := 0; z < W; z++ {
			C[x][z] = A[x][0][z]
			for y := 1; y < 5; y++ {
				C[x][z] = cmp.Xor(C[x][z], A[x][y][z])
			}
		}
	}
	var D [5][W]bool
	for x := 0; x < 5; x++ {
		for z := 0; z < W; z++ {
			D[x][z] = cmp.Xor(C[cmp.Mod(x-1, 5)][z], C[cmp.Mod(x+1, 5)][cmp.Mod(z-1, W)])
		}
	}
	var R [5][5][W]bool
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < W; z++ {
				R[x][y][z] = cmp.Xor(A[x][y][z], D[x][z])
			}
		}
	}
	return R
}

func Rho(A [5][5][W]bool) [5][5][W]bool {
	var R [5][5][W]bool
	for z := 0; z < W; z++ {
		R[0][0][z] = A[0][0][z]
	}
	x, y := 1, 0
	for t := 0; t < 24; t++ {
		for z := 0; z < W; z++ {
			R[x][y][z] = A[x][y][cmp.Mod(z-(t+1)*(t+2)/2, W)]
		}
		x, y = y, cmp.Mod(2*x+3*y, 5)
	}
	return R
}

func Pi(A [5][5][W]bool) [5][5][W]bool {
	var R [5][5][W]bool
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < W; z++ {
				R[x][y][z] = A[cmp.Mod(x+3*y, 5)][x][z]
			}
		}
	}
	return R
}

func Chi(A [5][5][W]bool) [5][5][W]bool {
	var R [5][5][W]bool
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < W; z++ {
				R[x][y][z] = cmp.Xor(A[x][y][z],
					!A[cmp.Mod(x+1, 5)][y][z] && A[cmp.Mod(x+2, 5)][y][z])
			}
		}
	}
	return R
}

func Rc(t int) bool {
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

func Iota(A [5][5][W]bool, round int) [5][5][W]bool {
	var R [5][5][W]bool
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			for z := 0; z < W; z++ {
				R[x][y][z] = A[x][y][z]
			}
		}
	}
	var RC [W]bool
	for j := 0; j <= L; j++ {
		RC[1<<j-1] = Rc(j + 7*round)
	}
	for z := 0; z < W; z++ {
		R[0][0][z] = cmp.Xor(R[0][0][z], RC[z])
	}
	return R
}
