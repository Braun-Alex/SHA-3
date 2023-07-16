package test

import (
	"github.com/Braun-Alex/SHA-3/pkg/conv"
	"testing"
)

func TestConvertingByteToBits(t *testing.T) {
	testCases := []struct {
		byteToConvert byte
		expectedBits  []bool
	}{
		{byteToConvert: 0,
			expectedBits: []bool{false, false, false, false, false, false, false, false}},
		{byteToConvert: 3,
			expectedBits: []bool{false, false, false, false, false, false, true, true}},
		{byteToConvert: 15,
			expectedBits: []bool{false, false, false, false, true, true, true, true}},
		{byteToConvert: 33,
			expectedBits: []bool{false, false, true, false, false, false, false, true}},
		{byteToConvert: 255,
			expectedBits: []bool{true, true, true, true, true, true, true, true}},
	}
	for _, tc := range testCases {
		actualResult := conv.ByteToBits(tc.byteToConvert)
		for i := range actualResult {
			if actualResult[i] != tc.expectedBits[i] {
				t.Errorf("Byte %v does not convert to bits correctly", tc.byteToConvert)
				break
			}
		}
	}
}

func TestConvertingBitsToBytes(t *testing.T) {
	testCases := []struct {
		bitsToConvert []bool
		expectedBytes []byte
	}{
		{bitsToConvert: []bool{},
			expectedBytes: []byte{}},
		{bitsToConvert: []bool{false, false, false, false, false, false, false, false},
			expectedBytes: []byte{0}},
		{bitsToConvert: []bool{false, false, false, false, false, false, false, false,
			false, false, false, false, false, false, true, true,
			false, false, false, false, true, true, true, true,
			false, false, true, false, false, false, false, true,
			true, true, true, true, true, true, true, true,
		},
			expectedBytes: []byte{0, 3, 15, 33, 255}},
	}
	for _, tc := range testCases {
		actualResult := conv.BitsToBytes(tc.bitsToConvert)
		for i := range actualResult {
			if actualResult[i] != tc.expectedBytes[i] {
				t.Errorf("Bits %v does not convert to bytes correctly", tc.bitsToConvert)
				break
			}
		}
	}
}
