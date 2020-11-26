package models

type CallbackPayload struct {
	ID		int64	`json:"id,omitempty"`
	Address string `json:address`
	Value   int    `json:value`
	TxHash  string `json:transactionHash`
	Status  bool   `json:status`
	TransactionID int64 `json:"transactionId,omitempty`
}