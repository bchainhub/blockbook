package xcb

type smartContractVerifier struct {
	verified []*VerifiedSC
}

type VerifiedSC struct {
	Address     string `json:"address"`
	Icon        string `json:"icon"`
	Web         string `json:"web"`
	TotalSupply int32  `json:"totalSupply"`
	Ticker      string `json:"ticker"`
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

func (s *smartContractVerifier) IsValidVerifiedSC(addr, ticker string) bool {
	for _, sc := range s.verified {
		if sc.Ticker == ticker {
			return sc.Address == addr
		}
	}
	return true
}
