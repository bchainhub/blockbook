package xcb

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/core-coin/go-core/v2/accounts/abi"
	"github.com/core-coin/go-core/v2/common"
	"github.com/juju/errors"
)

var stableABI, _ = abi.JSON(strings.NewReader(`[{"inputs":[{"internalType":"string","name":"name_","type":"string"},{"internalType":"string","name":"symbol_","type":"string"},{"internalType":"string[]","name":"initKeys","type":"string[]"},{"internalType":"string[]","name":"initValues","type":"string[]"},{"internalType":"bool[]","name":"initSealedFlags","type":"bool[]"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"spender","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Approval","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Transfer","type":"event"},{"inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"spender","type":"address"}],"name":"allowance","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"approve","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"count","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"decimals","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"subtractedValue","type":"uint256"}],"name":"decreaseAllowance","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"getAllKeys","outputs":[{"internalType":"string[]","name":"keys","type":"string[]"},{"internalType":"string[]","name":"values","type":"string[]"},{"internalType":"bool[]","name":"sealedFlags","type":"bool[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"index","type":"uint256"}],"name":"getByIndex","outputs":[{"internalType":"string","name":"","type":"string"},{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"key","type":"string"}],"name":"getValue","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"key","type":"string"}],"name":"hasKey","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"addedValue","type":"uint256"}],"name":"increaseAllowance","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"key","type":"string"}],"name":"isSealed","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"listKeys","outputs":[{"internalType":"string[]","name":"","type":"string[]"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"mint","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"renounceOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"key","type":"string"}],"name":"sealKey","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"key","type":"string"},{"internalType":"string","name":"value","type":"string"}],"name":"setValue","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"totalSupply","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transfer","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transferFrom","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"}]`))
var knownKeys = []string{"lab", "documents", "tradingStop", "tokenExpiration"}

type smartContractVerifier struct {
	RPC      CVMRPCClient
	abi      abi.ABI
	supabase *SupabaseClient
}

type VerifiedSC struct {
	Address           string   `json:"address" db:"address"`
	Icon              string   `json:"icon" db:"icon"`
	Web               string   `json:"web" db:"web"`
	TotalSupply       *big.Int `json:"total_supply" db:"total_supply"`
	CirculatingSupply *big.Int `json:"circulating_supply" db:"circulating_supply"`
	Ticker            string   `json:"ticker" db:"ticker"`
	Aliases           []string `json:"aliases" db:"aliases"`

	// RWA Metadata
	Metadata ContractMetadata `json:"metadata,omitempty" db:"metadata"`

	KnownMetadata ContractMetadata `json:"knownMetadata,omitempty" db:"known_metadata"`
	Documents     Documents        `json:"documents,omitempty" db:"documents"`
	LabResults    LabResults       `json:"labResults,omitempty" db:"lab_results"`
}

type LabResults map[string]struct {
	Value interface{} `json:"value" db:"value"`
	Unit  string      `json:"unit,omitempty" db:"unit"`
}

type Documents []struct {
	Name        string `json:"name"`        // Document name
	Fingerprint string `json:"fingerprint"` // Document fingerprint
	Url         string `json:"location"`    // Document URL
}

func newSmartContractVerifier(supabase *SupabaseClient) *smartContractVerifier {
	verifier := &smartContractVerifier{
		supabase: supabase,
		abi:      stableABI,
	}
	return verifier
}

func (s *smartContractVerifier) GetVerified(addr string) *VerifiedSC {
	for _, sc := range s.GetAllSmartContracts() {
		if sc.Address == addr {
			return sc
		}
	}
	return nil
}

func (s *smartContractVerifier) IsValidVerifiedSC(addr, ticker string) bool {
	for _, sc := range s.GetAllSmartContracts() {
		fmt.Println("Checking smart contract:", sc.Address, "Ticker:", sc.Ticker, "Addr:", addr, "Ticker:", ticker)
		if sc.Ticker == ticker {
			return sc.Address == addr
		}
	}
	return true
}

func (s *smartContractVerifier) GetAllSmartContracts() []*VerifiedSC {
	verifiedSC, err := s.supabase.GetVerifiedSmartContracts()
	if err != nil {
		fmt.Printf("ERROR: failed to get verified smart contracts: %v\n", errors.ErrorStack(err))
		return nil
	}

	for _, sc := range verifiedSC {
		if sc.CirculatingSupply.Cmp(big.NewInt(0)) < 0 { // RWA Smart Contract
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			var response string
			err := s.RPC.CallContext(ctx, &response, "xcb_call", map[string]interface{}{
				"data": "0x1f1881f8",
				"from": sc.Address,
				"to":   sc.Address,
			}, "latest")
			if err != nil {
				fmt.Println("ERROR: failed to get total supply for smart contract", sc.Address, ":", err)
				return nil
			}

			responseBytes, err := common.Hex2BytesWithError(response)
			if err != nil {
				fmt.Println("ERROR: failed to convert response to bytes for smart contract", sc.Address, ":", err)
				return nil
			}
			var result *big.Int = big.NewInt(0)
			// result.SetBytes(responseBytes)
			err = s.abi.UnpackIntoInterface(&result, "totalSupply", responseBytes)
			if err != nil {
				fmt.Println("ERROR: failed to unpack response for smart contract", sc.Address, ":", err)
				return nil
			}

			// divide by 10^18 to get the correct value
			sc.CirculatingSupply = new(big.Int).Div(result, big.NewInt(1e18))
			sc.TotalSupply = new(big.Int).Div(result, big.NewInt(1e18))

			// Add RWA Metadata
			metadata, err := s.GetAllMetadata(sc.Address)
			if err != nil {
				fmt.Println("ERROR: failed to get all metadata for smart contract", sc.Address, ":", err)
				return nil
			}

			sc.Metadata, sc.KnownMetadata = SplitMetadata(metadata, knownKeys)

			labRes, err := s.GetLabResults(metadata["lab"])
			if err != nil {
				fmt.Println("ERROR: failed to get lab results for smart contract", sc.Address, ":", err)
				return nil
			}
			sc.LabResults = labRes

			documents, err := s.GetDocuments(metadata["documents"])
			if err != nil {
				fmt.Println("ERROR: failed to get documents for smart contract", sc.Address, ":", err)
				return nil
			}
			sc.Documents = documents
		}
	}
	return verifiedSC
}

