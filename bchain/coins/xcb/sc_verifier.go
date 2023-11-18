package xcb

type smartContractVerifier struct {
	verified []*VerifiedSC
}

type VerifiedSC struct {
	Address     string `json:"address"`
	Icon        string `json:"icon"`
	Web         string `json:"web"`
	TotalSupply int32  `json:"totalSupply"`
}

func newSmartContractVerifier(verified []*VerifiedSC) *smartContractVerifier {
	verifier := &smartContractVerifier{
		verified: verified,
	}
	return verifier
}

func (s *smartContractVerifier) GetVerified(addr string) *VerifiedSC {
	for _, sc := range s.verified {
		if sc.Address == addr {
			return sc
		}
	}
	return nil
}
