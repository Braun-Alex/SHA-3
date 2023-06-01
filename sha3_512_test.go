package sha3_512_test

import (
	sha3_512 "SHA-3"
	"golang.org/x/crypto/sha3"
	"testing"
)

func TestLessOneBlock(test *testing.T) {
	data := []byte("Value")
	expectedResult := sha3.Sum512(data)
	actualResult := sha3_512.SHA3_512(data)
	if expectedResult != actualResult {
		test.Errorf("Реалізована геш-функція працює некоректно на значенні: %x", data)
	}
}

func TestOneBlock(test *testing.T) {
	data := []byte("Hiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii")
	expectedResult := sha3.Sum512(data)
	actualResult := sha3_512.SHA3_512(data)
	if expectedResult != actualResult {
		test.Errorf("Реалізована геш-функція працює некоректно на значенні: %x", data)
	}
}

func TestMoreBlocks(test *testing.T) {
	data := []byte("Hiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii")
	expectedResult := sha3.Sum512(data)
	actualResult := sha3_512.SHA3_512(data)
	if expectedResult != actualResult {
		test.Errorf("Реалізована геш-функція працює некоректно на значенні: %x", data)
	}
}
