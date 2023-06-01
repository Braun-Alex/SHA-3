package main

import "fmt"

func sha3(data []byte) {
	bits := make([]bool, 8*len(data))
	for i, byteValue := range data {
		for j := 0; j < 8; j++ {
			bits[8*i+j] = (byteValue & (1 << byte(7-j))) == 1
		}
	}
}

func main() {
	fmt.Print("AAA")
}
