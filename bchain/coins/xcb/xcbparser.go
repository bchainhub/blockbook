package xcb

import (
	"encoding/hex"
	"math/big"
	"strconv"

	"github.com/core-coin/go-core/common/hexutil"
	"github.com/golang/protobuf/proto"
	"github.com/juju/errors"
	"github.com/trezor/blockbook/bchain"
)

// CoreblockchainTypeAddressDescriptorLen - in case of EthereumType, the AddressDescriptor has fixed length
const CoreblockchainTypeAddressDescriptorLen = 22

// XcbAmountDecimalPoint defines number of decimal points in Ether amounts
const XcbAmountDecimalPoint = 18

// CoreblockchainParser handle
type CoreblockchainParser struct {
	*bchain.BaseParser
}

// NewCoreblockchainParser returns new CoreblockchainParser instance
func NewCoreblockchainParser(b int) *CoreblockchainParser {
	return &CoreblockchainParser{&bchain.BaseParser{
		BlockAddressesToKeep: b,
		AmountDecimalPoint:   XcbAmountDecimalPoint,
	}}
}

type rpcHeader struct {
	Hash       string `json:"hash"`
	ParentHash string `json:"parentHash"`
	Difficulty string `json:"difficulty"`
	Number     string `json:"number"`
	Time       string `json:"timestamp"`
	Size       string `json:"size"`
	Nonce      string `json:"nonce"`
}

type rpcTransaction struct {
	AccountNonce     string `json:"nonce"`
	EnergyPrice      string `json:"energyPrice"`
	EnergyLimit      string `json:"energy"`
	To               string `json:"to"` // nil means contract creation
	Value            string `json:"value"`
	Payload          string `json:"input"`
	Hash             string `json:"hash"`
	BlockNumber      string `json:"blockNumber"`
	BlockHash        string `json:"blockHash,omitempty"`
	From             string `json:"from"`
	TransactionIndex string `json:"transactionIndex"`
	// Signature values - ignored
	// V string `json:"v"`
	// R string `json:"r"`
	// S string `json:"s"`
}

type rpcLog struct {
	Address string   `json:"address"`
	Topics  []string `json:"topics"`
	Data    string   `json:"data"`
}

type rpcLogWithTxHash struct {
	rpcLog
	Hash string `json:"transactionHash"`
}

type rpcReceipt struct {
	EnergyUsed string    `json:"energyUsed"`
	Status     string    `json:"status"`
	Logs       []*rpcLog `json:"logs"`
}

type completeTransaction struct {
	Tx      *rpcTransaction `json:"tx"`
	Receipt *rpcReceipt     `json:"receipt,omitempty"`
}

type rpcBlockTransactions struct {
	Transactions []rpcTransaction `json:"transactions"`
}

type rpcBlockTxids struct {
	Transactions []string `json:"transactions"`
}

func xcbNumber(n string) (int64, error) {
	if len(n) > 2 {
		return strconv.ParseInt(n[2:], 16, 64)
	}
	return 0, errors.Errorf("Not a number: '%v'", n)
}

func (p *CoreblockchainParser) xcbTxToTx(tx *rpcTransaction, receipt *rpcReceipt, blocktime int64, confirmations uint32, fixEIP55 bool) (*bchain.Tx, error) {
	txid := tx.Hash
	var (
		fa, ta []string
		err    error
	)
	if len(tx.From) > 2 {
		fa = []string{tx.From}
	}
	if len(tx.To) > 2 {
		ta = []string{tx.To}
	}

	ct := completeTransaction{
		Tx:      tx,
		Receipt: receipt,
	}
	vs, err := hexutil.DecodeBig(tx.Value)
	if err != nil {
		return nil, err
	}
	return &bchain.Tx{
		Blocktime:     blocktime,
		Confirmations: confirmations,
		// Hex
		// LockTime
		Time: blocktime,
		Txid: txid,
		Vin: []bchain.Vin{
			{
				Addresses: fa,
				// Coinbase
				// ScriptSig
				// Sequence
				// Txid
				// Vout
			},
		},
		Vout: []bchain.Vout{
			{
				N:        0, // there is always up to one To address
				ValueSat: *vs,
				ScriptPubKey: bchain.ScriptPubKey{
					// Hex
					Addresses: ta,
				},
			},
		},
		CoinSpecificData: ct,
	}, nil
}

// GetAddrDescFromVout returns internal address representation of given transaction output
func (p *CoreblockchainParser) GetAddrDescFromVout(output *bchain.Vout) (bchain.AddressDescriptor, error) {
	if len(output.ScriptPubKey.Addresses) != 1 {
		return nil, bchain.ErrAddressMissing
	}
	return p.GetAddrDescFromAddress(output.ScriptPubKey.Addresses[0])
}

