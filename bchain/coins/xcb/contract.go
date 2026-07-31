package xcb

import (
	"bytes"
	"context"
	"encoding/hex"
	"math/big"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/core-coin/go-core/v2/common"
	"github.com/core-coin/go-core/v2/common/hexutil"
	"github.com/golang/glog"
	"github.com/juju/errors"

	"github.com/cryptohub-digital/blockbook/bchain"
)

const tokenTransferEventSignature = "0xc17a9d92b89f27cb79cc390f23a1a5d302fefab8c7911075ede952ac2b5607a1"

// doing the parsing/processing without using go-core/accounts/abi library, it is simple to get data from Transfer event
const cbc20TransferMethodSignature = "0x4b40e901"

const nameSignature = "0x07ba2a17"
const symbolSignature = "0x231782d8"
const decimalsSignature = "0x5d1fb5f9"
const totalSupplySignature = "0x1f1881f8"
const balanceOfSignature = "0x1d7976f3"
const supportsInterfaceSignature = "0x80ada41b"
const erc721InterfaceID = "0b911da1"

const cbc721TransferFromMethodSignature = "0x31f2e679"             // transferFrom(address,address,uint256)
const cbc721SafeTransferFromMethodSignature = "0x3453ba4a"         // safeTransferFrom(address,address,uint256)
const cbc721SafeTransferFromWithDataMethodSignature = "0xf3d63809" // safeTransferFrom(address,address,uint256,bytes)

var cachedContracts = make(map[string]*bchain.ContractInfo)
var cachedContractsMux sync.Mutex
var cachedContractsTimestamps = make(map[string]time.Time)

const contractCacheTTL = 3 * time.Minute

type cachedSupply struct {
	supply    *big.Int
	timestamp time.Time
}

var cachedSupplies = make(map[string]cachedSupply)
var cachedSuppliesMux sync.Mutex

const supplyCacheTTL = 2 * time.Minute

func addressFromPaddedHex(s string) (string, error) {
	var t big.Int
	var ok bool
	if has0xPrefix(s) {
		_, ok = t.SetString(s[2:], 16)
	} else {
		_, ok = t.SetString(s, 16)
	}
	if !ok {
		return "", errors.New("Data is not a number")
	}
	a := common.BigToAddress(&t)
	return a.String(), nil
}

func getTokenTransfersFromLog(logs []*RpcLog) (bchain.TokenTransfers, error) {
	var r bchain.TokenTransfers
	var tt *bchain.TokenTransfer
	var err error
	for _, l := range logs {
		tl := len(l.Topics)
		if tl > 0 {
			signature := l.Topics[0]
			if signature == tokenTransferEventSignature {
				tt, err = processtokenTransferEventFromLogs(l)
			} else {
				continue
			}
			if err != nil {
				return nil, err
			}
			if tt != nil {
				r = append(r, tt)
			}
		}
	}
	return r, nil
}

func processtokenTransferEventFromLogs(log *RpcLog) (*bchain.TokenTransfer, error) {
	tl := len(log.Topics)
	var ttt bchain.TokenType
	var value big.Int
	if tl == 3 {
		ttt = bchain.FungibleToken
		_, ok := value.SetString(log.Data, 0)
		if !ok {
			return nil, errors.New("CBC20 log Data is not a number")
		}
	} else if tl == 4 {
		ttt = bchain.NonFungibleToken
		_, ok := value.SetString(log.Topics[3], 0)
		if !ok {
			return nil, errors.New("CBC721 log Topics[3] is not a number")
		}
	} else {
		return nil, nil
	}

	from, err := addressFromPaddedHex(log.Topics[1])
	if err != nil {
		return nil, err
	}
	to, err := addressFromPaddedHex(log.Topics[2])
	if err != nil {
		return nil, err
	}
	return &bchain.TokenTransfer{
		Type:     ttt,
		Contract: log.Address,
		From:     from,
		To:       to,
		Value:    value,
	}, nil
}

