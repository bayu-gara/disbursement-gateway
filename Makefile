pkgs = $(shell go list ./... | grep -v vendor | grep -v mock)

gorun:
	@go build -v -o disbursement-gateway main.go && \
	./disbursement-gateway

test:
	@echo "RUN TESTING..."
	@go test -v -cover -gcflags=-l $(pkgs) -coverprofile coverage.out