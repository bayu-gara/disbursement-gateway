package http

import (
	gohttp "net/http"
)

func InitHTTPServer() error {
	mux := gohttp.NewServeMux()
	registerHandlerHTTP(mux, "/disburse", "POST", disburse)
	registerHandlerHTTP(mux, "/transaction", "GET", transactionHistory)

	return gohttp.ListenAndServe(":8080", mux)
}

func registerHandlerHTTP(
	mux *gohttp.ServeMux,
	path string,
	method string,
	handler func(respWriter gohttp.ResponseWriter, req *gohttp.Request),
) {
	if mux == nil {
		return
	}

	httpMethod := gohttp.MethodGet
	switch method {
	case gohttp.MethodHead:
		httpMethod = gohttp.MethodHead
	case gohttp.MethodPost:
		httpMethod = gohttp.MethodPost
	case gohttp.MethodPut:
		httpMethod = gohttp.MethodPut
	case gohttp.MethodPatch:
		httpMethod = gohttp.MethodPatch
	case gohttp.MethodDelete:
		httpMethod = gohttp.MethodDelete
	case gohttp.MethodConnect:
		httpMethod = gohttp.MethodConnect
	case gohttp.MethodOptions:
		httpMethod = gohttp.MethodOptions
	case gohttp.MethodTrace:
		httpMethod = gohttp.MethodTrace
	}

	mux.HandleFunc(path, func(respWriter gohttp.ResponseWriter, req *gohttp.Request) {
		if req.Method != httpMethod {
			gohttp.Error(respWriter, "Method not allowed", gohttp.StatusMethodNotAllowed)
			return
		}

		handler(respWriter, req)
		respWriter.WriteHeader(gohttp.StatusOK)
	})
}
