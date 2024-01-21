package httpsms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCipherService(t *testing.T) {
	// Arrange
	key := "Password123"
	message := "This is a test text message"
	client := New()

	// Act
	encryptedMessage, encryptErr := client.Cipher.Encrypt(key, message)
	decryptedMessage, decryptErr := client.Cipher.Decrypt(key, encryptedMessage)

	// Assert
	assert.Nil(t, encryptErr)
	assert.Nil(t, decryptErr)
	assert.Equal(t, message, decryptedMessage)
}