type ContractMetadata map[string]Metadata

type Metadata struct {
	Value  string `json:"value"`
	Sealed bool   `json:"sealed"`
}

func (s *smartContractVerifier) GetAllMetadata(addr string) (ContractMetadata, error) {
	// Encode the call to getAllKeys()
	data, err := s.abi.Pack("getAllKeys")
	if err != nil {
		return nil, errors.Annotate(err, "failed to pack getAllKeys()")
	}

	var result string
	err = s.RPC.CallContext(context.Background(), &result, "xcb_call", map[string]interface{}{
		"data": fmt.Sprintf("0x%x", data),
		"from": addr,
		"to":   addr,
	}, "latest")
	if err != nil {
		return nil, errors.Annotatef(err, "failed to get all keys for smart contract %s", addr)
	}

	responseBytes, err := common.Hex2BytesWithError(result)
	if err != nil {
		return nil, errors.Annotatef(err, "failed to convert getAllKeys response to bytes for smart contract %s", addr)
	}

	// Unpack the response into keys, values, and sealedFlags
	var keys []string
	var values []string
	var sealedFlags []bool
	err = s.abi.UnpackIntoInterface(&[]interface{}{&keys, &values, &sealedFlags}, "getAllKeys", responseBytes)
	if err != nil {
		return nil, errors.Annotatef(err, "failed to unpack getAllKeys response for smart contract %s", addr)
	}

	if len(keys) != len(values) || len(keys) != len(sealedFlags) {
		return nil, fmt.Errorf("getAllKeys: keys, values, and sealedFlags length mismatch for smart contract %s", addr)
	}

	metadata := ContractMetadata{}
	for i := range keys {
		// Create a new entry in the metadata map
		metadata[keys[i]] = struct {
			Value  string `json:"value"`
			Sealed bool   `json:"sealed"`
		}{
			Value:  values[i],
			Sealed: sealedFlags[i],
		}
	}

	return metadata, nil
}

func (s *smartContractVerifier) GetLabResults(labResultsMetadata Metadata) (LabResults, error) {
	// Download the file from the URI
	resp, err := http.Get(labResultsMetadata.Value)
	if err != nil {
		return nil, errors.Annotatef(err, "failed to download lab results file from URI %s", labResultsMetadata.Value)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download lab results file: HTTP %d", resp.StatusCode)
	}

	// Parse the JSON content
	labResults := LabResults{}
	err = json.NewDecoder(resp.Body).Decode(&labResults)
	if err != nil {
		return nil, errors.Annotatef(err, "failed to parse lab results JSON from URI %s", labResultsMetadata.Value)
	}

	return labResults, nil
}

func (s *smartContractVerifier) GetDocuments(documentsMetadata Metadata) (Documents, error) {

	// Download the file from the URI
	resp, err := http.Get(documentsMetadata.Value)
	if err != nil {
		return nil, errors.Annotatef(err, "failed to download documents results file from URI %s", documentsMetadata.Value)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download documents results file: HTTP %d", resp.StatusCode)
	}

	// Parse the JSON content
	documentsResults := Documents{}
	err = json.NewDecoder(resp.Body).Decode(&documentsResults)
	if err != nil {
		return nil, errors.Annotatef(err, "failed to parse lab results JSON from URI %s", documentsMetadata.Value)
	}

	return documentsResults, nil
}

// SplitMetadata separates known metadata keys from general metadata.
// knownKeys is a list of keys to move to KnownMetadata.
// Returns (metadata, knownMetadata)
func SplitMetadata(metadata ContractMetadata, knownKeys []string) (restMetadata ContractMetadata, knownMetadata ContractMetadata) {
	knownSet := make(map[string]struct{}, len(knownKeys))
	for _, k := range knownKeys {
		knownSet[k] = struct{}{}
	}
	restMetadata = make(ContractMetadata)
	knownMetadata = make(ContractMetadata)

	for k, v := range metadata {
		if _, isKnown := knownSet[k]; isKnown {
			knownMetadata[k] = v
		} else {
			restMetadata[k] = v
		}
	}
	return restMetadata, knownMetadata
}

func (s *smartContractVerifier) GetTicker(addr string) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var response string
	err := s.RPC.CallContext(ctx, &response, "xcb_call", map[string]interface{}{
		"data": "0x231782d8", // sha3("symbol()")
		"from": addr,
		"to":   addr,
	}, "latest")
	if err != nil {
		fmt.Println("ERROR: failed to get ticker for smart contract", addr, ":", err)
		return ""
	}

	responseBytes, err := common.Hex2BytesWithError(response)
	if err != nil {
		fmt.Println("ERROR: failed to convert response to bytes for smart contract", addr, ":", err)
		return ""
	}

	var result string
	err = s.abi.UnpackIntoInterface(&result, "symbol", responseBytes)
	if err != nil {
		fmt.Println("ERROR: failed to unpack response for smart contract", addr, ":", err)
		return ""
	}

	return result
}
