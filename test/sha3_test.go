package test

import (
	"encoding/hex"
	isha3 "github.com/Braun-Alex/SHA-3/pkg/sha3"
	"golang.org/x/crypto/sha3"
	"testing"
)

func TestLessOneBlock(t *testing.T) {
	data := []byte("Value")
	expectedResult := sha3.Sum512(data)
	actualResult := isha3.Sum512(data)
	if expectedResult != actualResult {
		t.Errorf("Expected: %s, actual: %s",
			hex.EncodeToString(expectedResult[:]),
			hex.EncodeToString(actualResult[:]))
	}
}

/*func TestOneBlock(t *testing.T) {
	data := []byte("Hiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii")
	expectedResult := sha3.Sum512(data)
	actualResult := isha3.Sum512(data)
	if expectedResult != actualResult {
		t.Errorf("Expected: %s, actual: %s",
			hex.EncodeToString(expectedResult[:]),
			hex.EncodeToString(actualResult[:]))
	}
}

func TestMoreBlocks(t *testing.T) {
	data := []byte("Hiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii")
	expectedResult := sha3.Sum512(data)
	actualResult := isha3.Sum512(data)
	if expectedResult != actualResult {
		t.Errorf("Expected: %s, actual: %s",
			hex.EncodeToString(expectedResult[:]),
			hex.EncodeToString(actualResult[:]))
	}
}*/
