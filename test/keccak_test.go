package test

import (
	"encoding/hex"
	"fmt"
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

func TestThetaFunction(t *testing.T) {
	data := "EE A4 41 43 C7 DE 5E 39 91 C9 B9 D8 AA 91 F0 61 8C 43 BB AF 67 3B F7 04 F6 " +
		"6C 4A 35 CD BF CF 6A 3B BC 9F 5A 78 F1 1F 2C 9D 16 F4 5D 31 C2 F4 33 06 BA " +
		"A9 FB E9 0C 5D 91 07 4F 52 1F 3F 25 51 76 AD E7 D2 88 63 12 66 EE 5E 0F " +
		"1C 19 CE 8D A7 33 68 EB E0 A0 84 93 48 14 EE C8 B5 E0 88 AC 28 68 80 " +
		"C1 80 F6 04 50 6B 17 37 38 70 FA 72 A5 60 A1 8A B0 73 65 40 E3 03 " +
		"42 83 C2 4E B9 BB 7A 48 35 9F 23 F3 C0 67 80 60 D0 7C DE 48 C1 " +
		"AB 7C EC 6B 82 25 DE D3 E7 C4 CD 48 4F E6 D1 92 23 AF FC C5 " +
		"EB A3 16 CF A7 FE 22 27 9A 99 88 9D 22 E7 95 A9 58 DC BA " +
		"66 7D 89 CF D8 84 49 58 69 CC 1A E1 FC FD C3 E3 0B 71 " +
		"37 D8 6C"
	state := ConvertDataToState(data)
	A := keccak.FromState(state)
	actualA := keccak.Theta(A)
	actualState := keccak.ToState(actualA)
	expectedData := "4A 80 3F C1 7E 74 21 0F BD 6F 83 D3 D1 ED A5 C0 25 FC B8 CB 06 C0 4C C6 " +
		"62 EF 55 AB 0E 0B 1F 55 B7 13 EB 37 52 31 EA A4 39 32 8A DF 88 68 8B 05 2A 1C 93 " +
		"F0 92 70 08 30 AE F0 51 7B 5E DE EA B4 39 64 CD 16 A0 A6 B6 D1 D2 A0 68 74 E4 " +
		"4D 52 BB CC CF 9E 22 3D 39 37 22 C2 6E 8F EB F3 D0 7D C9 29 7E 83 92 65 AB " +
		"D0 D5 A3 BB 6F 64 B1 11 B0 9E 06 1F 07 08 6A 23 F6 CA 27 E6 30 3B 02 D0 " +
		"37 03 B3 85 C9 CB 1C FC 35 71 D5 61 4B A5 CA 87 57 A9 16 A6 C1 4D 24 " +
		"70 1D 77 C3 49 A5 FF 09 6F 09 4D 4F 87 68 4D 1E 54 5D 11 B6 3F B2 " +
		"96 59 9B C0 08 F1 63 B9 02 1C 72 74 1A 10 CA 47 F7 0F AE 31 C3 " +
		"71 6C 97 66 5B F7 2D E4"
	expectedState := ConvertDataToState(expectedData)
	fmt.Println(hex.EncodeToString(conv.BitsToBytes(state[:])))
	fmt.Println(hex.EncodeToString(conv.BitsToBytes(actualState[:])))
	fmt.Print(hex.EncodeToString(conv.BitsToBytes(expectedState[:])))
	for i := range actualState {
		if actualState[i] != expectedState[i] {
			t.Error("Theta function has not been properly implemented")
			break
		}
	}
}

func TestRhoFunction(t *testing.T) {
	data := "4A 80 3F C1 7E 74 21 0F BD 6F 83 D3 D1 ED A5 C0 25 FC B8 CB 06 C0 4C C6 " +
		"62 EF 55 AB 0E 0B 1F 55 B7 13 EB 37 52 31 EA A4 39 32 8A DF 88 68 8B 05 2A 1C 93 " +
		"F0 92 70 08 30 AE F0 51 7B 5E DE EA B4 39 64 CD 16 A0 A6 B6 D1 D2 A0 68 74 E4 " +
		"4D 52 BB CC CF 9E 22 3D 39 37 22 C2 6E 8F EB F3 D0 7D C9 29 7E 83 92 65 AB " +
		"D0 D5 A3 BB 6F 64 B1 11 B0 9E 06 1F 07 08 6A 23 F6 CA 27 E6 30 3B 02 D0 " +
		"37 03 B3 85 C9 CB 1C FC 35 71 D5 61 4B A5 CA 87 57 A9 16 A6 C1 4D 24 " +
		"70 1D 77 C3 49 A5 FF 09 6F 09 4D 4F 87 68 4D 1E 54 5D 11 B6 3F B2 " +
		"96 59 9B C0 08 F1 63 B9 02 1C 72 74 1A 10 CA 47 F7 0F AE 31 C3 " +
		"71 6C 97 66 5B F7 2D E4"
	state := ConvertDataToState(data)
	A := keccak.FromState(state)
	actualA := keccak.Rho(A)
	actualState := keccak.ToState(actualA)
	expectedData := "4A 80 3F C1 7E 74 21 0F 7B DF 06 A7 A3 DB 4B 81 09 3F EE B2 01 30 93 71 " +
		"B0 F0 51 25 F6 5E B5 EA 8A 51 27 BD 9D 58 BF 91 8D 88 B6 58 90 23 A3 F8 09 2F 09 " +
		"87 00 A3 C2 31 AD 2B 7C D4 9E 97 B7 3A B2 66 0B 50 53 DB E8 1C 24 B5 2B 0D 8A " +
		"46 47 DE 61 7E F6 14 E9 C9 B9 11 25 0B BB 3D AE CF 43 F7 94 2C 5B 85 AE 4E " +
		"F1 1B 23 60 3D 47 77 DF C8 62 04 B5 11 7B 65 83 8F 03 76 04 A0 6F 06 4E " +
		"CC 61 79 99 83 BF 26 6E B6 30 AB D4 EA B0 A5 52 E5 C3 AE E3 CE C2 34 " +
		"B8 89 04 4D C3 49 A5 FF 09 6F 09 75 45 3C 1D A2 35 79 50 D8 FE C8 " +
		"5A 66 6D 02 23 7E 2C 57 80 43 8E 4E 23 CA 47 F7 0F AE 31 C3 10 " +
		"0B 79 1C DB A5 D9 D6 7D"
	expectedState := ConvertDataToState(expectedData)
	fmt.Println(hex.EncodeToString(conv.BitsToBytes(state[:])))
	fmt.Println(hex.EncodeToString(conv.BitsToBytes(actualState[:])))
	fmt.Print(hex.EncodeToString(conv.BitsToBytes(expectedState[:])))
	for i := range actualState {
		if actualState[i] != expectedState[i] {
			t.Error("Rho function has not been properly implemented")
			break
		}
	}
}

func TestPiFunction(t *testing.T) {
	data := "4A 80 3F C1 7E 74 21 0F 7B DF 06 A7 A3 DB 4B 81 09 3F EE B2 01 30 93 71 " +
		"B0 F0 51 25 F6 5E B5 EA 8A 51 27 BD 9D 58 BF 91 8D 88 B6 58 90 23 A3 F8 09 2F 09 " +
		"87 00 A3 C2 31 AD 2B 7C D4 9E 97 B7 3A B2 66 0B 50 53 DB E8 1C 24 B5 2B 0D 8A " +
		"46 47 DE 61 7E F6 14 E9 C9 B9 11 25 0B BB 3D AE CF 43 F7 94 2C 5B 85 AE 4E " +
		"F1 1B 23 60 3D 47 77 DF C8 62 04 B5 11 7B 65 83 8F 03 76 04 A0 6F 06 4E " +
		"CC 61 79 99 83 BF 26 6E B6 30 AB D4 EA B0 A5 52 E5 C3 AE E3 CE C2 34 " +
		"B8 89 04 4D C3 49 A5 FF 09 6F 09 75 45 3C 1D A2 35 79 50 D8 FE C8 " +
		"5A 66 6D 02 23 7E 2C 57 80 43 8E 4E 23 CA 47 F7 0F AE 31 C3 10 " +
		"0B 79 1C DB A5 D9 D6 7D"
	state := ConvertDataToState(data)
	A := keccak.FromState(state)
	actualA := keccak.Pi(A)
	actualState := keccak.ToState(actualA)
	expectedData := "4A 80 3F C1 7E 74 21 0F 09 2F 09 87 00 A3 C2 31 94 2C 5B 85 AE 4E F1 1B " +
		"AE E3 CE C2 34 B8 89 04 0B 79 1C DB A5 D9 D6 7D B0 F0 51 25 F6 5E B5 EA 24 B5 2B " +
		"0D 8A 46 47 DE 61 7E F6 14 E9 C9 B9 11 79 99 83 BF 26 6E B6 30 7E 2C 57 80 43 " +
		"8E 4E 23 7B DF 06 A7 A3 DB 4B 81 AD 2B 7C D4 9E 97 B7 3A 23 60 3D 47 77 DF " +
		"C8 62 4D C3 49 A5 FF 09 6F 09 75 45 3C 1D A2 35 79 50 8A 51 27 BD 9D 58 " +
		"BF 91 8D 88 B6 58 90 23 A3 F8 25 0B BB 3D AE CF 43 F7 AB D4 EA B0 A5 " +
		"52 E5 C3 CA 47 F7 0F AE 31 C3 10 09 3F EE B2 01 30 93 71 B2 66 0B " +
		"50 53 DB E8 1C 04 B5 11 7B 65 83 8F 03 76 04 A0 6F 06 4E CC 61 " +
		"D8 FE C8 5A 66 6D 02 23"
	expectedState := ConvertDataToState(expectedData)
	for i := range actualState {
		if actualState[i] != expectedState[i] {
			t.Error("Pi function has not been properly implemented")
			break
		}
	}
}

func TestChiFunction(t *testing.T) {
	data := "4A 80 3F C1 7E 74 21 0F 09 2F 09 87 00 A3 C2 31 94 2C 5B 85 AE 4E F1 1B " +
		"AE E3 CE C2 34 B8 89 04 0B 79 1C DB A5 D9 D6 7D B0 F0 51 25 F6 5E B5 EA 24 B5 2B " +
		"0D 8A 46 47 DE 61 7E F6 14 E9 C9 B9 11 79 99 83 BF 26 6E B6 30 7E 2C 57 80 43 " +
		"8E 4E 23 7B DF 06 A7 A3 DB 4B 81 AD 2B 7C D4 9E 97 B7 3A 23 60 3D 47 77 DF " +
		"C8 62 4D C3 49 A5 FF 09 6F 09 75 45 3C 1D A2 35 79 50 8A 51 27 BD 9D 58 " +
		"BF 91 8D 88 B6 58 90 23 A3 F8 25 0B BB 3D AE CF 43 F7 AB D4 EA B0 A5 " +
		"52 E5 C3 CA 47 F7 0F AE 31 C3 10 09 3F EE B2 01 30 93 71 B2 66 0B " +
		"50 53 DB E8 1C 04 B5 11 7B 65 83 8F 03 76 04 A0 6F 06 4E CC 61 " +
		"D8 FE C8 5A 66 6D 02 23"
	state := ConvertDataToState(data)
	A := keccak.FromState(state)
	actualA := keccak.Chi(A)
	actualState := keccak.ToState(actualA)
	expectedData := "DE 80 6D C1 D0 38 10 05 23 EC 8D C5 10 13 CA 35 95 34 4B 9C 2F 0F " +
		"A7 62 EE 63 ED C2 6E 9C A8 06 0A 56 1C DD A5 5A 14 4D F1 BA 85 35 97 D7 0D " +
		"EB 3C 34 2A A6 8C 60 41 FE 67 5A A2 14 A8 49 F1 12 F9 49 83 9A 92 3E 07 " +
		"F8 7A 29 7D 88 4B 8E 0C 37 79 9F 07 A4 C2 93 03 C1 E1 A8 3C 74 16 97 " +
		"90 33 13 64 09 5F 77 EB D8 32 47 59 4B 07 FE C3 6D 88 F1 65 44 4D " +
		"BE 31 CD 6A AA 52 2E 98 B3 94 FF 96 07 5C F6 D8 91 33 07 F8 65 " +
		"08 AE 32 A4 EE 41 E7 AB C4 EA 00 B4 1A D9 42 CF CF 67 4F AE " +
		"12 C3 78 0D AE FE 99 25 30 94 72 C0 66 AB 54 51 97 A8 " +
		"7C 8C 4F 59 6B 05 A2 8D 01 77 05 86 CF 07 5E 5D 31 " +
		"6A BE C9 1A 34 A6 6A 2F"
	expectedState := ConvertDataToState(expectedData)
	for i := range actualState {
		if actualState[i] != expectedState[i] {
			t.Error("Chi function has not been properly implemented")
			break
		}
	}
}

func TestIotaFunction(t *testing.T) {
	data := "DE 80 6D C1 D0 38 10 05 23 EC 8D C5 10 13 CA 35 95 34 4B 9C 2F 0F " +
		"A7 62 EE 63 ED C2 6E 9C A8 06 0A 56 1C DD A5 5A 14 4D F1 BA 85 35 97 D7 0D " +
		"EB 3C 34 2A A6 8C 60 41 FE 67 5A A2 14 A8 49 F1 12 F9 49 83 9A 92 3E 07 " +
		"F8 7A 29 7D 88 4B 8E 0C 37 79 9F 07 A4 C2 93 03 C1 E1 A8 3C 74 16 97 " +
		"90 33 13 64 09 5F 77 EB D8 32 47 59 4B 07 FE C3 6D 88 F1 65 44 4D " +
		"BE 31 CD 6A AA 52 2E 98 B3 94 FF 96 07 5C F6 D8 91 33 07 F8 65 " +
		"08 AE 32 A4 EE 41 E7 AB C4 EA 00 B4 1A D9 42 CF CF 67 4F AE " +
		"12 C3 78 0D AE FE 99 25 30 94 72 C0 66 AB 54 51 97 A8 " +
		"7C 8C 4F 59 6B 05 A2 8D 01 77 05 86 CF 07 5E 5D 31 " +
		"6A BE C9 1A 34 A6 6A 2F"
	state := ConvertDataToState(data)
	A := keccak.FromState(state)
	actualA := keccak.Iota(A, 0)
	actualState := keccak.ToState(actualA)
	expectedData := "DF 80 6D C1 D0 38 10 05 23 EC 8D C5 10 13 CA 35 95 34 4B 9C 2F 0F " +
		"A7 62 EE 63 ED C2 6E 9C A8 06 0A 56 1C DD A5 5A 14 4D F1 BA 85 35 97 D7 0D " +
		"EB 3C 34 2A A6 8C 60 41 FE 67 5A A2 14 A8 49 F1 12 F9 49 83 9A 92 3E 07 " +
		"F8 7A 29 7D 88 4B 8E 0C 37 79 9F 07 A4 C2 93 03 C1 E1 A8 3C 74 16 97 " +
		"90 33 13 64 09 5F 77 EB D8 32 47 59 4B 07 FE C3 6D 88 F1 65 44 4D " +
		"BE 31 CD 6A AA 52 2E 98 B3 94 FF 96 07 5C F6 D8 91 33 07 F8 65 " +
		"08 AE 32 A4 EE 41 E7 AB C4 EA 00 B4 1A D9 42 CF CF 67 4F AE " +
		"12 C3 78 0D AE FE 99 25 30 94 72 C0 66 AB 54 51 97 A8 7C " +
		"8C 4F 59 6B 05 A2 8D 01 77 05 86 CF 07 5E 5D 31 6A BE " +
		"C9 1A 34 A6 6A 2F"
	expectedState := ConvertDataToState(expectedData)
	for i := range actualState {
		if actualState[i] != expectedState[i] {
			t.Error("Iota function has not been properly implemented")
			break
		}
	}
}
