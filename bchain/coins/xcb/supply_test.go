package xcb

import (
	"encoding/json"
	"math/big"
	"testing"
)

func bigOrNil(t *testing.T, s string) *big.Int {
	t.Helper()
	if s == "" {
		return nil
	}
	v, ok := new(big.Int).SetString(s, 10)
	if !ok {
		t.Fatalf("bad test fixture %q", s)
	}
	return v
}

func TestFormatTokenSupply(t *testing.T) {
	tests := []struct {
		name     string
		supply   string
		decimals int
		want     string
	}{
		// real mainnet values, base units as returned by totalSupply()
		{name: "USDX", supply: "1027551000", decimals: 6, want: "1,027.551"},
		{name: "EURX", supply: "1267597246", decimals: 6, want: "1,267.597246"},
		{name: "THBX not yet minted", supply: "0", decimals: 6, want: "0"},
		{name: "CTN", supply: "1000000000000000000000000", decimals: 18, want: "1,000,000"},

		{name: "unknown supply renders empty", supply: "", decimals: 6, want: ""},
		{name: "below one whole token keeps leading zeroes", supply: "50", decimals: 6, want: "0.00005"},
		{name: "smallest representable unit", supply: "1", decimals: 18, want: "0.000000000000000001"},
		{name: "no decimals is a plain grouped integer", supply: "1234567", decimals: 0, want: "1,234,567"},
		{name: "fraction is padded not truncated", supply: "1000001", decimals: 6, want: "1.000001"},
		{name: "trailing zeroes are dropped", supply: "1500000", decimals: 6, want: "1.5"},
		{name: "value larger than uint64", supply: "123456789012345678901234567890", decimals: 6, want: "123,456,789,012,345,678,901,234.56789"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatTokenSupply(bigOrNil(t, tt.supply), tt.decimals)
			if got != tt.want {
				t.Errorf("formatTokenSupply(%q, %d) = %q, want %q", tt.supply, tt.decimals, got, tt.want)
			}
		})
	}
}

// A supply the registry does not state must be distinguishable from a stated zero,
// otherwise every token without a supply in the registry is reported as having none.
func TestRawMessageToBigIntDistinguishesAbsentFromZero(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want string // "" means nil
	}{
		{name: "absent", raw: "", want: ""},
		{name: "null", raw: "null", want: ""},
		{name: "empty string", raw: `""`, want: ""},
		{name: "explicit zero", raw: "0", want: "0"},
		{name: "number", raw: "1000000000000000000000000", want: "1000000000000000000000000"},
		{name: "string", raw: `"1027551000"`, want: "1027551000"},
		{name: "hex string", raw: `"0x3d3f2f18"`, want: "1027551000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := rawMessageToBigInt(json.RawMessage(tt.raw))
			if err != nil {
				t.Fatalf("rawMessageToBigInt(%q) error: %v", tt.raw, err)
			}
			if tt.want == "" {
				if got != nil {
					t.Errorf("rawMessageToBigInt(%q) = %v, want nil", tt.raw, got)
				}
				return
			}
			if got == nil || got.String() != tt.want {
				t.Errorf("rawMessageToBigInt(%q) = %v, want %s", tt.raw, got, tt.want)
			}
		})
	}
}

// The registry publishes supply in base units, the same unit the contract reports,
// so a token entry must survive the trip from JSON to VerifiedSC unscaled.
func TestConvertTokenPayloadKeepsRegistrySupplyInBaseUnits(t *testing.T) {
	var payload tokenPayload
	// trimmed copy of the live .well-known entry for CTN
	raw := `{"address":"cb19c7acc4c292d2943ba23c2eaa5d9c5a6652a8710c","ticker":"CTN",
		"decimals":18,"totalSupply":1000000000000000000000000,"categories":["corepass","payment"]}`
	if err := json.Unmarshal([]byte(raw), &payload); err != nil {
		t.Fatal(err)
	}
	sc, skip, err := convertTokenPayload(payload, nil)
	if err != nil || skip {
		t.Fatalf("convertTokenPayload: err=%v skip=%v", err, skip)
	}
	if sc.TotalSupply == nil || sc.TotalSupply.String() != "1000000000000000000000000" {
		t.Errorf("TotalSupply = %v, want 1000000000000000000000000", sc.TotalSupply)
	}
	if sc.CirculatingSupply != nil {
		t.Errorf("CirculatingSupply = %v, want nil (registry does not state it)", sc.CirculatingSupply)
	}
	if got := formatTokenSupply(sc.TotalSupply, 18); got != "1,000,000" {
		t.Errorf("rendered supply = %q, want %q", got, "1,000,000")
	}
}