func getTokenTransfersFromTx(tx *RpcTransaction) (bchain.TokenTransfers, error) {
	var r bchain.TokenTransfers
	if len(tx.Payload)%(128+len(cbc20TransferMethodSignature)) == 0 && strings.HasPrefix(tx.Payload, cbc20TransferMethodSignature) {
		to, err := addressFromPaddedHex(tx.Payload[len(cbc20TransferMethodSignature) : 64+len(cbc20TransferMethodSignature)])
		if err != nil {
			return nil, err
		}
		var t big.Int
		_, ok := t.SetString(tx.Payload[len(cbc20TransferMethodSignature)+64:], 16)
		if !ok {
			return nil, errors.New("Data is not a number")
		}
		r = append(r, &bchain.TokenTransfer{
			Contract: tx.To,
			From:     tx.From,
			To:       to,
			Value:    t,
			Type:     bchain.FungibleToken,
		})
	} else if len(tx.Payload) >= 10+192 &&
		(strings.HasPrefix(tx.Payload, cbc721TransferFromMethodSignature) ||
			strings.HasPrefix(tx.Payload, cbc721SafeTransferFromMethodSignature) ||
			strings.HasPrefix(tx.Payload, cbc721SafeTransferFromWithDataMethodSignature)) {
		from, err := addressFromPaddedHex(tx.Payload[10 : 10+64])
		if err != nil {
			return nil, err
		}
		to, err := addressFromPaddedHex(tx.Payload[10+64 : 10+128])
		if err != nil {
			return nil, err
		}
		var t big.Int
		_, ok := t.SetString(tx.Payload[10+128:10+192], 16)
		if !ok {
			return nil, errors.New("Data is not a number")
		}
		r = append(r, &bchain.TokenTransfer{
			Type:     bchain.NonFungibleToken,
			Contract: tx.To,
			From:     from,
			To:       to,
			Value:    t,
		})
	}
	return r, nil
}

func (b *CoreblockchainRPC) contractSupportsInterface(contractDesc bchain.AddressDescriptor, address string, interfaceID string) (bool, error) {
	if len(interfaceID) != 8 {
		return false, errors.New("invalid interface id length")
	}
	callData := supportsInterfaceSignature + interfaceID + strings.Repeat("0", 64-len(interfaceID))
	data, err := b.xcbCall(callData, address)
	if err != nil {
		if strings.Contains(err.Error(), "execution reverted") {
			return false, nil
		}
		return false, err
	}
	result := parseCBC20NumericProperty(contractDesc, data)
	if result == nil {
		return false, nil
	}
	return result.Sign() != 0, nil
}

func (b *CoreblockchainRPC) xcbCall(data, to string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), b.Timeout)
	defer cancel()
	var r string
	err := b.RPC.CallContext(ctx, &r, "xcb_call", map[string]interface{}{
		"data": data,
		"to":   to,
	}, "latest")
	if err != nil {
		return "", err
	}
	return r, nil
}

func parseCBC20NumericProperty(contractDesc bchain.AddressDescriptor, data string) *big.Int {
	if has0xPrefix(data) {
		data = data[2:]
	}
	if len(data) > 64 {
		data = data[:64]
	}
	if len(data) == 64 {
		var n big.Int
		_, ok := n.SetString(data, 16)
		if ok {
			return &n
		}
	}
	if glog.V(1) {
		glog.Warning("Cannot parse '", data, "' for contract ", contractDesc)
	}
	return nil
}

func parseCBC20StringProperty(contractDesc bchain.AddressDescriptor, data string) string {
	if has0xPrefix(data) {
		data = data[2:]
	}
	if len(data) > 128 {
		n := parseCBC20NumericProperty(contractDesc, data[64:128])
		if n != nil {
			l := n.Uint64()
			if l > 0 && 2*int(l) <= len(data)-128 {
				b, err := hex.DecodeString(data[128 : 128+2*l])
				if err == nil {
					return string(b)
				}
			}
		}
	}
	// allow string properties as UTF-8 data
	b, err := hex.DecodeString(data)
	if err == nil {
		i := bytes.Index(b, []byte{0})
		if i > 32 {
			i = 32
		}
		if i > 0 {
			b = b[:i]
		}
		if utf8.Valid(b) {
			return string(b)
		}
	}
	if glog.V(1) {
		glog.Warning("Cannot parse '", data, "' for contract ", contractDesc)
	}
	return ""
}

