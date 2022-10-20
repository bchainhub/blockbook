//go:generate sh generate.sh

package contracts

import (
	"context"
	"github.com/core-coin/go-core/accounts/abi/bind"
	"github.com/core-coin/go-core/common"
	"golang.org/x/crypto/sha3"
)

func GetChequableToken(
	ctx context.Context,
	registryAddress common.Address,
	provider bind.ContractBackend,
) (*ChequableToken, error) {
	registry, err := NewRegistry(registryAddress, provider)
	if err != nil {
		return nil, err
	}

	key := sha3.Sum256([]byte("CTN"))
	token, err := registry.Get(&bind.CallOpts{
		Context: ctx,
	}, key)
	if err != nil {
		return nil, err
	}

	tokenAddr, err := common.HexToAddress(string(token))
	if err != nil {
		return nil, err
	}

	return NewChequableToken(tokenAddr, provider)
}

func GetBountiableToken(
	ctx context.Context,
	registryAddress common.Address,
	provider bind.ContractBackend,
) (*BountiableToken, error) {
	registry, err := NewRegistry(registryAddress, provider)
	if err != nil {
		return nil, err
	}

	key := sha3.Sum256([]byte("CTN"))
	token, err := registry.Get(&bind.CallOpts{
		Context: ctx,
	}, key)
	if err != nil {
		return nil, err
	}

	tokenAddr, err := common.HexToAddress(string(token))
	if err != nil {
		return nil, err
	}

	return NewBountiableToken(tokenAddr, provider)
}
