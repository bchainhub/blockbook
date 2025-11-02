package xcb

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"path"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/juju/errors"
)

// TokenRegistry provides read-only access to the public .well-known token registry.
type TokenRegistry struct {
	baseURL    *url.URL
	httpClient *http.Client
	ttl        time.Duration

	mu            sync.RWMutex
	cache         *registrySnapshot
	cacheExpiry   time.Time
	addressPrefix string
}

type registrySnapshot struct {
	fetchedAt time.Time
	tokens    []*VerifiedSC
}

type registryPayload struct {
	Tokens []tokenPayload `json:"tokens"`
}

type tokenPayload struct {
	Address           string          `json:"address"`
	URL               string          `json:"url"`
	LegacyWeb         string          `json:"web"`
	TotalSupply       json.RawMessage `json:"totalSupply"`
	CirculatingSupply json.RawMessage `json:"circulatingSupply"`
	Ticker            string          `json:"ticker"`
	Aliases           stringList      `json:"aliases"`
	LegacyIcon        string          `json:"icon"`
	Logos             []logoPayload   `json:"logos"`
}

type logoPayload struct {
	Size int    `json:"size"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

type stringList []string

func (l *stringList) UnmarshalJSON(data []byte) error {
	if len(data) == 0 {
		*l = nil
		return nil
	}
	trimmed := strings.TrimSpace(string(data))
	if trimmed == "" || strings.EqualFold(trimmed, "null") {
		*l = nil
		return nil
	}
	switch data[0] {
	case '[':
		var tmp []string
		if err := json.Unmarshal(data, &tmp); err != nil {
			return errors.Annotate(err, "failed to parse aliases array")
		}
		normalized := make([]string, 0, len(tmp))
		for _, item := range tmp {
			item = strings.TrimSpace(item)
			if item != "" {
				normalized = append(normalized, item)
			}
		}
		*l = normalized
		return nil
	case '"':
		var single string
		if err := json.Unmarshal(data, &single); err != nil {
			return errors.Annotate(err, "failed to parse aliases string")
		}
		single = strings.TrimSpace(single)
		if single == "" {
			*l = nil
			return nil
		}
		parts := strings.Split(single, ",")
		normalized := make([]string, 0, len(parts))
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if part != "" {
				normalized = append(normalized, part)
			}
		}
		*l = normalized
		return nil
	default:
		return errors.Errorf("unsupported aliases format: %s", trimmed)
	}
}

// NewTokenRegistry creates a registry reader backed by a .well-known HTTP endpoint.
func NewTokenRegistry(rawURL string) (*TokenRegistry, error) {
	rawURL = strings.TrimSpace(rawURL)
	if rawURL == "" {
		return nil, errors.New("tokens registry url is empty")
	}
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, errors.Annotatef(err, "invalid tokens registry url %q", rawURL)
	}
	if u.Scheme == "" || u.Host == "" {
		return nil, errors.Errorf("tokens registry url must include scheme and host: %q", rawURL)
	}
	u.Path = strings.TrimSuffix(u.Path, "/")

	return &TokenRegistry{
		baseURL:    u,
		httpClient: &http.Client{Timeout: 10 * time.Second},
		ttl:        30 * time.Minute, // Cache for 30 minutes to reduce API calls
	}, nil
}

// SetAddressPrefix configures address filtering by prefix (expects lowercase).
func (r *TokenRegistry) SetAddressPrefix(prefix string) {
	prefix = strings.ToLower(strings.TrimSpace(prefix))
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.addressPrefix == prefix {
		return
	}
	r.addressPrefix = prefix
	// invalidate cache to force refetch with new prefix
	r.cache = nil
	r.cacheExpiry = time.Time{}
}

func (r *TokenRegistry) GetVerifiedSmartContracts() ([]*VerifiedSC, error) {
	glog.Info("TokenRegistry: GetVerifiedSmartContracts() called")
	start := time.Now()
	snapshot, err := r.getSnapshot()
	if err != nil {
		glog.Errorf("TokenRegistry: GetVerifiedSmartContracts() failed after %v: %v", time.Since(start), err)
		return nil, err
	}
	out := make([]*VerifiedSC, len(snapshot.tokens))
	copy(out, snapshot.tokens)
	glog.Infof("TokenRegistry: GetVerifiedSmartContracts() completed in %v, returned %d contracts", time.Since(start), len(out))
	return out, nil
}

func (r *TokenRegistry) getSnapshot() (*registrySnapshot, error) {
	now := time.Now()
	r.mu.RLock()
	if r.cache != nil && now.Before(r.cacheExpiry) {
		glog.Info("TokenRegistry: getSnapshot() - returning cached data")
		snap := r.cache
		r.mu.RUnlock()
		return snap, nil
	}
	r.mu.RUnlock()

	glog.Info("TokenRegistry: getSnapshot() - cache miss or expired, acquiring write lock")
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.cache != nil && now.Before(r.cacheExpiry) {
		glog.Info("TokenRegistry: getSnapshot() - another goroutine updated cache")
		return r.cache, nil
	}

	glog.Info("TokenRegistry: getSnapshot() - fetching fresh data from registry")
	fetchStart := time.Now()
	// Read addressPrefix while holding lock and pass it through to avoid data race
	prefix := r.addressPrefix
	snapshot, err := r.fetchRegistry(prefix)
	if err != nil {
		glog.Errorf("TokenRegistry: getSnapshot() - fetch failed after %v: %v", time.Since(fetchStart), err)
		return nil, err
	}
	glog.Infof("TokenRegistry: getSnapshot() - fetch completed in %v", time.Since(fetchStart))
	r.cache = snapshot
	r.cacheExpiry = time.Now().Add(r.ttl)
	return snapshot, nil
}

func (r *TokenRegistry) fetchRegistry(addressPrefix string) (*registrySnapshot, error) {
	// The well-known registry uses tokens.json to list addresses,
	// then individual {address}.json files for token details
	glog.Info("TokenRegistry: fetchRegistry() - fetching token addresses from tokens.json")
	start := time.Now()

	addresses, err := r.fetchTokenAddresses(addressPrefix)
	if err != nil {
		glog.Errorf("TokenRegistry: fetchRegistry() - fetchTokenAddresses failed after %v: %v", time.Since(start), err)
		return nil, err
	}
	glog.Infof("TokenRegistry: fetchRegistry() - got %d addresses in %v", len(addresses), time.Since(start))

	// Fetch individual token details
	fetchStart := time.Now()
	payload := registryPayload{
		Tokens: make([]tokenPayload, 0, len(addresses)),
	}

	for i, addr := range addresses {
		var token tokenPayload
		endpoint := r.endpointWithPrefix(fmt.Sprintf("%s.json", addr), addressPrefix)
		glog.V(2).Infof("TokenRegistry: fetchRegistry() - fetching token %d/%d: %s", i+1, len(addresses), addr)
		tokenStart := time.Now()
		notFound, err := r.getJSON(endpoint, &token)
		if err != nil {
			glog.Errorf("TokenRegistry: fetchRegistry() - getJSON failed for %s after %v: %v", addr, time.Since(tokenStart), err)
			// Continue with other tokens instead of failing completely
			continue
		}
		if notFound {
			glog.V(2).Infof("TokenRegistry: fetchRegistry() - token %s not found (404)", addr)
			continue
		}
		glog.V(2).Infof("TokenRegistry: fetchRegistry() - token %s fetched in %v", addr, time.Since(tokenStart))
		if strings.TrimSpace(token.Address) == "" {
			token.Address = addr
		}
		payload.Tokens = append(payload.Tokens, token)
	}
	glog.Infof("TokenRegistry: fetchRegistry() - fetched %d token details in %v", len(payload.Tokens), time.Since(fetchStart))

	return r.convertRegistryPayload(&payload, addressPrefix)
}

func (r *TokenRegistry) fetchTokenAddresses(addressPrefix string) ([]string, error) {
	glog.Info("TokenRegistry: fetchTokenAddresses() - starting")
	start := time.Now()

	endpoint := r.endpointWithPrefix("tokens.json", addressPrefix)
	glog.Infof("TokenRegistry: fetchTokenAddresses() - fetching from %s", endpoint)

	var response struct {
		Tokens     []string `json:"tokens"`
		Pagination struct {
			HasNext bool   `json:"hasNext"`
			Cursor  string `json:"cursor"`
		} `json:"pagination"`
	}

	notFound, err := r.getJSON(endpoint, &response)
	if err != nil {
		glog.Errorf("TokenRegistry: fetchTokenAddresses() - failed after %v: %v", time.Since(start), err)
		return nil, err
	}
	if notFound {
		glog.Errorf("TokenRegistry: fetchTokenAddresses() - token list endpoint not found at %s", endpoint)
		return nil, errors.Errorf("token list endpoint not found at %s", endpoint)
	}

	addresses := make([]string, 0, len(response.Tokens))
	prefix := strings.TrimSpace(addressPrefix)

	for _, addr := range response.Tokens {
		addr = strings.ToLower(strings.TrimSpace(addr))
		if addr == "" {
			continue
		}
		if prefix != "" && !strings.HasPrefix(addr, prefix) {
			continue
		}
		addresses = append(addresses, addr)
	}

	glog.Infof("TokenRegistry: fetchTokenAddresses() - completed in %v, got %d addresses", time.Since(start), len(addresses))
	return addresses, nil
}

func (r *TokenRegistry) getJSON(endpoint string, target interface{}) (bool, error) {
	glog.V(2).Infof("TokenRegistry: getJSON() - starting request to %s", endpoint)
	start := time.Now()

	// Create context with timeout to prevent hanging
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		glog.Errorf("TokenRegistry: getJSON() - failed to build request for %s: %v", endpoint, err)
		return false, errors.Annotatef(err, "failed to build request for %s", endpoint)
	}
	req.Header.Set("Accept", "application/json")

	glog.V(2).Infof("TokenRegistry: getJSON() - executing HTTP request to %s", endpoint)
	httpStart := time.Now()
	resp, err := r.httpClient.Do(req)
	if err != nil {
		glog.Errorf("TokenRegistry: getJSON() - HTTP request failed for %s after %v: %v", endpoint, time.Since(httpStart), err)
		return false, errors.Annotatef(err, "failed to call %s", endpoint)
	}
	defer resp.Body.Close()
	glog.V(2).Infof("TokenRegistry: getJSON() - HTTP request completed for %s in %v, status: %d", endpoint, time.Since(httpStart), resp.StatusCode)

	switch resp.StatusCode {
	case http.StatusOK:
		glog.V(2).Infof("TokenRegistry: getJSON() - decoding JSON response from %s", endpoint)
		decodeStart := time.Now()
		decoder := json.NewDecoder(resp.Body)
		decoder.UseNumber()
		if err := decoder.Decode(target); err != nil {
			glog.Errorf("TokenRegistry: getJSON() - JSON decode failed for %s after %v: %v", endpoint, time.Since(decodeStart), err)
			return false, errors.Annotatef(err, "failed to decode response from %s", endpoint)
		}
		glog.V(2).Infof("TokenRegistry: getJSON() - successfully decoded JSON from %s in %v (total: %v)", endpoint, time.Since(decodeStart), time.Since(start))
		return false, nil
	case http.StatusNotFound:
		glog.V(2).Infof("TokenRegistry: getJSON() - got 404 for %s after %v", endpoint, time.Since(start))
		_, _ = io.Copy(io.Discard, resp.Body)
		return true, nil
	default:
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
		glog.Errorf("TokenRegistry: getJSON() - unexpected status %d from %s after %v: %s", resp.StatusCode, endpoint, time.Since(start), strings.TrimSpace(string(body)))
		return false, errors.Errorf("unexpected status %d from %s: %s", resp.StatusCode, endpoint, strings.TrimSpace(string(body)))
	}
}

// endpointWithPrefix builds endpoint URL with explicit prefix (thread-safe)
func (r *TokenRegistry) endpointWithPrefix(filename string, prefix string) string {
	copyURL := *r.baseURL
	segments := []string{copyURL.Path}

	// Insert network-specific path prefix for token endpoints
	// For testnet (address prefix "ab"), use path: tokens/xab/...
	// For mainnet (address prefix "cb"), use path: tokens/...
	if filename == "tokens.json" {
		// tokens.json endpoint
		if prefix == "ab" {
			segments = append(segments, "tokens", "xab", "tokens.json")
		} else {
			segments = append(segments, "tokens", "tokens.json")
		}
	} else if strings.HasSuffix(filename, ".json") {
		// individual token: {address}.json
		if prefix == "ab" {
			segments = append(segments, "tokens", "xab", filename)
		} else {
			segments = append(segments, "tokens", filename)
		}
	} else {
		// fallback: just add filename
		segments = append(segments, filename)
	}

	joined := path.Join(segments...)
	if !strings.HasPrefix(joined, "/") {
		joined = "/" + joined
	}
	copyURL.Path = joined
	return copyURL.String()
}

func (r *TokenRegistry) convertRegistryPayload(payload *registryPayload, addressPrefix string) (*registrySnapshot, error) {
	snapshot := &registrySnapshot{
		fetchedAt: time.Now(),
		tokens:    make([]*VerifiedSC, 0, len(payload.Tokens)),
	}

	prefix := strings.TrimSpace(addressPrefix)
	include := func(addr string) bool {
		if prefix == "" {
			return true
		}
		return strings.HasPrefix(strings.ToLower(addr), prefix)
	}

	for _, token := range payload.Tokens {
		sc, skip, err := convertTokenPayload(token, include)
		if err != nil {
			return nil, err
		}
		if skip {
			continue
		}
		snapshot.tokens = append(snapshot.tokens, sc)
	}

	sort.Slice(snapshot.tokens, func(i, j int) bool { return snapshot.tokens[i].Address < snapshot.tokens[j].Address })

	return snapshot, nil
}

func convertTokenPayload(payload tokenPayload, include func(string) bool) (*VerifiedSC, bool, error) {
	address := strings.ToLower(strings.TrimSpace(payload.Address))
	if address == "" {
		return nil, false, errors.New("token entry missing address")
	}
	if include != nil && !include(address) {
		return nil, true, nil
	}
	totalSupply, err := rawMessageToBigInt(payload.TotalSupply)
	if err != nil {
		return nil, false, errors.Annotatef(err, "token %s total_supply", address)
	}
	circulatingSupply, err := rawMessageToBigInt(payload.CirculatingSupply)
	if err != nil {
		return nil, false, errors.Annotatef(err, "token %s circulating_supply", address)
	}
	aliases := make([]string, len(payload.Aliases))
	copy(aliases, payload.Aliases)

	iconURL := selectLogoURL(payload.Logos)
	if iconURL == "" {
		iconURL = strings.TrimSpace(payload.LegacyIcon)
	}
	webURL := strings.TrimSpace(payload.URL)
	if webURL == "" {
		webURL = strings.TrimSpace(payload.LegacyWeb)
	}

	sc := &VerifiedSC{
		Address:           address,
		Icon:              iconURL,
		Web:               webURL,
		TotalSupply:       totalSupply,
		CirculatingSupply: circulatingSupply,
		Ticker:            strings.TrimSpace(payload.Ticker),
		Aliases:           aliases,
	}

	return sc, false, nil
}

func rawMessageToBigInt(raw json.RawMessage) (*big.Int, error) {
	if len(raw) == 0 {
		return big.NewInt(0), nil
	}
	trimmed := strings.TrimSpace(string(raw))
	if trimmed == "" || strings.EqualFold(trimmed, "null") {
		return big.NewInt(0), nil
	}
	if raw[0] == '"' {
		var str string
		if err := json.Unmarshal(raw, &str); err != nil {
			return nil, errors.Annotate(err, "failed to parse big integer string")
		}
		return parseBigIntString(str)
	}
	var num json.Number
	if err := json.Unmarshal(raw, &num); err != nil {
		return nil, errors.Annotatef(err, "failed to parse big integer value %s", trimmed)
	}
	return parseBigIntString(num.String())
}

func parseBigIntString(value string) (*big.Int, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return big.NewInt(0), nil
	}
	sign := 1
	if strings.HasPrefix(value, "+") {
		value = value[1:]
	} else if strings.HasPrefix(value, "-") {
		sign = -1
		value = value[1:]
	}
	base := 10
	if strings.HasPrefix(value, "0x") || strings.HasPrefix(value, "0X") {
		base = 16
		value = value[2:]
	}
	if value == "" {
		return big.NewInt(0), nil
	}
	intVal := new(big.Int)
	if _, ok := intVal.SetString(value, base); ok {
		if sign < 0 {
			intVal.Neg(intVal)
		}
		return intVal, nil
	}
	if strings.ContainsAny(value, ".eE") {
		return nil, errors.Errorf("floating point values are not supported for big integers: %q", value)
	}
	if sign < 0 {
		value = "-" + value
	}
	if i, err := strconv.ParseInt(value, 10, 64); err == nil {
		return big.NewInt(i), nil
	}
	return nil, errors.Errorf("invalid integer value %q", value)
}

func selectLogoURL(logos []logoPayload) string {
	if len(logos) == 0 {
		return ""
	}
	var fallback string
	for _, logo := range logos {
		url := strings.TrimSpace(logo.URL)
		if url == "" {
			continue
		}
		if fallback == "" {
			fallback = url
		}
		if logo.Size == 32 {
			return url
		}
	}
	return fallback
}
