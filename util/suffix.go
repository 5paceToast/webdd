package util

import (
	"bytes"
)

// ReplaceSuffix replaces a before suffix with an after suffix
func ReplaceSuffix(s, before, after []byte) []byte {
	var b bytes.Buffer
	b.Write(bytes.TrimSuffix(s, before))
	b.Write(after)
	return b.Bytes()
}
