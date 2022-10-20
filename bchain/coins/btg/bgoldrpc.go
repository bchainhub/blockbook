package btg

import (
	"encoding/json"
	"github.com/core-coin/go-core/xcbclient"
	"github.com/cryptohub-digital/blockbook/contracts"

	"github.com/cryptohub-digital/blockbook/bchain"
	"github.com/cryptohub-digital/blockbook/bchain/coins/btc"
	"github.com/golang/glog"
)

// BGoldRPC is an interface to JSON-RPC bitcoind service.
type BGoldRPC struct {
	*btc.BitcoinRPC
}

// NewBGoldRPC returns new BGoldRPC instance.
func NewBGoldRPC(config json.RawMessage, pushHandler func(bchain.NotificationType)) (bchain.BlockChain, error) {
	b, err := btc.NewBitcoinRPC(config, pushHandler)
	if err != nil {
		return nil, err
	}

	s := &BGoldRPC{
		b.(*btc.BitcoinRPC),
	}

	return s, nil
}

func (b *BGoldRPC) GetRPCClient() *xcbclient.Client {
	return nil
}

func (b *BGoldRPC) GetSmartContracts() (*contracts.ChequableToken, *contracts.BountiableToken) {
	return nil, nil
}

// Initialize initializes BGoldRPC instance.
func (b *BGoldRPC) Initialize() error {
	ci, err := b.GetChainInfo()
	if err != nil {
		return err
	}
	chainName := ci.Chain

	params := GetChainParams(chainName)

	// always create parser
	b.Parser = NewBGoldParser(params, b.ChainConfig)

	// parameters for getInfo request
	if params.Net == MainnetMagic {
		b.Testnet = false
		b.Network = "livenet"
	} else {
		b.Testnet = true
		b.Network = "testnet"
	}

	glog.Info("rpc: block chain ", params.Name)

	return nil
}
