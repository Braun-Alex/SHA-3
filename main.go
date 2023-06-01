package main

import "fmt"

const blockSize int = 576

func sha3(data []byte) {
	countOfBits := 8 * len(data)
	bits := make([]bool, countOfBits)
	for i, byteValue := range data {
		for j := 0; j < 8; j++ {
			bits[8*i+j] = (byteValue & (1 << byte(7-j))) == 1
		}
	}
	lastBlockSize := countOfBits % blockSize
	difference := blockSize - 2
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
}

func main() {
	fmt.Print("AAA")
}
