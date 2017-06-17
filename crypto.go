package config

import (
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

func (c *Config) encryptAES(input []byte) ([]byte, error) {
	gcm, err := cipher.NewGCM(c.block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, input, nil), nil
}

func (c *Config) decryptAES(cipherText []byte) ([]byte, error) {
	gcm, err := cipher.NewGCM(c.block)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return nil, errors.New("cipherText too short")
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	return gcm.Open(nil, nonce, cipherText, nil)
}
