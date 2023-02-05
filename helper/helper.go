package helper

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func IsAuthTokenValid(authBearer string) bool {
	token := strings.Replace(authBearer, "Bearer ", "", 1)
	tokenData := strings.Split(token, ":")

	apiKey := tokenData[0]
	nonceStr := tokenData[1]
	signature := tokenData[2]

	nonce, err := strconv.ParseInt(nonceStr, 10, 64)
	if err != nil {
		return false
	}

	selfSignedSignature := generateSelfSignature(apiKey, nonce)
	return selfSignedSignature == signature
}

func generateSelfSignature(apiKey string, nonce int64) string {
	secretKey := []byte("SIMPLESHOPPINGTEST")
	content := fmt.Sprintf("%s:%d", apiKey, nonce)
	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(content))
	return hex.EncodeToString(mac.Sum(nil))
}