func has0xPrefix(s string) bool {
	return len(s) >= 2 && s[0] == '0' && (s[1]|32) == 'x'
}

// GetAddrDescFromAddress returns internal address representation of given address
func (p *CoreblockchainParser) GetAddrDescFromAddress(address string) (bchain.AddressDescriptor, error) {
	// github.com/core-coin/go-core/common.HexToAddress does not handle address errors, using own decoding
	if has0xPrefix(address) {
		address = address[2:]
	}
	if len(address) != CoreblockchainTypeAddressDescriptorLen*2 {
		return nil, bchain.ErrAddressMissing
	}
	return hex.DecodeString(address)
}

// EIP55Address returns an EIP55-compliant hex string representation of the address
func EIP55Address(addrDesc bchain.AddressDescriptor) string {
	raw := hexutil.Encode(addrDesc)

	if len(raw) > 2 {
		raw = raw[2:]
	}

	return raw
}

// EIP55AddressFromAddress returns an EIP55-compliant hex string representation of the address
func EIP55AddressFromAddress(address string) string {
	if has0xPrefix(address) {
		address = address[2:]
	}
	b, err := hex.DecodeString(address)
	if err != nil {
		return address
	}
	return EIP55Address(b)
}

// GetAddressesFromAddrDesc returns addresses for given address descriptor with flag if the addresses are searchable
func (p *CoreblockchainParser) GetAddressesFromAddrDesc(addrDesc bchain.AddressDescriptor) ([]string, bool, error) {
	return []string{EIP55Address(addrDesc)}, true, nil
}

// GetScriptFromAddrDesc returns output script for given address descriptor
func (p *CoreblockchainParser) GetScriptFromAddrDesc(addrDesc bchain.AddressDescriptor) ([]byte, error) {
	return addrDesc, nil
}

func hexDecode(s string) ([]byte, error) {
	b, err := hexutil.Decode(s)
	if err != nil && err != hexutil.ErrEmptyString {
		return nil, err
	}
	return b, nil
}

