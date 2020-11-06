package testutils

import (
	"bdsm/utils"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
)

// NewBlockchain is used to generate a simulated blockchain
func NewBlockchain() (*bind.TransactOpts, *backends.SimulatedBackend, error) {
	auth, _, err := utils.NewAccount()
	if err != nil {
		return nil, nil, err
	}
	// https://medium.com/coinmonks/unit-testing-solidity-contracts-on-ethereum-with-go-3cc924091281
	balance := new(big.Int).Mul(big.NewInt(999999999999999), big.NewInt(999999999999999))
	gAlloc := map[common.Address]core.GenesisAccount{
		auth.From: {Balance: balance},
	}
	sim := backends.NewSimulatedBackend(gAlloc, 8000000)
	return auth, sim, nil
}
