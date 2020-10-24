package responses

type TransactionListResponse struct {
	SourceUser string  `json:"source_user"`
	DestUser   string  `json:"dest_user"`
	Amount     float64 `json:"amount"`
	Type       string  `json:"type"`
	CreatedAt  string  `json:"created_at"`
}
