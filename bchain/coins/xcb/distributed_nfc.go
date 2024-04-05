package xcb

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/core-coin/go-core/v2/accounts/abi"
	"github.com/core-coin/go-core/v2/common"

	"github.com/cryptohub-digital/blockbook/bchain"
)

const definition = `[{"inputs":[{"internalType":"uint64","name":"timestamp","type":"uint64"}],"name":"accessDenied","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint64","name":"timestamp","type":"uint64"}],"name":"accessGranted","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"clear","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"showAttempts","outputs":[{"components":[{"internalType":"bool","name":"flag","type":"bool"},{"internalType":"uint64","name":"timestamp","type":"uint64"},{"internalType":"uint128","name":"id","type":"uint128"}],"internalType":"struct AttemptsLog.Attempt[]","name":"","type":"tuple[]"}],"stateMutability":"view","type":"function"}]`

type distributedNFCUseCase struct {
	RPC     CVMRPCClient
	Senders []DistributedNFCSender
	abi     abi.ABI
}

type DistributedNFCSender struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type DistributedNFCData struct {
	Sender  DistributedNFCSender
	Records []DistributedNFCAccessRecord
}

type DistributedNFCAccessRecord struct {
	Flag      bool     `json:"flag"`
	Timestamp int64    `json:"timestamp"`
	Id        *big.Int `json:"id"`
}

func newdistributedNFCUseCase(senders []DistributedNFCSender) *distributedNFCUseCase {
	abi, err := abi.JSON(strings.NewReader(definition))
	if err != nil {
		panic(err)
	}

	return &distributedNFCUseCase{Senders: senders, abi: abi}

}

func (d *distributedNFCUseCase) getAccesses(address *bchain.VerifiedAddress) []DistributedNFCData {
	res := []DistributedNFCData{}
	type scRecord struct {
		Flag      bool     `json:"flag"`
		Timestamp uint64   `json:"timestamp"`
		Id        *big.Int `json:"id"`
	}

	for _, sender := range d.Senders {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		var response string
		err := d.RPC.CallContext(ctx, &response, "xcb_call", map[string]interface{}{
			"data": "0x8ec66670",
			"from": sender.Address,
			"to":   address.Address,
		}, "latest")
		if err != nil {
			fmt.Println(err)
			return nil
		}

		responseBytes, err := common.Hex2BytesWithError(response)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		var records []scRecord
		err = d.abi.UnpackIntoInterface(&records, "showAttempts", responseBytes)
		if err != nil {
			fmt.Println(err)
			return nil
		}

		uiRecords := []DistributedNFCAccessRecord{}
		for _, scRecord := range records {
			uiRecords = append(uiRecords, DistributedNFCAccessRecord{
				Flag:      scRecord.Flag,
				Timestamp: time.Unix(0, int64(scRecord.Timestamp)*int64(time.Millisecond)).Unix(),
				Id:        scRecord.Id,
			})
		}

		res = append(res, DistributedNFCData{
			Sender:  sender,
			Records: uiRecords,
		})
	}
	return res
}