// getTokenSupply calls totalSupply() on a contract and returns the result in base units.
// It returns nil if the contract does not implement totalSupply() or the call failed.
// Values are cached for supplyCacheTTL so that page views do not hit the backend every time.
func (b *CoreblockchainRPC) getTokenSupply(contractDesc bchain.AddressDescriptor, address string) *big.Int {
	key := common.Bytes2Hex(contractDesc)
	now := time.Now()

	cachedSuppliesMux.Lock()
	cached, found := cachedSupplies[key]
	cachedSuppliesMux.Unlock()
	if found && now.Sub(cached.timestamp) <= supplyCacheTTL {
		return cached.supply
	}

	var supply *big.Int
	data, err := b.xcbCall(totalSupplySignature, address)
	if err != nil {
		// a contract without totalSupply() reverts, that is not an error worth logging
		if !strings.Contains(err.Error(), "execution reverted") {
			glog.Warning(errors.Annotatef(err, "cbc20TotalSupplySignature %v", address))
		}
	} else {
		supply = parseCBC20NumericProperty(contractDesc, data)
	}

	cachedSuppliesMux.Lock()
	cachedSupplies[key] = cachedSupply{supply: supply, timestamp: now}
	cachedSuppliesMux.Unlock()
	return supply
}

// groupDigits inserts thousands separators into a decimal integer, e.g. "1027551" ->
// "1,027,551". message.Printer cannot do this, it does not format big.Int values.
func groupDigits(value string) string {
	sign := ""
	if strings.HasPrefix(value, "-") {
		sign, value = "-", value[1:]
	}
	if len(value) <= 3 {
		return sign + value
	}
	var grouped strings.Builder
	grouped.WriteString(sign)
	first := len(value) % 3
	if first == 0 {
		first = 3
	}
	grouped.WriteString(value[:first])
	for i := first; i < len(value); i += 3 {
		grouped.WriteByte(',')
		grouped.WriteString(value[i : i+3])
	}
	return grouped.String()
}

// formatTokenSupply renders a supply given in base units as a human readable amount,
// e.g. 1027551000 with 6 decimals becomes "1,027.551". A nil supply renders as an empty
// string so that the caller can tell "unknown" apart from a genuine zero supply.
func formatTokenSupply(supply *big.Int, decimals int) string {
	if supply == nil {
		return ""
	}
	if decimals <= 0 {
		return groupDigits(supply.String())
	}
	scale := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	whole, fraction := new(big.Int).QuoRem(supply, scale, new(big.Int))
	formatted := groupDigits(whole.String())
	if fraction.Sign() == 0 {
		return formatted
	}
	if whole.Sign() == 0 && fraction.Sign() < 0 {
		// QuoRem keeps the sign on the remainder, the integer part lost it by truncating to zero
		formatted = "-" + formatted
	}
	// left pad the fraction to the full number of decimals, then drop trailing zeroes
	digits := new(big.Int).Abs(fraction).String()
	if len(digits) < decimals {
		digits = strings.Repeat("0", decimals-len(digits)) + digits
	}
	return formatted + "." + strings.TrimRight(digits, "0")
}

// resolveSupply decides which supply to report for a token. The registry and the contract
// both state supply in base units; a supply published by the registry is authoritative and
// overrides the contract, which is only consulted (via fromChain) when the registry is
// silent. Circulating supply equals total supply unless the registry states it separately.
// A nil result means the supply is unknown, which is not the same as a supply of zero.
func resolveSupply(sc *VerifiedSC, fromChain func() *big.Int) (total *big.Int, circulating *big.Int) {
	if sc != nil {
		total = sc.TotalSupply
		circulating = sc.CirculatingSupply
	}
	if total == nil {
		total = fromChain()
	}
	if circulating == nil {
		circulating = total
	}
	return total, circulating
}

// addSupplyData fills in the supply of a fungible token, formatted for display using the
// token's own decimals.
func (b *CoreblockchainRPC) addSupplyData(contract *bchain.ContractInfo, sc *VerifiedSC) {
	if contract.Type != CBC20TokenType {
		return
	}
	total, circulating := resolveSupply(sc, func() *big.Int {
		contractDesc, err := b.Parser.GetAddrDescFromAddress(contract.Contract)
		if err != nil {
			glog.Warning(errors.Annotatef(err, "addSupplyData %v", contract.Contract))
			return nil
		}
		return b.getTokenSupply(contractDesc, contract.Contract)
	})
	contract.TotalSupply = formatTokenSupply(total, contract.Decimals)
	contract.CirculatingSupply = formatTokenSupply(circulating, contract.Decimals)
}

