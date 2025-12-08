package xcb

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"sync"
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
	registry *TokenRegistry

	// Cache for processed smart contracts
	cacheMu          sync.RWMutex
	cachedContracts  []*VerifiedSC
	cacheExpiry      time.Time
	cacheTTL         time.Duration

	// Background refresh control
	stopRefresh chan struct{}
	refreshDone sync.WaitGroup
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

func newSmartContractVerifier(registry *TokenRegistry) *smartContractVerifier {
	verifier := &smartContractVerifier{
		registry:    registry,
		abi:         stableABI,
		cacheTTL:    30 * time.Minute, // Cache processed contracts for 30 minutes
		stopRefresh: make(chan struct{}),
	}
	// Start background refresh goroutine
	verifier.startBackgroundRefresh()
	return verifier
}

// startBackgroundRefresh starts a background goroutine that periodically refreshes the cache
func (s *smartContractVerifier) startBackgroundRefresh() {
	s.refreshDone.Add(1)
	go func() {
		defer s.refreshDone.Done()
		// Initial delay before first refresh (5 seconds)
		timer := time.NewTimer(5 * time.Second)
		defer timer.Stop()

		for {
			select {
			case <-timer.C:
				// Refresh the cache
				s.refreshCache()
				// Schedule next refresh
				timer.Reset(s.cacheTTL)
			case <-s.stopRefresh:
				return
			}
		}
	}()
}

// Stop stops the background refresh goroutine
func (s *smartContractVerifier) Stop() {
	close(s.stopRefresh)
	s.refreshDone.Wait()
}

// refreshCache refreshes the cache in the background
func (s *smartContractVerifier) refreshCache() {
	// This will trigger a refresh if needed
	_ = s.GetAllSmartContracts()
}

func (s *smartContractVerifier) GetVerified(addr string) *VerifiedSC {
	for _, sc := range s.GetAllSmartContracts() {
		if strings.EqualFold(sc.Address, addr) {
			return sc
		}
	}
	return nil
}

func (s *smartContractVerifier) IsValidVerifiedSC(addr, ticker string) bool {
	for _, sc := range s.GetAllSmartContracts() {
		if sc.Ticker == ticker {
			return strings.EqualFold(sc.Address, addr)
		}
	}
	return true
}

