package test

import (
	"encoding/hex"
	"github.com/Braun-Alex/SHA-3/pkg/conv"
	"github.com/Braun-Alex/SHA-3/pkg/keccak"
	"strings"
	"testing"
)

func ConvertDataToState(data string) [keccak.B]bool {
	data = strings.Replace(data, " ", "", -1)
	data = strings.ToLower(data)
	bytes, err := hex.DecodeString(data)
	if err != nil {
		panic(err)
	}
	stateSlice := conv.BytesToBits(bytes)
	var state [keccak.B]bool
	for i := range stateSlice {
		state[i] = stateSlice[i]
	}
	return state
}

/*
Unit tests data for the Keccak permutation and
functions: https://csrc.nist.gov/CSRC/media/Projects/Cryptographic-Standards-and-Guidelines/documents/examples/SHA3-512_1600.pdf
*/

func TestKeccakPermutation(t *testing.T) {
	data := "EE A4 41 43 C7 DE 5E 39 91 C9 B9 D8 AA 91 F0 61 8C 43 BB AF 67 3B F7 04 F6 6C " +
		"4A 35 CD BF CF 6A 3B BC 9F 5A 78 F1 1F 2C 9D 16 F4 5D 31 C2 F4 33 06 BA A9 FB E9 " +
		"0C 5D 91 07 4F 52 1F 3F 25 51 76 AD E7 D2 88 63 12 66 EE 5E 0F 1C 19 CE 8D A7 " +
		"33 68 EB E0 A0 84 93 48 14 EE C8 B5 E0 88 AC 28 68 80 C1 80 F6 04 50 6B 17 " +
		"37 38 70 FA 72 A5 60 A1 8A B0 73 65 40 E3 03 42 83 C2 4E B9 BB 7A 48 35 " +
		"9F 23 F3 C0 67 80 60 D0 7C DE 48 C1 AB 7C EC 6B 82 25 DE D3 E7 C4 CD " +
		"48 4F E6 D1 92 23 AF FC C5 EB A3 16 CF A7 FE 22 27 9A 99 88 9D 22 " +
		"E7 95 A9 58 DC BA 66 7D 89 CF D8 84 49 58 69 CC 1A E1 FC FD C3 " +
		"E3 0B 71 37 D8 6C"
	state := ConvertDataToState(data)
	actualState := keccak.Permute(state)
	expectedData := "E7 6D FA D2 20 84 A8 B1 46 7F CF 2F FA 58 36 1B EC 76 28 ED F5 F3 " +
		"FD C0 E4 80 5D C4 8C AE EC A8 1B 7C 13 C3 0A DF 52 A3 65 95 84 73 9A 2D F4 " +
		"6B E5 89 C5 1C A1 A4 A8 41 6D F6 54 5A 1C E8 BA 00 86 15 6D 2B 6C 44 A5 " +
		"89 CA ED 88 69 AC FD 4B CF EA 28 6E 9D 1F 26 EC AA 39 EB 6E 49 66 58 " +
		"85 C8 95 6E 91 B2 33 20 AA 5E D0 03 63 FB A6 E7 77 5C 70 E1 82 79 " +
		"44 5B 95 A9 8E C5 C4 5E 93 48 99 6B BC B9 46 5A 17 C2 CA 3E FB " +
		"A9 2B 0E AF D9 F1 AD CC 3C 7D 8C 20 13 E8 E6 A2 16 AB 6A 55 " +
		"A0 8A 58 D9 7C 4A E2 B1 B8 7B 1A 59 9D FE 84 93 12 52 91 " +
		"9F 76 37 83 72 20 D7 CE 4D A0 20 96 F3 73 55 8A 46 99 " +
		"A3 FA 9D A7 69 92"
	expectedState := ConvertDataToState(expectedData)
	for i := range actualState {
		if actualState[i] != expectedState[i] {
			t.Error("Keccak permutation has not been properly implemented")
			break
		}
	}
}