func (b *CoreblockchainRPC) AddVerifiedSCData(contract *bchain.ContractInfo) *bchain.ContractInfo {
	if contract != nil {
		// if smart contract ticker is verified but address is wrong -> do not show SC symbol (ticker)
		if !b.smartContractVerifier.IsValidVerifiedSC(contract.Contract, contract.Symbol) {
			contract.Symbol = ""
			return contract
		}
		sc := b.smartContractVerifier.GetVerified(contract.Contract)
		b.addSupplyData(contract, sc)
		// if smart contract address is verified -> add verifying data
		if sc != nil {
			contract.Icon = sc.Icon
			contract.VerifierWebAddress = sc.Web
			contract.Symbol = sc.Ticker

			// RWA
			if sc.Metadata != nil {
				converted := make(bchain.ContractMetadata)
				for k, v := range sc.Metadata {
					converted[k] = bchain.Metadata{
						Value:  v.Value,
						Sealed: v.Sealed,
					}
				}
				contract.Metadata = converted
			}
			if sc.KnownMetadata != nil {
				converted := make(bchain.ContractMetadata)
				for k, v := range sc.KnownMetadata {
					converted[k] = bchain.Metadata{
						Value:  v.Value,
						Sealed: v.Sealed,
					}
				}
				contract.KnownMetadata = converted
			}
			if sc.Documents != nil {
				contract.Documents = sc.Documents
			}
			if sc.LabResults != nil {
				contract.LabResults = sc.LabResults
			}
		}
	}
	return contract
}

func (b *CoreblockchainRPC) FindVerifiedByName(query string) *bchain.AddressDescriptor {
	contains := func(s []string, e string) bool {
		for _, a := range s {
			if strings.EqualFold(a, e) {
				return true
			}
		}
		return false
	}
	for _, sc := range b.smartContractVerifier.GetAllSmartContracts() {
		if contains(sc.Aliases, query) {
			ad, _ := bchain.AddressDescriptorFromString("ad:" + sc.Address)
			return &ad
		}
	}
	return nil
}

func (b *CoreblockchainRPC) IsVerified(address bchain.AddressDescriptor) bool {
	// check if address is verified smart contract
	if sc := b.smartContractVerifier.GetVerified(common.Bytes2Hex(address)); sc != nil {
		return true
	}
	return false
}

