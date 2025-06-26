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

var stableABI, _ = abi.JSON(strings.NewReader(`[
  { "inputs":[{"internalType":"string","name":"name_","type":"string"},{"internalType":"string","name":"symbol_","type":"string"}],"stateMutability":"nonpayable","type":"constructor" },

  { "anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"spender","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Approval","type":"event" },
  { "anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"from","type":"address"},{"indexed":true,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"Transfer","type":"event" },

  { "inputs":[],"name":"name","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function" },
  { "inputs":[],"name":"symbol","outputs":[{"internalType":"string","name":"","type":"string"}],"stateMutability":"view","type":"function" },
  { "inputs":[],"name":"decimals","outputs":[{"internalType":"uint8","name":"","type":"uint8"}],"stateMutability":"view","type":"function" },
  { "inputs":[],"name":"totalSupply","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function" },
  { "inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"balanceOf","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function" },

  { "inputs":[{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transfer","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function" },

  { "inputs":[{"internalType":"address","name":"owner","type":"address"},{"internalType":"address","name":"spender","type":"address"}],"name":"allowance","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function" },
  { "inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"approve","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function" },

  { "inputs":[{"internalType":"address","name":"from","type":"address"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"transferFrom","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function" },

  { "inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"addedValue","type":"uint256"}],"name":"increaseAllowance","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function" },
  { "inputs":[{"internalType":"address","name":"spender","type":"address"},{"internalType":"uint256","name":"subtractedValue","type":"uint256"}],"name":"decreaseAllowance","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"nonpayable","type":"function" },
   {
    "inputs": [],
    "name": "documents",
    "outputs": [
      {
        "components": [
          { "internalType": "string", "name": "name", "type": "string" },
          { "internalType": "string", "name": "fingerprint", "type": "string" },
          { "internalType": "string", "name": "url", "type": "string" }
        ],
        "internalType": "struct IMetadata.Document[]",
        "name": "",
        "type": "tuple[]"
      }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "labResultsURI",
    "outputs": [
      { "internalType": "string", "name": "", "type": "string" }
    ],
    "stateMutability": "view",
    "type": "function"
  },
  {
    "inputs": [],
    "name": "metadata",
    "outputs": [
      { "internalType": "string[]", "name": "keys", "type": "string[]" },
      { "internalType": "string[]", "name": "values", "type": "string[]" }
    ],
    "stateMutability": "view",
    "type": "function"
  }
]
`))

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
	Metadata   map[string]string `json:"metadata,omitempty" db:"metadata"`
	Documents  Documents         `json:"documents,omitempty" db:"documents"`
	LabResults LabResults        `json:"labResults,omitempty" db:"lab_results"`
}

type LabResults map[string]struct {
	Value interface{} `json:"value" db:"value"`
	Unit  string      `json:"unit,omitempty" db:"unit"`
}

type Documents []struct {
	Name        string `json:"name"`        // Document name
	Fingerprint string `json:"fingerprint"` // Document fingerprint
	Url         string `json:"url"`         // Document URL
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
		if sc.Ticker == ticker {
			return sc.Address == addr
		}
	}
	return true
}

func (s *smartContractVerifier) GetAllSmartContracts() []*VerifiedSC {
	verifiedSC, err := s.supabase.GetVerifiedSmartContracts()
	if err != nil {
		fmt.Println("ERROR: failed to get verified smart contracts: %v", errors.ErrorStack(err))
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
			metadata, err := s.GetLabResults(sc.Address)
			if err != nil {
				fmt.Println("ERROR: failed to get lab results for smart contract", sc.Address, ":", err)
				return nil
			}
			sc.LabResults = metadata

			documents, err := s.GetDocuments(sc.Address)
			if err != nil {
				fmt.Println("ERROR: failed to get documents for smart contract", sc.Address, ":", err)
				return nil
			}
			sc.Documents = documents

			metadataMap, err := s.GetMetadata(sc.Address)
			if err != nil {
				fmt.Println("ERROR: failed to get metadata for smart contract", sc.Address, ":", err)
				return nil
			}
			sc.Metadata = metadataMap
		}
	}
	return verifiedSC
}

func (s *smartContractVerifier) GetLabResults(addr string) (LabResults, error) {
	var labResult string
	err := s.RPC.CallContext(context.Background(), &labResult, "xcb_call", map[string]interface{}{
		"data": "0xa538d22c",
		"from": addr,
		"to":   addr,
	}, "latest")
	if err != nil {
		return nil, errors.Annotatef(err, "failed to get lab results for smart contract %s", addr)
	}

	responseBytes, err := common.Hex2BytesWithError(labResult)
	if err != nil {
		return nil, errors.Annotatef(err, "failed to convert lab results response to bytes for smart contract %s", addr)
	}
	var result string
	err = s.abi.UnpackIntoInterface(&result, "labResultsURI", responseBytes)
	if err != nil {
		return nil, errors.Annotatef(err, "failed to unpack lab results response for smart contract %s", addr)
	}

	// Download the file from the URI
	resp, err := http.Get(result)
	if err != nil {
		return nil, errors.Annotatef(err, "failed to download lab results file from URI %s", result)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download lab results file: HTTP %d", resp.StatusCode)
	}

	// Parse the JSON content
	labResults := LabResults{}
	err = json.NewDecoder(resp.Body).Decode(&labResults)
	if err != nil {
		return nil, errors.Annotatef(err, "failed to parse lab results JSON from URI %s", result)
	}

	return labResults, nil
}

func (s *smartContractVerifier) GetDocuments(addr string) (Documents, error) {
	var result string
	err := s.RPC.CallContext(context.Background(), &result, "xcb_call", map[string]interface{}{
		"data": "0x29203b91",
		"from": addr,
		"to":   addr,
	}, "latest")
	if err != nil {
		return nil, errors.Annotatef(err, "failed to get documents for smart contract %s", addr)
	}

	responseBytes, err := common.Hex2BytesWithError(result)
	if err != nil {
		return nil, errors.Annotatef(err, "failed to convert documents response to bytes for smart contract %s", addr)
	}

	var documents Documents
	err = s.abi.UnpackIntoInterface(&documents, "documents", responseBytes)
	if err != nil {
		return nil, errors.Annotatef(err, "failed to unpack documents response for smart contract %s", addr)
	}

	return documents, nil
}

func (s *smartContractVerifier) GetMetadata(addr string) (map[string]string, error) {
	var result string
	err := s.RPC.CallContext(context.Background(), &result, "xcb_call", map[string]interface{}{
		"data": "0xee57f833",
		"from": addr,
		"to":   addr,
	}, "latest")
	if err != nil {
		return nil, errors.Annotatef(err, "failed to get metadata keys for smart contract %s", addr)
	}

	responseBytes, err := common.Hex2BytesWithError(result)
	if err != nil {
		return nil, errors.Annotatef(err, "failed to convert metadata response to bytes for smart contract %s", addr)
	}

	// Unpack the response into keys and values
	var metadataResponse struct {
		Keys   []string
		Values []string
	}
	err = s.abi.UnpackIntoInterface(&metadataResponse, "metadata", responseBytes)
	if err != nil {
		return nil, errors.Annotatef(err, "failed to unpack metadata response for smart contract %s", addr)
	}

	// Combine keys and values into a map
	if len(metadataResponse.Keys) != len(metadataResponse.Values) {
		return nil, fmt.Errorf("metadata keys and values length mismatch for smart contract %s", addr)
	}

	metadata := make(map[string]string)
	for i := range metadataResponse.Keys {
		metadata[metadataResponse.Keys[i]] = metadataResponse.Values[i]
	}

	return metadata, nil
}
