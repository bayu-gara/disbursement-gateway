package bank

type Bank struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	UserID        int64  `json:"user_id"`
	Code          string `json:"code"`
	AccountNumber string `json:"account_number"`
}
