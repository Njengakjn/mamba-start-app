package handlers

import "github.com/kjnjenga/mamba"

func (h *Handlers) encrypt(text string) (string, error) {
	enc := mamba.Encryption{Key: []byte(h.mamba.EncryptionKey)}

	encrypted, err := enc.Encrypt(text)
	if err != nil {
		return "", err
	}

	return encrypted, nil
}

func (h *Handlers) decrypt(crypto string) (string, error) {
	enc := mamba.Encryption{Key: []byte(h.mamba.EncryptionKey)}

	decrypted, err := enc.Decrypt(crypto)
	if err != nil {
		return "", err
	}

	return decrypted, nil
}

func randomString(h *Handlers, n int) string {
	return h.mamba.RandomString(n)
}
