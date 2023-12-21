package tools

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func SignImageUrl(secret string, path string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(path))
	sign := base64.URLEncoding.EncodeToString(mac.Sum(nil))[:40]
	return fmt.Sprintf("%s/%s", sign, path)
}
