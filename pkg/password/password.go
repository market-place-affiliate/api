package password

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

func HashPassword(password string, salt []byte) string {
	block, err := aes.NewCipher(salt)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(password), nil)
	return hex.EncodeToString(ciphertext)
}

func VerifyPassword(password, saltBase64, hashedPasswordBase64 string) bool {
	pass := Decrypt([]byte(saltBase64), hashedPasswordBase64)
	return pass == password
}

func Decrypt(key []byte, encryptedHex string) string {
	ciphertext, err := hex.DecodeString(encryptedHex)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		panic(errors.New("len(ciphertext) < nonceSize"))
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}

	return string(plaintext)
}
