package crypto

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
)

func DerivePublicPrivateKey() {
	// Decode a hex-encoded private key string to []byte
	privateKeyBytes, err := hex.DecodeString("67789d8a530ab6838ab497dbd2615195825dd3e87a1a9db663f5ea17a15887c1")
	if err != nil {
		fmt.Println(err)
		return
	}

	// derive private key and public key from []byte to *secp256k1.PrivateKey and *secp256k1.PrivateKey, respectively
	privateKeySecp256k1, publicKeySecp256k1 := btcec.PrivKeyFromBytes(privateKeyBytes)
	fmt.Printf("The public key is: %v, and the private key is: %v\n", publicKeySecp256k1, privateKeySecp256k1)

	// private key in ECDSA
	privateKeyEcdsa := privateKeySecp256k1.ToECDSA()
	fmt.Printf("The private key in ECDSA is: %v\n", privateKeyEcdsa)

	// prepare message
	message := "Example message"
	data := fmt.Sprintf("\x19Tron Signed Message:\n%d%s", len(message), message)
	//messageHash := chainhash.DoubleHashB([]byte(data))

	// sign the message
	// signature := ecdsa.Sign(privateKeySecp256k1, messageHash)
	hash, err := hex.DecodeString(data)
	if err != nil {
		fmt.Errorf("Error in hex.DecodeString, err: %v", err)
	}

	signature, err := ecdsa.SignCompact(privateKeySecp256k1, hash, false)
	if err != nil {
		fmt.Errorf("Error in ecdsa.SignCompact, err: %v", err)
	}

	// Serialize and display the signature.
	fmt.Printf("Serialized Signature: %x\n", hex.EncodeToString(signature))

	// // Verify the signature for the message using the public key.
	// verified := signature.Verify(messageHash, publicKeySecp256k1)
	// fmt.Printf("Signature Verified? %v\n", verified)
}
