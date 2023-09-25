package util

import (
	"encoding/hex"
)

func Hash_Password(pass string) string {
	Hasher.Write([]byte(pass))
	var hashed_password string = hex.EncodeToString(Hasher.Sum(nil))
	Hasher.Reset()
	return hashed_password
}
