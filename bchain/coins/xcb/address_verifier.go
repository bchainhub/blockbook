package xcb

import "github.com/cryptohub-digital/blockbook/bchain"

type addressVerifier struct {
	verified []*bchain.VerifiedAddress
}

func newAddressVerifier(verified []*bchain.VerifiedAddress) *addressVerifier {
	verifier := &addressVerifier{
		verified: verified,
	}
	return verifier
}

func (v *addressVerifier) GetVerified(addr string) *bchain.VerifiedAddress {
	for _, address := range v.verified {
		if address.Address == addr {
			return address
		}
	}
	return nil
}

func (v *addressVerifier) GetAllAddresses() []*bchain.VerifiedAddress {
	return v.verified
}
