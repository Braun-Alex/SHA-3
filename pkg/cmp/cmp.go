package cmp

func Xor(x, y bool) bool {
	return (x && !y) || (!x && y)
}

func Mod(a, p int) int {
	remainder := a % p
	if remainder < 0 {
		return remainder + p
	} else {
		return remainder
	}
}
