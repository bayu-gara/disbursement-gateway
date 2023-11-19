package transaction

type Transaction struct {
	ID     int64   `json:"id"`
	FromID int64   `json:"from_id"`
	ToID   int64   `json:"to_id"`
	Amount float64 `json:"amount"`
	Status int     `json:"status"`
}
