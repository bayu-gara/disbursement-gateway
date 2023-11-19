package http

import (
	"fmt"
	"log"
	gohttp "net/http"
	"strconv"

	"github.com/bayu-gara/disbursement-gateway/usecase/disbursement"
	jsoniter "github.com/json-iterator/go"
)

func disburse(responseWriter gohttp.ResponseWriter, httpReq *gohttp.Request) {
	var request disbursement.DisburseRequest

	err := jsoniter.NewDecoder(httpReq.Body).Decode(&request)
	if err != nil {
		log.Println(fmt.Sprintf("[disburse][Decode][error:%+v]", err))
		WriteError(responseWriter, err)
		return
	}

	err = disbursement.Disburse(request)
	if err != nil {
		log.Println(fmt.Sprintf("[disburse][Disburse][error:%+v]", err))
		WriteError(responseWriter, err)
	}

	WriteSuccess(responseWriter, nil)
}

func transactionHistory(responseWriter gohttp.ResponseWriter, httpReq *gohttp.Request) {
	userIDStr := httpReq.URL.Query().Get("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		log.Println(fmt.Sprintf("[transactionHistory][ParseInt][error:%+v]", err))
		WriteError(responseWriter, err)
		return
	}

	resp, err := disbursement.TransactionHistory(userID)
	if err != nil {
		log.Println(fmt.Sprintf("[disburse][Disburse][error:%+v]", err))
		WriteError(responseWriter, err)
	}

	WriteSuccess(responseWriter, resp)
}
