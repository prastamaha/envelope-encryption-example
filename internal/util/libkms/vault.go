package libkms

import (
	"crypto/tls"
	"log"
	"net/http"

	vault "github.com/hashicorp/vault/api"
)

type VaultKMS struct {
	Address   string
	token     string
	tlsConfig *tls.Config
}

func NewVaultKMS(address string, token string, tlsConfig *tls.Config) *VaultKMS {
	return &VaultKMS{
		Address:   address,
		token:     token,
		tlsConfig: tlsConfig,
	}
}

func (v *VaultKMS) NewClient() (*vault.Client, error) {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: v.tlsConfig,
		},
	}

	config := &vault.Config{
		Address:    v.Address,
		HttpClient: httpClient,
	}

	client, err := vault.NewClient(config)
	if err != nil {
		log.Fatalf("Unable to initialize Vault client: %v", err)
	}

	client.SetToken(v.token)

	return client, nil
}