func TestThetaFunction(t *testing.T) {
	data := "9C 8C E4 4A 56 69 F3 00 95 32 C3 5E A7 A5 F5 23 B1 D6 4B 03 41 FB CD 6E 71 63 " +
		"01 7E 61 AA 78 4A 14 01 56 EF F6 DD CD C0 C3 4D D2 1F A8 21 37 8D C7 97 1E AD 07 " +
		"53 35 22 6D 2C 72 8E 58 CA B7 F2 34 98 7C C8 16 B9 82 F7 A5 1D 54 20 81 5A F3 " +
		"45 06 E3 E4 7B 40 55 4D 59 48 8E E7 36 DE DA 7F 97 7A 90 9B 38 74 6D 2D 69 " +
		"CF B2 56 31 48 C6 97 5B 33 D1 E7 81 D6 64 A3 70 77 00 59 F2 20 FB 59 28 " +
		"A8 EC 61 F4 68 7D 67 4E B7 97 76 A2 A7 01 97 C6 B2 E5 33 C8 F1 9B 3D " +
		"FB 74 30 A5 16 E8 11 3D 53 A2 7F 97 A8 12 35 4A 79 B3 42 F9 B3 07 " +
		"DF CC E9 86 81 64 09 5C 67 86 47 5A 39 78 1D DD CD 10 9B 36 96 " +
		"3B 85 19 A1 1B 35"
	state := ConvertDataToState(data)
	A := keccak.FromState(state)
	actualA := keccak.Theta(A)
	actualState := keccak.ToState(actualA)
	expectedData := "5E ED DA 92 25 26 61 F1 37 96 BF 17 07 03 E3 4E 75 79 28 25 76 73 58 " +
		"53 67 C8 47 DA 56 36 48 19 6F 2F 0E 55 FD F9 B8 4D 01 2C EC C7 DB 6E A5 7C 65 " +
		"33 62 E4 A7 F5 23 4F A9 83 11 A8 6F 42 22 CF 22 33 3A 6C 21 25 B2 A4 DE 33 " +
		"0C 9A 8A 7E 86 C8 C4 82 DA A3 33 1A DF A8 EA 2A 9B 7F 7E 7C 69 FA BE 3F " +
		"F8 1E 43 E5 B8 54 D9 19 10 95 7F 5A A7 08 48 FF BF 3B DD 40 D6 FD B5 " +
		"61 67 2A 53 B4 CB D9 0A 48 1D BD C8 DB 71 23 73 38 15 84 90 89 02 " +
		"FB A4 4E 75 6C C6 07 0D A8 0F 1E FD AC E3 35 48 DE 60 1E A9 70 " +
		"61 7A D8 88 11 E6 85 FA A7 79 DA 84 42 2E 07 2F 6B EF 13 7A " +
		"4C 92 3E B9 EA 51 20 C8 4D B8 63 3F 12 85 6E B8"
	expectedState := ConvertDataToState(expectedData)
	for i := range actualState {
		if actualState[i] != expectedState[i] {
			t.Error("Theta function has not been properly implemented")
			break
		}
	}
}

