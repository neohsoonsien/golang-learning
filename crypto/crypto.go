package crypto

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
)

func DerivePublicPrivateKey() {
	// Decode a hex-encoded private key string to []byte
	privateKeyBytes, err := hex.DecodeString("67789d8a530ab6838ab497dbd2615195825dd3e87a1a9db663f5ea17a15887c1")
	if err != nil {
		fmt.Println(err)
		return
	}

	// derive private key and public key from []byte to *secp256k1.PrivateKey and *secp256k1.PrivateKey, respectively
	privateKey, publicKey := btcec.PrivKeyFromBytes(privateKeyBytes)

	fmt.Printf("The public key is: %v, and the private key is: %v", publicKey, privateKey)
}
