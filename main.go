package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type tokenInfo struct {
	Ver     int    `json:"ver"`
	Hash    string `json:"hash"`
	Nonce   string `json:"nonce"`
	Expired int64  `json:"expired"`
}

func GetToken(appid uint32, app_sign string, userID string, expired_add int64) (ret string, err error) {
	nonce := strings.ReplaceAll(uuid.New().String(), "-", "")
	nonce = "7994243b2c0648cb"
	expired := time.Now().Unix() + expired_add // 单位:秒
	expired = 1619782106

	if len(app_sign) < 32 {
		return "", fmt.Errorf("app_sign wrong")
	}

	app_sign_32 := app_sign[0:32]
	// fmt.Printf("before md5: appid:%d, app_sign_32:%s, userID%s, nonce[:16]:%s, expired:%d", appid, app_sign_32, userID, nonce[:16], expired)
	source := fmt.Sprintf("%d%s%s%s%d", appid, app_sign_32, userID, nonce[:16], expired)
	println(source)
	// sum := md5.Sum([]byte(source))
	sum := GetMD5Encode(source)
	println(string(sum[:len(sum)]))

	token := tokenInfo{}
	token.Ver = 1
	token.Hash = sum
	token.Nonce = nonce
	token.Expired = expired

	buf, err := json.Marshal(token)
	if err != nil {
		return "", err
	}

	encodeString := base64.StdEncoding.EncodeToString(buf)
	return encodeString, nil
}

func main() {
	t, _ := GetToken(2215113147, "a5f568a58a4e7b30919be8f602abb0a3", "2000003288", 1619784517)
	fmt.Printf("token:%s\n", t)
}

func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