func TestRhoFunction(t *testing.T) {
	data := "5E ED DA 92 25 26 61 F1 37 96 BF 17 07 03 E3 4E 75 79 28 25 76 73 58 " +
		"53 67 C8 47 DA 56 36 48 19 6F 2F 0E 55 FD F9 B8 4D 01 2C EC C7 DB 6E A5 7C 65 " +
		"33 62 E4 A7 F5 23 4F A9 83 11 A8 6F 42 22 CF 22 33 3A 6C 21 25 B2 A4 DE 33 " +
		"0C 9A 8A 7E 86 C8 C4 82 DA A3 33 1A DF A8 EA 2A 9B 7F 7E 7C 69 FA BE 3F " +
		"F8 1E 43 E5 B8 54 D9 19 10 95 7F 5A A7 08 48 FF BF 3B DD 40 D6 FD B5 " +
		"61 67 2A 53 B4 CB D9 0A 48 1D BD C8 DB 71 23 73 38 15 84 90 89 02 " +
		"FB A4 4E 75 6C C6 07 0D A8 0F 1E FD AC E3 35 48 DE 60 1E A9 70 " +
		"61 7A D8 88 11 E6 85 FA A7 79 DA 84 42 2E 07 2F 6B EF 13 7A " +
		"4C 92 3E B9 EA 51 20 C8 4D B8 63 3F 12 85 6E B8"
	state := ConvertDataToState(data)
	A := keccak.FromState(state)
	actualA := keccak.Rho(A)
	actualState := keccak.ToState(actualA)
	expectedData := "5E ED DA 92 25 26 61 F1 6E 2C 7F 2F 0E 06 C6 9D 5D 1E 4A 89 " +
		"DD 1C D6 54 65 83 94 71 86 7C A4 6D CF C7 6D 7A 7B 71 A8 EA BC ED 56 " +
		"CA 17 C0 C2 7E 46 7E 5A 3F F2 54 36 23 73 EA 60 04 EA 9B 90 C8 19 " +
		"1D B6 90 12 59 52 91 67 88 EC 3D C3 A0 A9 E8 25 16 D4 1E 9D D1 " +
		"F8 46 E9 AB AB 6C FE F9 F1 A5 F7 18 2A C7 A5 F2 FD C1 B4 4E " +
		"11 B2 33 20 2A FF 9D 6E 20 EB 7E A4 FF DF 54 A6 68 97 B3 " +
		"6B C3 CE A3 17 79 3B 6E 44 01 A9 81 FD 39 9C 0A 42 C8 " +
		"44 A0 01 95 D4 A9 8E CD F8 DE 0F 1E FD AC E3 35 48 " +
		"61 23 82 79 A4 C2 85 E9 46 98 17 EA 9F E6 69 13 " +
		"C8 E5 E0 65 ED 7D 42 4F 92 3E B9 EA 51 20 C8 " +
		"4C 1B 6E 13 EE D8 8F 44 A1"
	expectedState := ConvertDataToState(expectedData)
	for i := range actualState {
		if actualState[i] != expectedState[i] {
			t.Error("Rho function has not been properly implemented")
			break
		}
	}
}

func TestPiFunction(t *testing.T) {
	data := "5E ED DA 92 25 26 61 F1 6E 2C 7F 2F 0E 06 C6 9D 5D 1E 4A 89 " +
		"DD 1C D6 54 65 83 94 71 86 7C A4 6D CF C7 6D 7A 7B 71 A8 EA BC ED 56 " +
		"CA 17 C0 C2 7E 46 7E 5A 3F F2 54 36 23 73 EA 60 04 EA 9B 90 C8 19 " +
		"1D B6 90 12 59 52 91 67 88 EC 3D C3 A0 A9 E8 25 16 D4 1E 9D D1 " +
		"F8 46 E9 AB AB 6C FE F9 F1 A5 F7 18 2A C7 A5 F2 FD C1 B4 4E " +
		"11 B2 33 20 2A FF 9D 6E 20 EB 7E A4 FF DF 54 A6 68 97 B3 " +
		"6B C3 CE A3 17 79 3B 6E 44 01 A9 81 FD 39 9C 0A 42 C8 " +
		"44 A0 01 95 D4 A9 8E CD F8 DE 0F 1E FD AC E3 35 48 " +
		"61 23 82 79 A4 C2 85 E9 46 98 17 EA 9F E6 69 13 " +
		"C8 E5 E0 65 ED 7D 42 4F 92 3E B9 EA 51 20 C8 " +
		"4C 1B 6E 13 EE D8 8F 44 A1"
	state := ConvertDataToState(data)
	A := keccak.FromState(state)
	actualA := keccak.Pi(A)
	actualState := keccak.ToState(actualA)
	expectedData := "5E ED DA 92 25 26 61 F1 46 7E 5A 3F F2 54 36 23 F7 18 2A C7 A5 " +
		"F2 FD C1 A0 01 95 D4 A9 8E CD F8 1B 6E 13 EE D8 8F 44 A1 65 83 94 71 86 " +
		"7C A4 6D 67 88 EC 3D C3 A0 A9 E8 25 16 D4 1E 9D D1 F8 46 A3 17 79 3B " +
		"6E 44 01 A9 C8 E5 E0 65 ED 7D 42 4F 6E 2C 7F 2F 0E 06 C6 9D 73 EA " +
		"60 04 EA 9B 90 C8 B4 4E 11 B2 33 20 2A FF DE 0F 1E FD AC E3 35 " +
		"48 61 23 82 79 A4 C2 85 E9 CF C7 6D 7A 7B 71 A8 EA BC ED 56 " +
		"CA 17 C0 C2 7E E9 AB AB 6C FE F9 F1 A5 81 FD 39 9C 0A 42 " +
		"C8 44 92 3E B9 EA 51 20 C8 4C 5D 1E 4A 89 DD 1C D6 54 " +
		"19 1D B6 90 12 59 52 91 9D 6E 20 EB 7E A4 FF DF 54 " +
		"A6 68 97 B3 6B C3 CE 46 98 17 EA 9F E6 69 13"
	expectedState := ConvertDataToState(expectedData)
	for i := range actualState {
		if actualState[i] != expectedState[i] {
			t.Error("Pi function has not been properly implemented")
			break
		}
	}
}

