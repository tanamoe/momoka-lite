package tools

import (
	"crypto/sha256"
	"fmt"
)

func SHA256(inp string) string {
	h := sha256.New()
	h.Write([]byte(inp))
	return fmt.Sprintf("%x", h.Sum(nil))
}
