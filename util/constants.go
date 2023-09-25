package util

import (
	"hash"

	"gorm.io/gorm"
)

var DB_Connection *gorm.DB

var Hasher hash.Hash
