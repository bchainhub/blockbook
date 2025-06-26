package xcb

import (
	"fmt"

	"github.com/cryptohub-digital/blockbook/bchain"
)

type addressVerifier struct {
	supabase *SupabaseClient
}

func newAddressVerifier(supabase *SupabaseClient) *addressVerifier {
	verifier := &addressVerifier{
		supabase: supabase,
	}
	return verifier
}

func (v *addressVerifier) GetVerified(addr string) *bchain.VerifiedAddress {
	for _, address := range v.GetAllAddresses() {
		if address.Address == addr {
			return address
		}
	}
	return nil
}

func (v *addressVerifier) GetAllAddresses() []*bchain.VerifiedAddress {
	verifiedAddrs, err := v.supabase.GetVerifiedAddresses()
	if err != nil {
		fmt.Println("ERROR: failed to get verified addresses:", err)
		return nil
	}
	return verifiedAddrs
}
