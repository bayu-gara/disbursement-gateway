package disbursement

type DisbursementData struct {
	UserID            int64   `json:"user_id"`
	BankCode          string  `json:"bank_code"`
	BankAccountNumber string  `json:"bank_account_number"`
	Amount            float64 `json:"amount"`
}

type DisburseRequest struct {
	DisbursementData []DisbursementData `json:"disbursement_data"`
}

type DisburseResponse struct{}

type TransferRequest struct {
	UserID            int64   `json:"user_id"`
	BankUserID        int64   `json:"bank_id"`
	BankAccountNumber string  `json:"bank_account_number"`
	Amount            float64 `json:"amount"`
}
