package transaction

var lastID = int64(7)
var data = []Transaction{
	{
		ID:     1,
		FromID: 0,
		ToID:   1,
		Amount: 10000,
		Status: 0,
	},
	{
		ID:     2,
		FromID: 1,
		ToID:   2,
		Amount: 1000,
		Status: 0,
	},
	{
		ID:     3,
		FromID: 0,
		ToID:   2,
		Amount: 10000,
		Status: 0,
	},
	{
		ID:     4,
		FromID: 0,
		ToID:   3,
		Amount: 10000,
		Status: 0,
	},
	{
		ID:     5,
		FromID: 0,
		ToID:   4,
		Amount: 10000,
		Status: 0,
	},
	{
		ID:     6,
		FromID: 0,
		ToID:   5,
		Amount: 10000,
		Status: 0,
	},
	{
		ID:     7,
		FromID: 0,
		ToID:   6,
		Amount: 10000,
		Status: 0,
	},
}
