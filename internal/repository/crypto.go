package repository

import (
	"encoding/base64"
	"fmt"

	vault "github.com/hashicorp/vault/api"
	"github.com/prastamaha/envelope-encryption-example/internal/model"
)

type Crypto struct {
	kmsClient *vault.Client
	kekName   string
}

func NewCrypto(kmsClient *vault.Client, kekName string) *Crypto {
	return &Crypto{
		kmsClient: kmsClient,
		kekName:   kekName,
	}
}

func (c *Crypto) GenerateDEK() (*model.DEK, error) {
	path := fmt.Sprintf("transit/datakey/plaintext/%s", c.kekName)
	secret, err := c.kmsClient.Logical().Write(path, map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	return &model.DEK{
		Plaintext:  secret.Data["plaintext"].(string),
		Ciphertext: secret.Data["ciphertext"].(string),
	}, nil
}

func (c *Crypto) DecryptDEK(ciphertext []byte) ([]byte, error) {
	path := fmt.Sprintf("transit/decrypt/%s", c.kekName)
	secret, err := c.kmsClient.Logical().Write(path, map[string]interface{}{
		"ciphertext": string(ciphertext),
	})
	if err != nil {
		return nil, err
	}

	decodedDek, err := base64.StdEncoding.DecodeString(secret.Data["plaintext"].(string))
	if err != nil {
		return nil, err
	}

	return decodedDek, nil
}

func (c *Crypto) Encrypt(dek []byte, plaintext []byte) ([]byte, error) {
	path := fmt.Sprintf("transit/encrypt/%s", c.kekName)
	secret, err := c.kmsClient.Logical().Write(path, map[string]interface{}{
		"plaintext": base64.StdEncoding.EncodeToString(plaintext),
	})
	if err != nil {
		return nil, err
	}

	return []byte(secret.Data["ciphertext"].(string)), nil
}

func (c *Crypto) Decrypt(dek []byte, ciphertext []byte) ([]byte, error) {
	path := fmt.Sprintf("transit/decrypt/%s", c.kekName)
	secret, err := c.kmsClient.Logical().Write(path, map[string]interface{}{
		"ciphertext": string(ciphertext),
	})
	if err != nil {
		return nil, err
	}

	decodedData, err := base64.StdEncoding.DecodeString(secret.Data["plaintext"].(string))
	if err != nil {
		return nil, err
	}

	return decodedData, nil
}
