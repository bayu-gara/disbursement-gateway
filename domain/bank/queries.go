package bank

func GetBank(id int64) Bank {
	for _, bank := range data {
		if bank.ID == id {
			return bank
		}
	}

	return Bank{}
}
