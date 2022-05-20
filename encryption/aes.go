package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
)

func AesEncrypt(key string, content string) (encContent string) {

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Panic(err)
	}

	// Never use more than 2^32 random nonces with a given key
	// because of the risk of repeat.
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Fatal(err)
	}

	encData := gcm.Seal(nonce, nonce, []byte(content), nil)

	return string(encData)

}

func AesDecrypt(key string, content string) (decContent string) {

	ciphertext := []byte(content)

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Panic(err)
	}

	nonce := ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Panic(err)
	}

	return string(plaintext)
}
