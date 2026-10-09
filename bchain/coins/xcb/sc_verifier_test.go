package xcb

import (
	"context"
	"encoding/hex"
	"testing"

	xcbcommon "github.com/core-coin/go-core/v2/common"
	"github.com/core-coin/go-core/v2/core/types"
	"github.com/juju/errors"
)

// fakeRPC is one connection to the node. Once closed it fails every call, the way
// the real rpc client does.
type fakeRPC struct {
	closed bool
	reply  string
}

func (f *fakeRPC) XcbSubscribe(ctx context.Context, channel interface{}, args ...interface{}) (CVMClientSubscription, error) {
	return fakeSubscription{}, nil
}

func (f *fakeRPC) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	if f.closed {
		return errors.New("client is closed")
	}
	*result.(*string) = f.reply
	return nil
}

func (f *fakeRPC) Close() { f.closed = true }

type fakeSubscription struct{}

func (fakeSubscription) Err() <-chan error { return nil }
func (fakeSubscription) Unsubscribe()      {}

// After getBestHeader reconnects, RWA metadata must still be readable. The verifier
// used to keep the closed connection, so every RWA contract lost its metadata,
// documents and lab results at the next cache refresh and never got them back.
func TestReconnectRPCKeepsVerifierConnected(t *testing.T) {
	packed, err := stableABI.Methods["getAllKeys"].Outputs.Pack([]string{"originCountry"}, []string{"TUR"}, []bool{true})
	if err != nil {
		t.Fatal(err)
	}
	reply := "0x" + hex.EncodeToString(packed)

	old := &fakeRPC{reply: reply}
	b := &CoreblockchainRPC{
		RPC:                   old,
		ChainConfig:           &Configuration{},
		NewBlock:              &CoreCoinNewBlock{channel: make(chan *types.Header)},
		NewTx:                 &CoreCoinNewTx{channel: make(chan xcbcommon.Hash)},
		smartContractVerifier: &smartContractVerifier{RPC: old, abi: stableABI},
		OpenRPC: func(string) (CVMRPCClient, CVMClient, error) {
			return &fakeRPC{reply: reply}, nil, nil
		},
	}

	if err := b.reconnectRPC(); err != nil {
		t.Fatal(err)
	}
	if !old.closed {
		t.Fatal("reconnectRPC did not close the old connection")
	}

	metadata, err := b.smartContractVerifier.GetAllMetadata("ab69e939452343effcc1882c1b06f702210e2e03241f")
	if err != nil {
		t.Fatalf("verifier cannot read metadata after reconnect: %v", err)
	}
	if got := metadata["originCountry"]; got.Value != "TUR" || !got.Sealed {
		t.Fatalf("originCountry = %+v, want sealed TUR", got)
	}
}
