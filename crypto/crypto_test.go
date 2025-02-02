package crypto

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	assert.Equal(t, len(privKey.Bytes()), privateKeyLength)
	pubKey := privKey.PublicKey()
	assert.Equal(t, len(pubKey.Bytes()), publicKeyLength)
}

func TestPrivateKey_SignWithValidMessage(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	message := []byte("hello")

	signature := privKey.Sign(message)

	assert.True(t, signature.Verify(pubKey, message))
}

func TestPrivateKey_SignWithInvalidMessage(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	originalMessage := []byte("hello")
	wrongMessage := []byte("world")
	signature := privKey.Sign(originalMessage)

	assert.False(t, signature.Verify(pubKey, wrongMessage))
}

func TestPrivateKey_SignWithWrongPublicKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	message := []byte("hello")
	signature := privKey.Sign(message)

	anotherPrivKey := GeneratePrivateKey()
	anotherPubKey := anotherPrivKey.PublicKey()

	assert.False(t, signature.Verify(anotherPubKey, message))
}

func TestPrivateKey_SignWithInvalidSignature(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	message := []byte("hello")
	invalidSignature := &Signature{value: []byte("invalid")}

	assert.False(t, invalidSignature.Verify(pubKey, message))
}

func TestAddress(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	address := pubKey.Address()

	assert.Equal(t, len(address.value), addressLength)
	assert.Equal(t, len(address.Bytes()), addressLength)
	fmt.Println(pubKey.String())
	fmt.Println(address)
}

func TestNewPrivateKeyFromString(t *testing.T) {
	seed := "f327edf67f00184f535f49596bc5d3080f918b9584ca7b6b21d234e42c651c6d"
	addressString := "c7006c0b769d3c9fe2cbbe9e9369bfe91acc9a45"

	privKey := NewPrivateKeyFromString(seed)
	assert.Equal(t, privateKeyLength, len(privKey.Bytes()))
	address := privKey.PublicKey().Address()
	assert.Equal(t, addressString, address.String())
	assert.Equal(t, addressLength, len(address.Bytes()))
}
