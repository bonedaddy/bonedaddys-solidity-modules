package utils

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
)

// NewAccount creates a new ethereum private key, and associated transactor
func NewAccount() (*bind.TransactOpts, *ecdsa.PrivateKey, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, nil, err
	}
	return bind.NewKeyedTransactor(key), key, err
}
