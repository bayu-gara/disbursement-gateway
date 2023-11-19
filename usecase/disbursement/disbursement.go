package disbursement

import (
	"errors"
	"fmt"
	"log"

	bankdmn "github.com/bayu-gara/disbursement-gateway/domain/bank"
	transactiondmn "github.com/bayu-gara/disbursement-gateway/domain/transaction"
	userdmn "github.com/bayu-gara/disbursement-gateway/domain/user"
)

func Disburse(request DisburseRequest) (err error) {
	//request to several worker
	for _, data := range request.DisbursementData {
		//for simplicity i use goroutine to do the transaction asynchronously
		//we can also use message queue/broker like NSQ. It also have retry mechanism which more secure to make sure the transaction is processed
		go doDisbursement(data)
	}

	return nil
}

func doDisbursement(request DisbursementData) (err error) {
	//user validation
	user := userdmn.GetUser(request.UserID)
	if user.ID == 0 {
		err = errors.New("Invalid User")
		log.Println(fmt.Sprintf("[doDisbursement][GetUser][error:%+v]", err))
		return err
	}

	//get bank pooling account
	bank := bankdmn.GetBank(1)
	if bank.ID == 0 {
		err = errors.New("Invalid Bank")
		log.Println(fmt.Sprintf("[doDisbursement][GetBank][error:%+v]", err))
		return err
	}

	//transfer to BSI pooling account e-wallet
	trxID, err := transferEmoney(request.UserID, bank.UserID, request.Amount, transactiondmn.TransactionPending)
	if err != nil {
		log.Println(fmt.Sprintf("[doDisbursement][transferEmoney][error:%+v]", err))
		return err
	}

	err = transaferToBank(bank.AccountNumber, bank.Code, request.BankAccountNumber, request.BankCode, request.Amount)
	if err != nil {
		log.Println(fmt.Sprintf("[doDisbursement][transaferToBank][error:%+v]", err))
		transaction := transactiondmn.GetTransactionByID(trxID)
		transaction.Status = transactiondmn.TransactionFailed
		transactiondmn.UpdateTransaction(transaction)
		return err
	}

	transaction := transactiondmn.GetTransactionByID(trxID)
	transaction.Status = transactiondmn.TransactionSuccess
	transactiondmn.UpdateTransaction(transaction)

	return nil
}

func TransactionHistory(userID int64) (result []transactiondmn.Transaction, err error) {
	result = transactiondmn.GetTransactionByUserID(userID)
	return result, nil
}