func TestChiFunction(t *testing.T) {
	data := "5E ED DA 92 25 26 61 F1 46 7E 5A 3F F2 54 36 23 F7 18 2A C7 A5 " +
		"F2 FD C1 A0 01 95 D4 A9 8E CD F8 1B 6E 13 EE D8 8F 44 A1 65 83 94 71 86 " +
		"7C A4 6D 67 88 EC 3D C3 A0 A9 E8 25 16 D4 1E 9D D1 F8 46 A3 17 79 3B " +
		"6E 44 01 A9 C8 E5 E0 65 ED 7D 42 4F 6E 2C 7F 2F 0E 06 C6 9D 73 EA " +
		"60 04 EA 9B 90 C8 B4 4E 11 B2 33 20 2A FF DE 0F 1E FD AC E3 35 " +
		"48 61 23 82 79 A4 C2 85 E9 CF C7 6D 7A 7B 71 A8 EA BC ED 56 " +
		"CA 17 C0 C2 7E E9 AB AB 6C FE F9 F1 A5 81 FD 39 9C 0A 42 " +
		"C8 44 92 3E B9 EA 51 20 C8 4C 5D 1E 4A 89 DD 1C D6 54 " +
		"19 1D B6 90 12 59 52 91 9D 6E 20 EB 7E A4 FF DF 54 " +
		"A6 68 97 B3 6B C3 CE 46 98 17 EA 9F E6 69 13"
	state := ConvertDataToState(data)
	A := keccak.FromState(state)
	actualA := keccak.Chi(A)
	actualState := keccak.ToState(actualA)
	expectedData := "EF ED FA 52 20 84 A8 31 46 7F CF 2F FA 58 36 1B EC 76 28 ED F5 F3 " +
		"FD C0 E4 80 5D C4 8C AE EC A8 1B 7C 13 C3 0A DF 52 A3 65 95 84 73 9A 2D F4 " +
		"6B E5 89 C5 1C A1 A4 A8 41 6D F6 54 5A 1C E8 BA 00 86 15 6D 2B 6C 44 A5 " +
		"89 CA ED 88 69 AC FD 4B CF EA 28 6E 9D 1F 26 EC AA 39 EB 6E 49 66 58 " +
		"85 C8 95 6E 91 B2 33 20 AA 5E D0 03 63 FB A6 E7 77 5C 70 E1 82 79 " +
		"44 5B 95 A9 8E C5 C4 5E 93 48 99 6B BC B9 46 5A 17 C2 CA 3E FB " +
		"A9 2B 0E AF D9 F1 AD CC 3C 7D 8C 20 13 E8 E6 A2 16 AB 6A 55 " +
		"A0 8A 58 D9 7C 4A E2 B1 B8 7B 1A 59 9D FE 84 93 12 52 91 " +
		"9F 76 37 83 72 20 D7 CE 4D A0 20 96 F3 73 55 8A 46 99 " +
		"A3 FA 9D A7 69 92"
	expectedState := ConvertDataToState(expectedData)
	for i := range actualState {
		if actualState[i] != expectedState[i] {
			t.Error("Chi function has not been properly implemented")
			break
		}
	}
}

