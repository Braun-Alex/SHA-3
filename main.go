package sha3_512

import (
	"fmt"
	"golang.org/x/crypto/sha3"
	"time"
)

func testTime() {
	data := []byte("Value")
	start := time.Now()
	expectedResult := sha3.Sum512(data)
	elapsed := time.Since(start)
	fmt.Printf("Час виконання: %s\n", elapsed)
	start = time.Now()
	actualResult := sha3_512.SHA3_512(data)
	elapsed = time.Since(start)
	fmt.Printf("Час виконання: %s\n", elapsed)
}
