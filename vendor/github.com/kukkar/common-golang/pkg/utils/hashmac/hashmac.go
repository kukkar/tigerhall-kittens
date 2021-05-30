package hashmac

import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"

	"github.com/kukkar/common-golang/pkg/utils/hmacrequestgenerator"
)

func GenerateMac(data hmacrequestgenerator.HMacRequest, secret string) []byte {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))
	hashGenerated := mac.Sum(nil)
	return hashGenerated
	// d := fmt.Sprintf("mac %x", hashGenerated)
	// fmt.Printf("output %v", hmac.Equal([]byte("XrC56CoUXo8KBy7UGMMUfrEbPVvyKHw2rJfaymyW8q4="), []byte(d)))
}

func GenerateBase64Mac(data hmacrequestgenerator.HMacRequest, secret string) []byte {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))
	hashGenerated := mac.Sum(nil)
	base64Data := b64.RawStdEncoding.EncodeToString(hashGenerated)
	base64Data = base64Data + "="
	return []byte(base64Data)
}

func Equal(val1 []byte, val2 []byte) bool {
	if fmt.Sprintf("%x", val1) == fmt.Sprintf("%x", val2) {
		return true
	}
	return false
	//return hmac.Equal(val1, val2)
}
