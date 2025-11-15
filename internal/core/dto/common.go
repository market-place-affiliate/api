package dto


type Response[T any] struct {
	HttpCode int               `json:"-"`
	Success  bool              `json:"success"`
	Code     int               `json:"code"`
	Message  string            `json:"message"`
	TxnID    string            `json:"txn_id"`
	Data     T                 `json:"data,omitempty"`
}
