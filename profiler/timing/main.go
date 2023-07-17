package main

import (
	"fmt"
	implSha3 "github.com/Braun-Alex/SHA-3/pkg/sha3"
	"golang.org/x/crypto/sha3"
	"time"
)

const rounds = 100

func main() {
	data := []byte("StarkNet is an open-source, decentralized platform for building scalable " +
		"and secure applications on Ethereum. It is designed to provide high " +
		"throughput, low-cost transactions, and strong privacy guarantees " +
		"for smart contract execution. StarkNet achieves these goals " +
		"by utilizing a technology called zk-rollups, which allows " +
		"for off-chain execution of computations while maintaining " +
		"the security and trustlessness of the Ethereum blockchain. StarkNet " +
		"operates as a Layer 2 solution, meaning it operates on top of " +
		"the Ethereum mainnet, leveraging its security and decentralization. " +
		"It uses Zero-Knowledge Proofs to bundle and validate multiple " +
		"transactions off-chain, compressing them into a single proof that is then " +
		"submitted to the Ethereum mainnet for verification. This approach significantly " +
		"reduces the transaction fees and congestion on the mainnet while maintaining the " +
		"security and trust of Ethereum's consensus mechanism. StarkNet has the " +
		"potential to greatly enhance the scalability of decentralized applications " +
		"on Ethereum, enabling a wide range of use cases. By leveraging off-chain " +
		"computation and the security of the Ethereum mainnet, StarkNet aims to provide " +
		"a powerful infrastructure for building scalable and efficient " +
		"blockchain applications.")
	// Deferred invoking has been implemented for the correct displaying all the data in the console
	defer fmt.Printf("The text to hash was: \n\"%s\"", data)
	standardStart := time.Now()
	for i := 0; i < rounds; i++ {
		_ = sha3.Sum512(data)
	}
	standardDuration := time.Since(standardStart)
	defer fmt.Printf("Standard version of SHA3-512 hashes text on average within %v microseconds\n",
		standardDuration.Microseconds()/rounds)
	implStart := time.Now()
	for i := 0; i < rounds; i++ {
		_ = implSha3.Sum512(data)
	}
	implDuration := time.Since(implStart)
	defer fmt.Printf("Implemented version of SHA3-512 hashes text on average within %v microseconds\n",
		implDuration.Microseconds()/rounds)
}
