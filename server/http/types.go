package http

import (
	"github.com/bayu-gara/disbursement-gateway/common/status"
)

type ResponseHTTP struct {
	Status status.Status `json:"status"`
	Data   interface{}   `json:"data,omitempty"`
}