func hexDecodeBig(s string) ([]byte, error) {
	b, err := hexutil.DecodeBig(s)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func hexEncodeBig(b []byte) string {
	var i big.Int
	i.SetBytes(b)
	return hexutil.EncodeBig(&i)
}

// PackTx packs transaction to byte array
func (p *CoreblockchainParser) PackTx(tx *bchain.Tx, height uint32, blockTime int64) ([]byte, error) {
	var err error
	var n uint64
	r, ok := tx.CoinSpecificData.(completeTransaction)
	if !ok {
		return nil, errors.New("Missing CoinSpecificData")
	}
	pt := &ProtoCompleteTransaction{}
	pt.Tx = &ProtoCompleteTransaction_TxType{}
	if pt.Tx.AccountNonce, err = hexutil.DecodeUint64(r.Tx.AccountNonce); err != nil {
		return nil, errors.Annotatef(err, "AccountNonce %v", r.Tx.AccountNonce)
	}
	// pt.BlockNumber = height
	if n, err = hexutil.DecodeUint64(r.Tx.BlockNumber); err != nil {
		return nil, errors.Annotatef(err, "BlockNumber %v", r.Tx.BlockNumber)
	}
	pt.BlockNumber = uint32(n)
	pt.BlockTime = uint64(blockTime)
	if pt.Tx.From, err = hexDecode(r.Tx.From); err != nil {
		return nil, errors.Annotatef(err, "From %v", r.Tx.From)
	}
	if pt.Tx.EnergyLimit, err = hexutil.DecodeUint64(r.Tx.EnergyLimit); err != nil {
		return nil, errors.Annotatef(err, "EnergyLimit %v", r.Tx.EnergyLimit)
	}
	if pt.Tx.Hash, err = hexDecode(r.Tx.Hash); err != nil {
		return nil, errors.Annotatef(err, "Hash %v", r.Tx.Hash)
	}
	if pt.Tx.Payload, err = hexDecode(r.Tx.Payload); err != nil {
		return nil, errors.Annotatef(err, "Payload %v", r.Tx.Payload)
	}
	if pt.Tx.EnergyPrice, err = hexDecodeBig(r.Tx.EnergyPrice); err != nil {
		return nil, errors.Annotatef(err, "Price %v", r.Tx.EnergyPrice)
	}
	// if pt.R, err = hexDecodeBig(r.R); err != nil {
	// 	return nil, errors.Annotatef(err, "R %v", r.R)
	// }
	// if pt.S, err = hexDecodeBig(r.S); err != nil {
	// 	return nil, errors.Annotatef(err, "S %v", r.S)
	// }
	// if pt.V, err = hexDecodeBig(r.V); err != nil {
	// 	return nil, errors.Annotatef(err, "V %v", r.V)
	// }
	if pt.Tx.To, err = hexDecode(r.Tx.To); err != nil {
		return nil, errors.Annotatef(err, "To %v", r.Tx.To)
	}
	if n, err = hexutil.DecodeUint64(r.Tx.TransactionIndex); err != nil {
		return nil, errors.Annotatef(err, "TransactionIndex %v", r.Tx.TransactionIndex)
	}
	pt.Tx.TransactionIndex = uint32(n)
	if pt.Tx.Value, err = hexDecodeBig(r.Tx.Value); err != nil {
		return nil, errors.Annotatef(err, "Value %v", r.Tx.Value)
	}
	if r.Receipt != nil {
		pt.Receipt = &ProtoCompleteTransaction_ReceiptType{}
		if pt.Receipt.EnergyUsed, err = hexDecodeBig(r.Receipt.EnergyUsed); err != nil {
			return nil, errors.Annotatef(err, "EnergyUsed %v", r.Receipt.EnergyUsed)
		}
		if r.Receipt.Status != "" {
			if pt.Receipt.Status, err = hexDecodeBig(r.Receipt.Status); err != nil {
				return nil, errors.Annotatef(err, "Status %v", r.Receipt.Status)
			}
		} else {
			// unknown status, use 'U' as status bytes
			// there is a potential for conflict with value 0x55 but this is not used by any chain at this moment
			pt.Receipt.Status = []byte{'U'}
		}
		ptLogs := make([]*ProtoCompleteTransaction_ReceiptType_LogType, len(r.Receipt.Logs))
		for i, l := range r.Receipt.Logs {
			a, err := hexutil.Decode(l.Address)
			if err != nil {
				return nil, errors.Annotatef(err, "Address cannot be decoded %v", l)
			}
			d, err := hexutil.Decode(l.Data)
			if err != nil {
				return nil, errors.Annotatef(err, "Data cannot be decoded %v", l)
			}
			t := make([][]byte, len(l.Topics))
			for j, s := range l.Topics {
				t[j], err = hexutil.Decode(s)
				if err != nil {
					return nil, errors.Annotatef(err, "Topic cannot be decoded %v", l)
				}
			}
			ptLogs[i] = &ProtoCompleteTransaction_ReceiptType_LogType{
				Address: a,
				Data:    d,
				Topics:  t,
			}

		}
		pt.Receipt.Log = ptLogs
	}
	return proto.Marshal(pt)
}

// UnpackTx unpacks transaction from byte array
func (p *CoreblockchainParser) UnpackTx(buf []byte) (*bchain.Tx, uint32, error) {
	var pt ProtoCompleteTransaction
	err := proto.Unmarshal(buf, &pt)
	if err != nil {
		return nil, 0, err
	}
	rt := rpcTransaction{
		AccountNonce: hexutil.EncodeUint64(pt.Tx.AccountNonce),
		BlockNumber:  hexutil.EncodeUint64(uint64(pt.BlockNumber)),
		From:         EIP55Address(pt.Tx.From),
		EnergyLimit:  hexutil.EncodeUint64(pt.Tx.EnergyLimit),
		Hash:         hexutil.Encode(pt.Tx.Hash),
		Payload:      hexutil.Encode(pt.Tx.Payload),
		EnergyPrice:  hexEncodeBig(pt.Tx.EnergyPrice),
		// R:                hexEncodeBig(pt.R),
		// S:                hexEncodeBig(pt.S),
		// V:                hexEncodeBig(pt.V),
		To:               EIP55Address(pt.Tx.To),
		TransactionIndex: hexutil.EncodeUint64(uint64(pt.Tx.TransactionIndex)),
		Value:            hexEncodeBig(pt.Tx.Value),
	}
	var rr *rpcReceipt
	if pt.Receipt != nil {
		logs := make([]*rpcLog, len(pt.Receipt.Log))
		for i, l := range pt.Receipt.Log {
			topics := make([]string, len(l.Topics))
			for j, t := range l.Topics {
				topics[j] = hexutil.Encode(t)
			}
			logs[i] = &rpcLog{
				Address: EIP55Address(l.Address),
				Data:    hexutil.Encode(l.Data),
				Topics:  topics,
			}
		}
		status := ""
		// handle a special value []byte{'U'} as unknown state
		if len(pt.Receipt.Status) != 1 || pt.Receipt.Status[0] != 'U' {
			status = hexEncodeBig(pt.Receipt.Status)
		}
		rr = &rpcReceipt{
			EnergyUsed: hexEncodeBig(pt.Receipt.EnergyUsed),
			Status:     status,
			Logs:       logs,
		}
	}
	tx, err := p.xcbTxToTx(&rt, rr, int64(pt.BlockTime), 0, false)
	if err != nil {
		return nil, 0, err
	}
	return tx, pt.BlockNumber, nil
}

// PackedTxidLen returns length in bytes of packed txid
func (p *CoreblockchainParser) PackedTxidLen() int {
	return 32
}

// PackTxid packs txid to byte array
func (p *CoreblockchainParser) PackTxid(txid string) ([]byte, error) {
	if has0xPrefix(txid) {
		txid = txid[2:]
	}
	return hex.DecodeString(txid)
}

// UnpackTxid unpacks byte array to txid
func (p *CoreblockchainParser) UnpackTxid(buf []byte) (string, error) {
	return hexutil.Encode(buf), nil
}

// PackBlockHash packs block hash to byte array
func (p *CoreblockchainParser) PackBlockHash(hash string) ([]byte, error) {
	if has0xPrefix(hash) {
		hash = hash[2:]
	}
	return hex.DecodeString(hash)
}

// UnpackBlockHash unpacks byte array to block hash
func (p *CoreblockchainParser) UnpackBlockHash(buf []byte) (string, error) {
	return hexutil.Encode(buf), nil
}

// GetChainType returns EthereumType
func (p *CoreblockchainParser) GetChainType() bchain.ChainType {
	return bchain.ChainEthereumType
}

// GetHeightFromTx returns coreblockchain specific data from bchain.Tx
func GetHeightFromTx(tx *bchain.Tx) (uint32, error) {
	var bn string
	csd, ok := tx.CoinSpecificData.(completeTransaction)
	if !ok {
		return 0, errors.New("Missing CoinSpecificData")
	}
	bn = csd.Tx.BlockNumber
	n, err := hexutil.DecodeUint64(bn)
	if err != nil {
		return 0, errors.Annotatef(err, "BlockNumber %v", bn)
	}
	return uint32(n), nil
}

// CoreblockchainTypeGetxrc20FromTx returns xrc20 data from bchain.Tx
func (p *CoreblockchainParser) CoreblockchainTypeGetxrc20FromTx(tx *bchain.Tx) ([]bchain.Erc20Transfer, error) {
	var r []bchain.Erc20Transfer
	var err error
	csd, ok := tx.CoinSpecificData.(completeTransaction)
	if ok {
		if csd.Receipt != nil {
			r, err = xrc20GetTransfersFromLog(csd.Receipt.Logs)
		} else {
			r, err = xrc20GetTransfersFromTx(csd.Tx)
		}
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

// TxStatus is status of transaction
type TxStatus int

// statuses of transaction
const (
	TxStatusUnknown = TxStatus(iota - 2)
	TxStatusPending
	TxStatusFailure
	TxStatusOK
)

// CoreblockchainTxData contains coreblockchain specific transaction data
type CoreblockchainTxData struct {
	Status      TxStatus `json:"status"` // 1 OK, 0 Fail, -1 pending, -2 unknown
	Nonce       uint64   `json:"nonce"`
	EnergyLimit *big.Int `json:"energylimit"`
	EnergyUsed  *big.Int `json:"energyused"`
	EnergyPrice *big.Int `json:"energyprice"`
	Data        string   `json:"data"`
}

// GetCoreblockchainTxData returns CoreblockchainTxData from bchain.Tx
func GetCoreblockchainTxData(tx *bchain.Tx) *CoreblockchainTxData {
	return GetCoreblockchainTxDataFromSpecificData(tx.CoinSpecificData)
}

// GetCoreblockchainTxDataFromSpecificData returns CoreblockchainTxData from coinSpecificData
func GetCoreblockchainTxDataFromSpecificData(coinSpecificData interface{}) *CoreblockchainTxData {
	etd := CoreblockchainTxData{Status: TxStatusPending}
	csd, ok := coinSpecificData.(completeTransaction)
	if ok {
		if csd.Tx != nil {
			etd.Nonce, _ = hexutil.DecodeUint64(csd.Tx.AccountNonce)
			etd.EnergyLimit, _ = hexutil.DecodeBig(csd.Tx.EnergyLimit)
			etd.EnergyPrice, _ = hexutil.DecodeBig(csd.Tx.EnergyPrice)
			etd.Data = csd.Tx.Payload
		}
		if csd.Receipt != nil {
			switch csd.Receipt.Status {
			case "0x1":
				etd.Status = TxStatusOK
			case "": // old transactions did not set status
				etd.Status = TxStatusUnknown
			default:
				etd.Status = TxStatusFailure
			}
			etd.EnergyUsed, _ = hexutil.DecodeBig(csd.Receipt.EnergyUsed)
		}
	}
	return &etd
}
