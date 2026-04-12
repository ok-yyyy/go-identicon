package digest

import (
	"crypto/md5"
)

// Sum computes the hash of the input data.
func Sum(data []byte) [16]byte {
	return md5.Sum(data)
}
