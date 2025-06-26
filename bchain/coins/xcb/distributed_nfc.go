package xcb

import (
	"context"
	"fmt"
	"math/big"
	"slices"
	"strings"
	"time"

	"github.com/core-coin/go-core/v2/accounts/abi"
	"github.com/core-coin/go-core/v2/common"

	"github.com/cryptohub-digital/blockbook/bchain"
)

const definition = `[{"inputs":[{"internalType":"uint64","name":"timestamp","type":"uint64"},{"internalType":"address","name":"coreid","type":"address"}],"name":"accessDenied","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint64","name":"timestamp","type":"uint64"},{"internalType":"address","name":"coreid","type":"address"}],"name":"accessGranted","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"clear","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"showAttempts","outputs":[{"components":[{"internalType":"bool","name":"flag","type":"bool"},{"internalType":"uint64","name":"timestamp","type":"uint64"},{"internalType":"uint128","name":"id","type":"uint128"},{"internalType":"address","name":"coreid","type":"address"}],"internalType":"struct AttemptsLog.Attempt[]","name":"","type":"tuple[]"}],"stateMutability":"view","type":"function"}]`

type distributedNFCUseCase struct {
	RPC      CVMRPCClient
	abi      abi.ABI
	supabase *SupabaseClient
}

type DistributedNFCSender struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type DistributedNFCData struct {
	Sender           DistributedNFCSender
	Records          []DistributedNFCAccessRecord
	AllRecordsLength int
}

type DistributedNFCAccessRecord struct {
	Flag      bool           `json:"flag"`
	Timestamp int64          `json:"timestamp"`
	Id        *big.Int       `json:"id"`
	CoreID    common.Address `json:"coreid"`
}

func newdistributedNFCUseCase(supabase *SupabaseClient) *distributedNFCUseCase {
	abi, err := abi.JSON(strings.NewReader(definition))
	if err != nil {
		panic(err)
	}

	return &distributedNFCUseCase{supabase: supabase, abi: abi}

}

func (d *distributedNFCUseCase) getAccesses(address *bchain.VerifiedAddress, senderName string, page uint32) []DistributedNFCData {
	res := []DistributedNFCData{}
	type scRecord struct {
		Flag      bool           `json:"flag"`
		Timestamp uint64         `json:"timestamp"`
		Id        *big.Int       `json:"id"`
		CoreID    common.Address `json:"coreid"`
	}

	nfcSenders, err := d.supabase.GetDistributedNFCSenders()
	if err != nil {
		fmt.Println("ERROR: failed to get distributed NFC senders:", err)
		return nil
	}

	for _, sender := range nfcSenders {
		if senderName == "" {
			senderName = nfcSenders[0].Name
		}
		if page == 0 {
			page = 1
		}
		uiRecords := []DistributedNFCAccessRecord{}
		var records []scRecord

		if sender.Name == senderName {
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
			err = d.abi.UnpackIntoInterface(&records, "showAttempts", responseBytes)
			if err != nil {
				fmt.Println(err)
				return nil
			}

			for i := len(records) - 1 - ((int(page) - 1) * 10); i >= 0; i-- {
				uiRecords = append(uiRecords, DistributedNFCAccessRecord{
					Flag:      records[i].Flag,
					Timestamp: time.Unix(0, int64(records[i].Timestamp)*int64(time.Millisecond)).Unix(),
					Id:        records[i].Id,
					CoreID:    records[i].CoreID,
				})
				if len(uiRecords) == 10 {
					break
				}
			}

		}
		res = append(res, DistributedNFCData{
			Sender:           sender,
			Records:          uiRecords,
			AllRecordsLength: len(records),
		})
	}

	slices.Reverse(res)
	return res
}
