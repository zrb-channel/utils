package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"math/rand"
	"os"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

// RandString
// @param n
// @date 2022-09-21 22:17:57
func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// FileExists
// @param name
// @date 2022-09-21 22:17:56
func FileExists(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}

	return os.IsNotExist(err)
}

func NewPublicKey(pubKey string) (*rsa.PublicKey, error) {

	key, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return nil, err
	}

	pk, err := x509.ParsePKIXPublicKey(key)
	if err != nil {
		return nil, err
	}

	publicKey, ok := pk.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("加载publicKey失败")
	}

	return publicKey, nil
}