func (s *smartContractVerifier) GetAllSmartContracts() []*VerifiedSC {
	fmt.Println("SmartContractVerifier: GetAllSmartContracts() called")
	start := time.Now()

	// Check cache first
	now := time.Now()
	s.cacheMu.RLock()
	if s.cachedContracts != nil && now.Before(s.cacheExpiry) {
		fmt.Printf("SmartContractVerifier: GetAllSmartContracts() - returning cached data (%d contracts)\n", len(s.cachedContracts))
		contracts := s.cachedContracts
		s.cacheMu.RUnlock()
		return contracts
	}
	s.cacheMu.RUnlock()

	fmt.Println("SmartContractVerifier: GetAllSmartContracts() - cache miss/expired, acquiring write lock")
	// Cache miss or expired, acquire write lock
	s.cacheMu.Lock()
	defer s.cacheMu.Unlock()

	// Double-check in case another goroutine just updated the cache
	if s.cachedContracts != nil && now.Before(s.cacheExpiry) {
		fmt.Println("SmartContractVerifier: GetAllSmartContracts() - another goroutine updated cache")
		return s.cachedContracts
	}

	// Fetch and process contracts
	fmt.Println("SmartContractVerifier: GetAllSmartContracts() - fetching from registry")
	registryStart := time.Now()
	verifiedSC, err := s.registry.GetVerifiedSmartContracts()
	if err != nil {
		fmt.Printf("ERROR: SmartContractVerifier: failed to get verified smart contracts after %v: %v\n", time.Since(registryStart), errors.ErrorStack(err))
		return nil
	}
	fmt.Printf("SmartContractVerifier: GetAllSmartContracts() - got %d contracts from registry in %v\n", len(verifiedSC), time.Since(registryStart))

	rwaCount := 0
	for i, sc := range verifiedSC {
		if sc.CirculatingSupply.Cmp(big.NewInt(0)) < 0 { // RWA Smart Contract
			rwaCount++
			fmt.Printf("SmartContractVerifier: Processing RWA contract %d/%d: %s\n", rwaCount, len(verifiedSC), sc.Address)
			rpcStart := time.Now()

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			var response string
			err := s.RPC.CallContext(ctx, &response, "xcb_call", map[string]interface{}{
				"data": "0x1f1881f8",
				"from": sc.Address,
				"to":   sc.Address,
			}, "latest")
			cancel()
			if err != nil {
				fmt.Printf("ERROR: failed to get total supply for smart contract %s after %v: %v\n", sc.Address, time.Since(rpcStart), err)
				continue // Skip this contract instead of returning nil
			}
			fmt.Printf("SmartContractVerifier: RPC call for %s completed in %v\n", sc.Address, time.Since(rpcStart))

			responseBytes, err := common.Hex2BytesWithError(response)
			if err != nil {
				fmt.Println("ERROR: failed to convert response to bytes for smart contract", sc.Address, ":", err)
				continue
			}
			var result *big.Int = big.NewInt(0)
			err = s.abi.UnpackIntoInterface(&result, "totalSupply", responseBytes)
			if err != nil {
				fmt.Println("ERROR: failed to unpack response for smart contract", sc.Address, ":", err)
				continue
			}

			// divide by 10^18 to get the correct value
			sc.CirculatingSupply = new(big.Int).Div(result, big.NewInt(1e18))
			sc.TotalSupply = new(big.Int).Div(result, big.NewInt(1e18))

			// Add RWA Metadata
			fmt.Printf("SmartContractVerifier: Fetching metadata for %s\n", sc.Address)
			metadataStart := time.Now()
			metadata, err := s.GetAllMetadata(sc.Address)
			if err != nil {
				fmt.Printf("ERROR: failed to get all metadata for smart contract %s after %v: %v\n", sc.Address, time.Since(metadataStart), err)
				continue
			}
			fmt.Printf("SmartContractVerifier: Metadata for %s fetched in %v\n", sc.Address, time.Since(metadataStart))

			sc.Metadata, sc.KnownMetadata = SplitMetadata(metadata, knownKeys)

			// Only fetch lab results if URL is provided
			if labURL := metadata["lab"].Value; strings.TrimSpace(labURL) != "" {
				fmt.Printf("SmartContractVerifier: Fetching lab results from %s\n", labURL)
				labStart := time.Now()
				labRes, err := s.GetLabResults(metadata["lab"])
				if err != nil {
					fmt.Printf("ERROR: failed to get lab results for smart contract %s after %v: %v\n", sc.Address, time.Since(labStart), err)
					// Continue without lab results
				} else {
					fmt.Printf("SmartContractVerifier: Lab results fetched in %v\n", time.Since(labStart))
					sc.LabResults = labRes
				}
			}

			// Only fetch documents if URL is provided
			if docURL := metadata["documents"].Value; strings.TrimSpace(docURL) != "" {
				fmt.Printf("SmartContractVerifier: Fetching documents from %s\n", docURL)
				docStart := time.Now()
				documents, err := s.GetDocuments(metadata["documents"])
				if err != nil {
					fmt.Printf("ERROR: failed to get documents for smart contract %s after %v: %v\n", sc.Address, time.Since(docStart), err)
					// Continue without documents
				} else {
					fmt.Printf("SmartContractVerifier: Documents fetched in %v\n", time.Since(docStart))
					sc.Documents = documents
				}
			}
		} else {
			fmt.Printf("SmartContractVerifier: Skipping non-RWA contract %d/%d: %s\n", i+1, len(verifiedSC), sc.Address)
		}
	}

	// Update cache
	s.cachedContracts = verifiedSC
	s.cacheExpiry = now.Add(s.cacheTTL)

	fmt.Printf("SmartContractVerifier: GetAllSmartContracts() completed in %v, cached %d contracts (including %d RWA)\n", time.Since(start), len(verifiedSC), rwaCount)
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
	uri := strings.TrimSpace(labResultsMetadata.Value)
	if uri == "" {
		return nil, errors.New("empty lab results URI")
	}

	// Download the file from the URI
	resp, err := http.Get(uri)
	if err != nil {
		return nil, errors.Annotatef(err, "failed to download lab results file from URI %s", uri)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download lab results file: HTTP %d", resp.StatusCode)
	}

	// Parse the JSON content
	labResults := LabResults{}
	err = json.NewDecoder(resp.Body).Decode(&labResults)
	if err != nil {
		return nil, errors.Annotatef(err, "failed to parse lab results JSON from URI %s", uri)
	}

	return labResults, nil
}

func (s *smartContractVerifier) GetDocuments(documentsMetadata Metadata) (Documents, error) {
	uri := strings.TrimSpace(documentsMetadata.Value)
	if uri == "" {
		return nil, errors.New("empty documents URI")
	}

	// Download the file from the URI
	resp, err := http.Get(uri)
	if err != nil {
		return nil, errors.Annotatef(err, "failed to download documents file from URI %s", uri)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download documents file: HTTP %d", resp.StatusCode)
	}

	// Parse the JSON content
	documentsResults := Documents{}
	err = json.NewDecoder(resp.Body).Decode(&documentsResults)
	if err != nil {
		return nil, errors.Annotatef(err, "failed to parse documents JSON from URI %s", uri)
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
