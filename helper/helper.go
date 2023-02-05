package helper

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type LogForm struct {
	Request   interface{} `json:"request"`
	Response  interface{} `json:"response"`
	CreatedAt time.Time   `json:"created_at"`
}

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

func LogToFile(fileName string, request interface{}, response interface{}) {
	destPath := path.Join("./logs/", fileName)
	f, _ := os.OpenFile(destPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	defer f.Close()

	WriteLog(f, request, response)
	log.SetOutput(os.Stdout)
}

func WriteLog(writer io.Writer, request interface{}, response interface{}) {
	var data LogForm
	data.Request = request
	data.Response = response
	data.CreatedAt = time.Now()

	log.SetFlags(0)
	log.SetOutput(writer)
	b, _ := json.Marshal(data)
	log.Println(string(b))
}
