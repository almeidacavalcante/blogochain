package crypto

import (
	"crypto/ed25519"
	"encoding/hex"
)

type PublicKey struct {
	key ed25519.PublicKey
}

func (p *PublicKey) Bytes() []byte {
	return p.key
}

func (p *PrivateKey) PublicKey() *PublicKey {
	b := make([]byte, publicKeyLength)
	copy(b, p.key[32:])
	return &PublicKey{key: b}
}

func (p *PublicKey) Address() Address {
	return Address{
		value: p.key[len(p.key)-addressLength:], // last 20 bytes
	}
}

func (p *PublicKey) String() string {
	return hex.EncodeToString(p.key)
}

func (a Address) Bytes() []byte {
	return a.value
}
