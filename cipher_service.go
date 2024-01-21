package httpsms

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

// CipherService is used to encrypt and decrypt SMS messages using the AES-256 algorithm
type CipherService service

// Encrypt the message content using the encryption key
func (service *CipherService) Encrypt(encryptionKey string, content string) (string, error) {
	block, err := aes.NewCipher(service.hash(encryptionKey))
	if err != nil {
		return "", errors.Join(err, errors.New("failed to create new cipher"))
	}

	text := []byte(content)

	iv, err := service.initializationVector()
	if err != nil {
		return "", errors.Join(err, errors.New("failed to create initialization vector"))
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	cipherText := make([]byte, len(text))
	stream.XORKeyStream(cipherText, text)

	return base64.StdEncoding.EncodeToString(append(iv, cipherText...)), nil
}

// Decrypt the message content using the encryption key
func (service *CipherService) Decrypt(encryptionKey string, cipherText string) (string, error) {
	content, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", errors.Join(err, errors.New("failed to decode cipher in base64"))
	}

	block, err := aes.NewCipher(service.hash(encryptionKey))
	if err != nil {
		return "", errors.Join(err, errors.New("failed to create new cipher"))
	}

	// Decrypt the message
	cipherTextBytes := content[16:]
	stream := cipher.NewCFBDecrypter(block, content[:16])
	stream.XORKeyStream(cipherTextBytes, cipherTextBytes)

	return string(cipherTextBytes), nil
}

// hash a key using the SHA-256 algorithm
func (service *CipherService) hash(key string) []byte {
	sha := sha256.Sum256([]byte(key))
	return sha[:]
}

func (service *CipherService) initializationVector() ([]byte, error) {
	iv := make([]byte, 16)
	_, err := rand.Read(iv)
	return iv, err
}
