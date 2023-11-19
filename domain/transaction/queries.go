package transaction

import (
	"sync"
)

var mutex sync.Mutex

func GetTransactionByID(trxID int64) (result Transaction) {
	for _, transaction := range data {
		if transaction.ID == trxID {
			return transaction
		}
	}

	return result
}

func GetTransactionByUserID(userID int64) (result []Transaction) {
	for _, transaction := range data {
		if transaction.FromID == userID {
			result = append(result, transaction)
			continue
		}

		if transaction.ToID == userID {
			result = append(result, transaction)
			continue
		}
	}

	return result
}

func GetTransactionByUserIDAndStatus(userID int64, status int) (result []Transaction) {
	for _, transaction := range data {
		if transaction.Status != status {
			continue
		}

		if transaction.FromID == userID {
			result = append(result, transaction)
			continue
		}

		if transaction.ToID == userID {
			result = append(result, transaction)
			continue
		}
	}

	return result
}

func GetSuccessTransactionByUserID(userID int64) []Transaction {
	return GetTransactionByUserIDAndStatus(userID, TransactionSuccess)
}

func GetUserBalance(userID int64) (result float64) {
	transactions := GetTransactionByUserID(userID)
	for _, transaction := range transactions {
		if transaction.FromID == userID {
			if transaction.Status != TransactionSuccess {
				result += transaction.Amount
				continue
			}

			result -= transaction.Amount
			continue
		}

		if transaction.ToID == userID {
			if transaction.Status == TransactionSuccess {
				result += transaction.Amount
				continue
			}
		}
	}

	return result
}

func InsertTransaction(transaction Transaction) (int64, error) {
	mutex.Lock()
	defer mutex.Unlock()
	lastID += 1
	transaction.ID = lastID
	data = append(data, transaction)

	return lastID, nil
}

func UpdateTransaction(transaction Transaction) (int64, error) {
	mutex.Lock()
	defer mutex.Unlock()

	for i := range data {
		if data[i].ID == transaction.ID {
			data[i].FromID = transaction.FromID
			data[i].ToID = transaction.ToID
			data[i].Amount = transaction.Amount
			data[i].Status = transaction.Status
		}
	}

	return lastID, nil
}
