package crypto

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// HashEmail chuẩn hoá email rồi hash SHA-256 để làm key lưu OTP trong Redis.
func HashEmail(email string) string {
	normalized := strings.TrimSpace(strings.ToLower(email))
	sum := sha256.Sum256([]byte(normalized))
	return hex.EncodeToString(sum[:])
}
