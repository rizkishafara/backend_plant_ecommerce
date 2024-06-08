package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var Secretkey = []byte("@9r33n3l394nt123")

func GetFileTypeFromBase64Header(header string) (string, error) {
	switch {
	case strings.Contains(header, "image/jpeg"):
		return "jpeg", nil
	case strings.Contains(header, "image/jpg"):
		return "jpg", nil
	case strings.Contains(header, "image/png"):
		return "png", nil
	default:
		return "", errors.New("unsupported file type")
	}
}

func DeleteFile(file, jenis string) error {
	path := fmt.Sprintf("%s/%s", jenis, file)
	err := os.Remove("/uploads/" + path)
	if err != nil {
		return err
	}
	return nil
}

func Encrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher(Secretkey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}
func Decrypt(encrypted string) (string, error) {
	ciphertext, err := base64.URLEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(Secretkey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
