package aesencryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	b64 "encoding/base64"
	"errors"
	"fmt"
)

type aesEncryption struct {
	superKey []byte
}

func GetAesEncyption(superKey string) (*aesEncryption, error) {
	var aesInstance = aesEncryption{
		superKey: []byte(superKey),
	}
	return &aesInstance, nil
}

func (this *aesEncryption) Encrypt(text string) (string, error) {
	block, err := aes.NewCipher(this.superKey)
	if err != nil {
		return "", err
	}
	b := []byte(text)

	b = PKCS5Padding(b, aes.BlockSize, len(text))
	ciphertext := make([]byte, len(b))

	mode := cipher.NewCBCEncrypter(block, []byte(VECTOR))
	mode.CryptBlocks(ciphertext, b)
	encryptedText := b64.RawStdEncoding.EncodeToString(ciphertext)
	return encryptedText, nil
}

func (this *aesEncryption) Decrypt(encText string) (string, error) {

	text, err := b64.RawStdEncoding.DecodeString(encText)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(this.superKey)
	if err != nil {
		return "", err
	}

	if len(text) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	decrypted := make([]byte, len(text))
	mode := cipher.NewCBCDecrypter(block, []byte(VECTOR))
	mode.CryptBlocks(decrypted, text)
	unpaddedData, err := PKCS5UnPadding(decrypted)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", unpaddedData), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int, after int) []byte {
	padding := (blockSize - len(ciphertext)%blockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(src []byte) ([]byte, error) {
	length := len(src)
	unpadding := int(src[length-1])
	if length-unpadding < 0 {
		return nil, fmt.Errorf("Use right key to decrypt negative index error")
	}
	return src[:(length - unpadding)], nil
}
