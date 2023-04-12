package integration

import (
	"net/http"
	"time"
)

func NewClient() *http.Client {
	return &http.Client{Timeout: 15 * time.Second}
}
