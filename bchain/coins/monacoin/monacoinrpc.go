package monacoin

import (
	"encoding/json"
	"github.com/core-coin/go-core/xcbclient"
	"github.com/cryptohub-digital/blockbook/contracts"

	"github.com/cryptohub-digital/blockbook/bchain"
	"github.com/cryptohub-digital/blockbook/bchain/coins/btc"
	"github.com/golang/glog"
)

// MonacoinRPC is an interface to JSON-RPC bitcoind service.
type MonacoinRPC struct {
	*btc.BitcoinRPC
}

// NewMonacoinRPC returns new MonacoinRPC instance.
func NewMonacoinRPC(config json.RawMessage, pushHandler func(bchain.NotificationType)) (bchain.BlockChain, error) {
	b, err := btc.NewBitcoinRPC(config, pushHandler)
	if err != nil {
		return nil, err
	}

	s := &MonacoinRPC{
		b.(*btc.BitcoinRPC),
	}
	s.RPCMarshaler = btc.JSONMarshalerV2{}
	s.ChainConfig.SupportsEstimateFee = false

	return s, nil
}

func (b *MonacoinRPC) GetRPCClient() *xcbclient.Client {
	return nil
}

func (b *MonacoinRPC) GetSmartContracts() (*contracts.ChequableToken, *contracts.BountiableToken) {
	return nil, nil
}

// Initialize initializes MonacoinRPC instance.
func (b *MonacoinRPC) Initialize() error {
	ci, err := b.GetChainInfo()
	if err != nil {
		return err
	}
	chainName := ci.Chain

	glog.Info("Chain name ", chainName)
	params := GetChainParams(chainName)

	// always create parser
	b.Parser = NewMonacoinParser(params, b.ChainConfig)

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