// A token whose entry carries no supply must fall through to the chain rather than
// report zero, which is what made every MoneyX stablecoin show a total supply of 0.
func TestConvertTokenPayloadLeavesAbsentSupplyUnset(t *testing.T) {
	var payload tokenPayload
	// trimmed copy of the live .well-known entry for USDX
	raw := `{"address":"cb313b5e681ecd5e6086cc5609bfba595c2b8885488f","ticker":"USDX",
		"decimals":6,"categories":["corepass","stable"]}`
	if err := json.Unmarshal([]byte(raw), &payload); err != nil {
		t.Fatal(err)
	}
	sc, skip, err := convertTokenPayload(payload, nil)
	if err != nil || skip {
		t.Fatalf("convertTokenPayload: err=%v skip=%v", err, skip)
	}
	if sc.TotalSupply != nil {
		t.Errorf("TotalSupply = %v, want nil so that the chain value is used", sc.TotalSupply)
	}
	if sc.CirculatingSupply != nil {
		t.Errorf("CirculatingSupply = %v, want nil so that the chain value is used", sc.CirculatingSupply)
	}
}

func TestResolveSupply(t *testing.T) {
	tests := []struct {
		name            string
		sc              *VerifiedSC
		chain           string
		chainCalls      int
		wantTotal       string
		wantCirculating string
	}{
		{
			name:            "registry silent, chain is used for both", // USDX, EURX, THBX, CNYX
			sc:              &VerifiedSC{},
			chain:           "1027551000",
			chainCalls:      1,
			wantTotal:       "1027551000",
			wantCirculating: "1027551000",
		},
		{
			name:            "registry total overrides the chain and the chain is not called", // CTN
			sc:              &VerifiedSC{TotalSupply: bigOrNil(t, "1000000000000000000000000")},
			chain:           "999",
			chainCalls:      0,
			wantTotal:       "1000000000000000000000000",
			wantCirculating: "1000000000000000000000000",
		},
		{
			name:            "registry states circulating separately",
			sc:              &VerifiedSC{TotalSupply: bigOrNil(t, "1000"), CirculatingSupply: bigOrNil(t, "400")},
			chain:           "999",
			chainCalls:      0,
			wantTotal:       "1000",
			wantCirculating: "400",
		},
		{
			name:            "registry states only circulating, total comes from the chain",
			sc:              &VerifiedSC{CirculatingSupply: bigOrNil(t, "400")},
			chain:           "1000",
			chainCalls:      1,
			wantTotal:       "1000",
			wantCirculating: "400",
		},
		{
			name:            "unverified token still gets its chain supply",
			sc:              nil,
			chain:           "42",
			chainCalls:      1,
			wantTotal:       "42",
			wantCirculating: "42",
		},
		{
			name:            "explicit zero in the registry is honoured, not treated as absent",
			sc:              &VerifiedSC{TotalSupply: bigOrNil(t, "0")},
			chain:           "500",
			chainCalls:      0,
			wantTotal:       "0",
			wantCirculating: "0",
		},
		{
			name:            "unreadable contract leaves the supply unknown rather than zero",
			sc:              &VerifiedSC{},
			chain:           "",
			chainCalls:      1,
			wantTotal:       "",
			wantCirculating: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calls := 0
			total, circulating := resolveSupply(tt.sc, func() *big.Int {
				calls++
				return bigOrNil(t, tt.chain)
			})
			if calls != tt.chainCalls {
				t.Errorf("chain was called %d times, want %d", calls, tt.chainCalls)
			}
			if got := formatTokenSupply(total, 0); got != groupDigits(tt.wantTotal) {
				t.Errorf("total = %q, want %q", got, tt.wantTotal)
			}
			if got := formatTokenSupply(circulating, 0); got != groupDigits(tt.wantCirculating) {
				t.Errorf("circulating = %q, want %q", got, tt.wantCirculating)
			}
		})
	}
}

// totalSupply() must be addressed by its Core SHA3-256 selector; a keccak selector
// would silently return empty data and every token would look supply-less.
func TestTotalSupplySignature(t *testing.T) {
	if totalSupplySignature != "0x1f1881f8" {
		t.Errorf("totalSupplySignature = %s, want 0x1f1881f8 (sha3-256(\"totalSupply()\")[:4])", totalSupplySignature)
	}
}
