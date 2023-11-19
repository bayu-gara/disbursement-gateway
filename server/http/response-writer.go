package http

import (
	"fmt"
	"github.com/bayu-gara/disbursement-gateway/common/status"
	jsoniter "github.com/json-iterator/go"
	"log"
	gohttp "net/http"
)

func WriteSuccess(responseWriter gohttp.ResponseWriter, dataResponse interface{}) {
	response := ResponseHTTP{
		Status: status.Success,
		Data:   dataResponse,
	}

	dataBytes, err := jsoniter.Marshal(response)
	if err != nil {
		log.Println(fmt.Sprintf("[WriteSuccess][Marshal][error:%+v]", err))
	}

	responseWriter.Write(dataBytes)
}

func WriteError(responseWriter gohttp.ResponseWriter, err error) {
	response := ResponseHTTP{
		Status: status.UnknownError,
		Data:   err,
	}

	dataBytes, err := jsoniter.Marshal(response)
	if err != nil {
		log.Println(fmt.Sprintf("[WriteError][Marshal][error:%+v]", err))
	}

	responseWriter.Write(dataBytes)
}
