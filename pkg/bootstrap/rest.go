package bootstrap

import (
	"net/http"
	"sync"
)

var (
	client         *http.Client
	restClientOnce sync.Once
)

func InitRESTClient() {
	restClientOnce.Do(func() {
		client = &http.Client{}
	})
}

func GetRESTClient() *http.Client {
	return client
}
