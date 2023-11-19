package main

import (
	"log"

	"github.com/bayu-gara/disbursement-gateway/server/http"
)

func main() {
	err := http.InitHTTPServer()
	if err != nil {
		log.Fatalln(err)
	}
}
