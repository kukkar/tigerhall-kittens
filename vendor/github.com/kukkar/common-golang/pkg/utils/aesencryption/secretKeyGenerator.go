package aesencryption

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateAesSecretKey() string {

	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower := strings.ToLower(upper)
	digits := "0123456789"
	alphnum := upper + lower + digits

	rand.Seed(time.Now().UnixNano())
	var secreKey string
	for i := 0; i < AesSecretKeyLength; i++ {
		secreKey += string(alphnum[rand.Intn(len(alphnum))])
	}
	return secreKey
}
