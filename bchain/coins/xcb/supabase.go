package xcb

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/cryptohub-digital/blockbook/bchain"
	"github.com/supabase-community/supabase-go"
)

type SupabaseClient struct {
	inner *supabase.Client
}

func NewSupabaseClient(url, key string) (*SupabaseClient, error) {
	client, err := supabase.NewClient(url, key, &supabase.ClientOptions{})
	if err != nil {
		return nil, errors.New(fmt.Sprint("failed to create Supabase client: %v", err))
	}

	return &SupabaseClient{
		inner: client,
	}, nil
}

func (s *SupabaseClient) GetVerifiedSmartContracts() ([]*VerifiedSC, error) {
	type VerifiedSCSupabase struct {
		Address           string   `json:"address" db:"address"`
		Icon              string   `json:"icon" db:"icon"`
		Web               string   `json:"web" db:"web"`
		TotalSupply       *big.Int `json:"total_supply" db:"total_supply"`
		CirculatingSupply *big.Int `json:"circulating_supply" db:"circulating_supply"`
		Ticker            string   `json:"ticker" db:"ticker"`
		Aliases           string   `json:"aliases" db:"aliases"`
	}
	var contractsSupabase []*VerifiedSCSupabase

	// Execute the query
	data, _, err := s.inner.From("verified_smart_contracts").Select("*", "exact", false).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch verified smart contracts: %v", err)
	}

	// Parse the data into the VerifiedSC struct
	if err := json.Unmarshal(data, &contractsSupabase); err != nil {
		return nil, fmt.Errorf("JSON unmarshal failed: %v", err)
	}
	contracts := []*VerifiedSC{}
	// Convert to bchain.VerifiedSC
	for _, sc := range contractsSupabase {
		contracts = append(contracts, &VerifiedSC{
			Address:           sc.Address,
			Icon:              sc.Icon,
			Web:               sc.Web,
			TotalSupply:       sc.TotalSupply,
			Ticker:            sc.Ticker,
			CirculatingSupply: sc.CirculatingSupply,
			Aliases:           strings.Split(sc.Aliases, ","),
		})
	}
	return contracts, nil
}

func (s *SupabaseClient) GetVerifiedAddresses() ([]*bchain.VerifiedAddress, error) {
	type VerifiedAddressSupabase struct {
		Type     bchain.AddressUseCaseType `json:"type" db:"type"`
		Address  string                    `json:"address" db:"address"`
		Name     string                    `json:"name" db:"name"`
		Icon     string                    `json:"icon" db:"icon"`
		URL      string                    `json:"url" db:"url"`
		URLTitle string                    `json:"urlTitle" db:"url_title"`
		Aliases  string                    `json:"aliases" db:"aliases"`
	}
	var addresses []VerifiedAddressSupabase

	// Execute the query
	data, _, err := s.inner.From("verified_addresses").Select("*", "exact", false).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch verified addresses: %v", err)
	}

	// Parse the data into the VerifiedAddress struct
	if err := json.Unmarshal(data, &addresses); err != nil {
		return nil, fmt.Errorf("JSON unmarshal failed: %v", err)
	}
	// Convert to bchain.VerifiedAddress
	var verifiedAddresses []*bchain.VerifiedAddress
	for _, addr := range addresses {
		verifiedAddresses = append(verifiedAddresses, &bchain.VerifiedAddress{
			Type:     addr.Type,
			Address:  addr.Address,
			Name:     addr.Name,
			Icon:     addr.Icon,
			URL:      addr.URL,
			URLTitle: addr.URLTitle,
			Aliases:  strings.Split(addr.Aliases, ","),
		})
	}
	return verifiedAddresses, nil
}

func (s *SupabaseClient) GetDistributedNFCSenders() ([]DistributedNFCSender, error) {
	var senders []DistributedNFCSender

	// Execute the query
	data, _, err := s.inner.From("distributed_nfc_senders").Select("*", "exact", false).Execute()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch distributed nfc senders: %v", err)
	}

	// Parse the data into the DistributedNFCSender struct
	if err := json.Unmarshal(data, &senders); err != nil {
		return nil, fmt.Errorf("JSON unmarshal failed: %v", err)
	}

	return senders, nil
}