// GetContractInfo returns information about smart contract
func (b *CoreblockchainRPC) GetContractInfo(contractDesc bchain.AddressDescriptor) (*bchain.ContractInfo, error) {
	cds, err := b.Parser.GetAddrDescFromAddress(common.Bytes2Hex(contractDesc[:]))
	if err != nil {
		return nil, err
	}

	now := time.Now()
	cachedContractsMux.Lock()
	contract, found := cachedContracts[common.Bytes2Hex(cds)]
	timestamp, tsFound := cachedContractsTimestamps[common.Bytes2Hex(cds)]
	cachedContractsMux.Unlock()

	// Invalidate cache if more than 2 minutes have passed
	if found && tsFound && now.Sub(timestamp) > contractCacheTTL {
		cachedContractsMux.Lock()
		delete(cachedContracts, common.Bytes2Hex(cds))
		delete(cachedContractsTimestamps, common.Bytes2Hex(cds))
		cachedContractsMux.Unlock()
		found = false
		contract = nil
	}

	if !found {
		address, err := common.HexToAddress(common.Bytes2Hex(cds))
		if err != nil {
			return nil, err
		}

		contractInfo := &bchain.ContractInfo{
			Contract: address.Hex(),
		}
		if sc := b.smartContractVerifier.GetVerified(common.Bytes2Hex(contractDesc[:])); sc != nil {
			contractInfo.Icon = sc.Icon
			contractInfo.VerifierWebAddress = sc.Web
			contractInfo.Symbol = sc.Ticker
			// supply is filled in by AddVerifiedSCData, once the decimals are known

			// RWA fields
			if sc.Metadata != nil {
				converted := make(bchain.ContractMetadata)
				for k, v := range sc.Metadata {
					converted[k] = bchain.Metadata{
						Value:  v.Value,
						Sealed: v.Sealed,
					}
				}
				contractInfo.Metadata = converted
			}
			if sc.KnownMetadata != nil {
				converted := make(bchain.ContractMetadata)
				for k, v := range sc.KnownMetadata {
					converted[k] = bchain.Metadata{
						Value:  v.Value,
						Sealed: v.Sealed,
					}
				}
				contractInfo.KnownMetadata = converted
			}
			if sc.Documents != nil {
				contractInfo.Documents = sc.Documents
			}
			if sc.LabResults != nil {
				contractInfo.LabResults = sc.LabResults
			}
		}

		isCBC721, err := b.contractSupportsInterface(contractDesc, address.Hex(), erc721InterfaceID)
		if err != nil {
			glog.Warning(errors.Annotatef(err, "cbcSupportsInterface %v", address))
		}
		if isCBC721 {
			if data, err := b.xcbCall(nameSignature, address.Hex()); err == nil {
				if name := parseCBC20StringProperty(contractDesc, data); name != "" {
					contractInfo.Name = name
				}
			} else if !strings.Contains(err.Error(), "execution reverted") {
				glog.Warning(errors.Annotatef(err, "cbc721NameSignature %v", address))
			}
			if data, err := b.xcbCall(symbolSignature, address.Hex()); err == nil {
				if symbol := parseCBC20StringProperty(contractDesc, data); symbol != "" {
					contractInfo.Symbol = symbol
				}
			} else if err != nil && !strings.Contains(err.Error(), "execution reverted") {
				glog.Warning(errors.Annotatef(err, "cbc721SymbolSignature %v", address))
			}
			if !b.smartContractVerifier.IsValidVerifiedSC(contractInfo.Contract, contractInfo.Symbol) {
				contractInfo.Symbol = ""
			}
			contractInfo.Type = CBC721TokenType
			contractInfo.Decimals = 0
			cachedContractsMux.Lock()
			cachedContracts[common.Bytes2Hex(cds)] = contractInfo
			cachedContractsTimestamps[common.Bytes2Hex(cds)] = now
			cachedContractsMux.Unlock()
			return contractInfo, nil
		}

		data, err := b.xcbCall(nameSignature, address.Hex())
		if err != nil {
			if strings.Contains(err.Error(), "execution reverted") {
				// if execution reverted -> it is not cbc20 smart contract
				contractInfo.Type = CBC721TokenType
				contractInfo.Decimals = 0
				return contractInfo, nil
			}
			return nil, nil
		}
		name := parseCBC20StringProperty(contractDesc, data)
		if name != "" {
			data, err = b.xcbCall(symbolSignature, address.Hex())
			if err != nil {
				glog.Warning(errors.Annotatef(err, "cbc20SymbolSignature %v", address))
				return nil, nil
				// return nil, errors.Annotatef(err, "cbc20SymbolSignature %v", address)
			}
			symbol := parseCBC20StringProperty(contractDesc, data)
			data, err = b.xcbCall(decimalsSignature, address.Hex())
			if err != nil {
				glog.Warning(errors.Annotatef(err, "cbc20DecimalsSignature %v", address))
				// return nil, errors.Annotatef(err, "cbc20DecimalsSignature %v", address)
			}
			contractInfo.Name = name
			if symbol != "" {
				contractInfo.Symbol = symbol
			}
			contractInfo.Type = CBC20TokenType

			// if smart contract ticker is verified but address is wrong -> do not show SC symbol (ticker)
			if !b.smartContractVerifier.IsValidVerifiedSC(contractInfo.Contract, contractInfo.Symbol) {
				contractInfo.Symbol = ""
			}
			d := parseCBC20NumericProperty(contractDesc, data)
			if d != nil {
				contractInfo.Decimals = int(uint8(d.Uint64()))
			} else {
				contractInfo.Decimals = CoreAmountDecimalPoint
			}
		} else {
			contractInfo = nil
		}
		cachedContractsMux.Lock()
		cachedContracts[common.Bytes2Hex(cds)] = contractInfo
		cachedContractsTimestamps[common.Bytes2Hex(cds)] = now
		cachedContractsMux.Unlock()
		return contractInfo, nil
	}
	return contract, nil
}

// CoreCoinTypeGetCbc20ContractBalance returns balance of cbc20 contract for given address
func (b *CoreblockchainRPC) CoreCoinTypeGetCbc20ContractBalance(addrDesc, contractDesc bchain.AddressDescriptor) (*big.Int, error) {
	addr := cutAddress(addrDesc)
	contract := "0x" + cutAddress(contractDesc)

	req := balanceOfSignature + "0000000000000000000000000000000000000000000000000000000000000000"[len(addr):] + addr
	data, err := b.xcbCall(req, contract)
	if err != nil {
		return nil, err
	}
	r := parseCBC20NumericProperty(contractDesc, data)
	if r == nil {
		return nil, errors.New("Invalid balance")
	}
	return r, nil
}

func cutAddress(addrDesc bchain.AddressDescriptor) string {
	raw := hexutil.Encode(addrDesc)

	if len(raw) > 2 {
		raw = raw[2:]
	}

	return raw
}
