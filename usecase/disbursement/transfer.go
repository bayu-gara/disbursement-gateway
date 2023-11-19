package disbursement

import (
	"errors"
	"fmt"
	"log"

	bankdmn "github.com/bayu-gara/disbursement-gateway/domain/bank"
	transactiondmn "github.com/bayu-gara/disbursement-gateway/domain/transaction"
)

// transfer from/to e-wallet
func transferEmoney(fromID, toID int64, amount float64, status int) (result int64, err error) {
	lock := getLock(fromID)
	defer removeLock(fromID)

	lock.Lock()
	defer lock.Unlock()

	currentBalance := transactiondmn.GetUserBalance(fromID)
	if (currentBalance - amount) < 0 {
		return 0, errors.New(fmt.Sprintf("[transferEmoney][GetUserBalance][Insufficient balance, userID:%d amount to transfer %.2f]", fromID, amount))
	}

	result, err = transactiondmn.InsertTransaction(transactiondmn.Transaction{
		FromID: fromID,
		ToID:   toID,
		Amount: amount,
		Status: status,
	})
	if err != nil {
		log.Println(fmt.Sprintf("[transferEmoney][InsertTransaction][error:%+v]", err))
		return 0, err
	}

	return result, nil
}

// transfer from bank account pooling to destination bank account
// in this scenario we only have connection to BSI API for transfering money
func transaferToBank(srcAccount, srcBankCode, destAccount, destBankCode string, amount float64) error {
	if srcBankCode != bankdmn.BankCodeBSI {
		return errors.New("Invalid source bank code")
	}

	if destAccount != bankdmn.BankCodeBSI {
		return transaferFromBSIToOtherBank(srcAccount, destAccount, destBankCode, amount)
	}

	return transaferFromBSIToBSI(srcAccount, destAccount, amount)
}

// BSI API from transfering money from BSI to BSI
func transaferFromBSIToBSI(srcAccount, destAccount string, amount float64) error {
	//hit BSI API
	//for this simulation we assume it always success
	log.Println(fmt.Sprintf("[transaferFromBSIToBSI][Success transfer from %s to %s]", srcAccount, destAccount))
	return nil
}

// BSI API from transfering money from BSI to Other bank
func transaferFromBSIToOtherBank(srcAccount, destAccount, destBankCode string, amount float64) error {
	//hit BSI API
	//for this simulation we assume it always success
	log.Println(fmt.Sprintf("[transaferFromBSIToOtherBank][Success transfer from %s to %s]", srcAccount, destAccount))
	return nil
}
