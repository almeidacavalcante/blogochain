package crypto

import "crypto/ed25519"

type Signature struct {
	value []byte
}

func (sig *Signature) Bytes() []byte {
	return sig.value
}

func (sig *Signature) Verify(pubKey *PublicKey, message []byte) bool {
	return ed25519.Verify(pubKey.key, message, sig.value)
}
