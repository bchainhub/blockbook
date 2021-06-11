package xcb

import "math/big"

// Xrc20Contract contains info about Xrc20 contract
type Xrc20Contract struct {
	Contract string `json:"contract"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
}

// Xrc20Transfer contains a single Xrc20 token transfer
type Xrc20Transfer struct {
	Contract string
	From     string
	To       string
	Tokens   big.Int
}
