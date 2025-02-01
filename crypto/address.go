package crypto

import "encoding/hex"

type Address struct {
	value []byte
}

func (a Address) String() string {
	return hex.EncodeToString(a.value)
}
