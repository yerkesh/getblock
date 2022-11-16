package domain

type BlockTransactions struct {
	Hash        string        `json:"hash"`
	Number      string        `json:"number"`
	ParentHash  string        `json:"parentHash"`
	Transaction []Transaction `json:"transactions"`
}

type Transaction struct {
	BlockHash   string `json:"blockHash"`
	BlockNumber string `json:"blockNumber"`
	From        string `json:"from"`
	To          string `json:"to"`
	Value       string `json:"value"`
	Gas         string `json:"gas"`
	GasPrice    string `json:"gasPrice"`
}
