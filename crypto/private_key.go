package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"io"
)

type PrivateKey struct {
	key ed25519.PrivateKey
}

func NewPrivateKeyFromString(s string) *PrivateKey {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return NewPrivateKeyFromSeed(b)
}

func NewPrivateKeyFromSeed(seed []byte) *PrivateKey {
	if len(seed) != seedLength {
		panic("invalid seed length")
	}
	return &PrivateKey{key: ed25519.NewKeyFromSeed(seed)}
}

func GeneratePrivateKey() *PrivateKey {
	seed := make([]byte, seedLength)
	_, err := io.ReadFull(rand.Reader, seed)
	if err != nil {
		panic(err)
	}
	return &PrivateKey{key: ed25519.NewKeyFromSeed(seed)}
}

func (p *PrivateKey) Bytes() []byte {
	return p.key
}

func (p *PrivateKey) Sign(message []byte) *Signature {
	return &Signature{
		value: ed25519.Sign(p.key, message),
	}
}

func (p *PrivateKey) String() string {
	return hex.EncodeToString(p.key)
}