func TestIotaFunction(t *testing.T) {
	data := "EF ED FA 52 20 84 A8 31 46 7F CF 2F FA 58 36 1B EC 76 28 ED F5 F3 " +
		"FD C0 E4 80 5D C4 8C AE EC A8 1B 7C 13 C3 0A DF 52 A3 65 95 84 73 9A 2D F4 " +
		"6B E5 89 C5 1C A1 A4 A8 41 6D F6 54 5A 1C E8 BA 00 86 15 6D 2B 6C 44 A5 " +
		"89 CA ED 88 69 AC FD 4B CF EA 28 6E 9D 1F 26 EC AA 39 EB 6E 49 66 58 " +
		"85 C8 95 6E 91 B2 33 20 AA 5E D0 03 63 FB A6 E7 77 5C 70 E1 82 79 " +
		"44 5B 95 A9 8E C5 C4 5E 93 48 99 6B BC B9 46 5A 17 C2 CA 3E FB " +
		"A9 2B 0E AF D9 F1 AD CC 3C 7D 8C 20 13 E8 E6 A2 16 AB 6A 55 " +
		"A0 8A 58 D9 7C 4A E2 B1 B8 7B 1A 59 9D FE 84 93 12 52 91 " +
		"9F 76 37 83 72 20 D7 CE 4D A0 20 96 F3 73 55 8A 46 99 " +
		"A3 FA 9D A7 69 92"
	state := ConvertDataToState(data)
	A := keccak.FromState(state)
	actualA := keccak.Iota(A, 23)
	actualState := keccak.ToState(actualA)
	expectedData := "E7 6D FA D2 20 84 A8 B1 46 7F CF 2F FA 58 36 1B EC 76 28 ED F5 F3 " +
		"FD C0 E4 80 5D C4 8C AE EC A8 1B 7C 13 C3 0A DF 52 A3 65 95 84 73 9A 2D F4 " +
		"6B E5 89 C5 1C A1 A4 A8 41 6D F6 54 5A 1C E8 BA 00 86 15 6D 2B 6C 44 A5 " +
		"89 CA ED 88 69 AC FD 4B CF EA 28 6E 9D 1F 26 EC AA 39 EB 6E 49 66 58 " +
		"85 C8 95 6E 91 B2 33 20 AA 5E D0 03 63 FB A6 E7 77 5C 70 E1 82 79 " +
		"44 5B 95 A9 8E C5 C4 5E 93 48 99 6B BC B9 46 5A 17 C2 CA 3E FB " +
		"A9 2B 0E AF D9 F1 AD CC 3C 7D 8C 20 13 E8 E6 A2 16 AB 6A 55 " +
		"A0 8A 58 D9 7C 4A E2 B1 B8 7B 1A 59 9D FE 84 93 12 52 91 " +
		"9F 76 37 83 72 20 D7 CE 4D A0 20 96 F3 73 55 8A 46 99 " +
		"A3 FA 9D A7 69 92"
	expectedState := ConvertDataToState(expectedData)
	for i := range actualState {
		if actualState[i] != expectedState[i] {
			t.Error("Iota function has not been properly implemented")
			break
		}
	}
}
