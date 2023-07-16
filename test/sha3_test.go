package test

import (
	isha3 "github.com/Braun-Alex/SHA-3/pkg/sha3"
	"golang.org/x/crypto/sha3"
	"testing"
)

func TestLessOneBlock(test *testing.T) {
	data := []byte("Value")
	expectedResult := sha3.Sum512(data)
	actualResult := isha3.Sum512(data)
	if expectedResult != actualResult {
		test.Errorf("Реалізована геш-функція працює некоректно на значенні: %x", data)
	}
}

func TestOneBlock(test *testing.T) {
	data := []byte("Hiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii")
	expectedResult := sha3.Sum512(data)
	actualResult := isha3.Sum512(data)
	if expectedResult != actualResult {
		test.Errorf("Реалізована геш-функція працює некоректно на значенні: %x", data)
	}
}

func TestMoreBlocks(test *testing.T) {
	data := []byte("Hiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii")
	expectedResult := sha3.Sum512(data)
	actualResult := isha3.Sum512(data)
	if expectedResult != actualResult {
		test.Errorf("Реалізована геш-функція працює некоректно на значенні: %x", data)
	}
}
