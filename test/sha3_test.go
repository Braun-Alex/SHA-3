package test

import (
	implSha3 "github.com/Braun-Alex/SHA-3/pkg/sha3"
	"golang.org/x/crypto/sha3"
	"testing"
)

func TestShortMessage(t *testing.T) {
	data := []byte("ZK-STARK has big impact on StarkNet")
	expectedResult := sha3.Sum512(data)
	actualResult := implSha3.Sum512(data)
	if expectedResult != actualResult {
		t.Errorf("Implemented SHA3-%v function does not properly hash message with length less than %v bits",
			implSha3.D, implSha3.Rate)
	}
}

func TestOneBlockMessage(t *testing.T) {
	data := []byte("StarkNet is an open-source, decentralized platform for dApps on Ethereum")
	expectedResult := sha3.Sum512(data)
	actualResult := implSha3.Sum512(data)
	if expectedResult != actualResult {
		t.Errorf("Implemented SHA3-%v function does not properly hash message with a length that equals %v bits",
			implSha3.D, implSha3.Rate)
	}
}

func TestLongMessage(t *testing.T) {
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
	expectedResult := sha3.Sum512(data)
	actualResult := implSha3.Sum512(data)
	if expectedResult != actualResult {
		t.Errorf("Implemented SHA3-%v function does not properly hash message with length greater than %v bits",
			implSha3.D, implSha3.Rate)
	}
}
