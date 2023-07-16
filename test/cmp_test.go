package test

import (
	"github.com/Braun-Alex/SHA-3/pkg/cmp"
	"testing"
)

func TestXorOperation(t *testing.T) {
	testCases := []struct {
		x, y, expectedResult bool
	}{
		{x: false, y: false, expectedResult: false},
		{x: false, y: true, expectedResult: true},
		{x: true, y: false, expectedResult: true},
		{x: true, y: true, expectedResult: false},
	}
	for _, tc := range testCases {
		actualResult := cmp.Xor(tc.x, tc.y)
		if actualResult != tc.expectedResult {
			t.Errorf("XOR(%v, %v) != %v", tc.x, tc.y, tc.expectedResult)
		}
	}
}

func TestModOperation(t *testing.T) {
	testCases := []struct {
		a, p, expectedResult int
	}{
		{a: 0, p: 3, expectedResult: 0},
		{a: -3, p: 29, expectedResult: 26},
		{a: 33, p: 5, expectedResult: 3},
	}
	for _, tc := range testCases {
		actualResult := cmp.Mod(tc.a, tc.p)
		if actualResult != tc.expectedResult {
			t.Errorf("%v mod %v != %v", tc.a, tc.p, tc.expectedResult)
		}
	}
}
