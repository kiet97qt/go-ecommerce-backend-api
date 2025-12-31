package random

import (
	"crypto/rand"
	"math/big"
)

// NumericOTP sinh OTP chỉ gồm chữ số với độ dài cho trước.
func NumericOTP(length int) (string, error) {
	const digits = "0123456789"
	res := make([]byte, length)
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		res[i] = digits[n.Int64()]
	}
	return string(res), nil
}
