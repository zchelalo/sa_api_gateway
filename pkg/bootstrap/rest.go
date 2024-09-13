package bootstrap

import (
	"net/http"
	"sync"
)

var (
	client *http.Client
	once   sync.Once
)

func InitRESTClient() {
	once.Do(func() {
		client = &http.Client{}
	})
}

func GetRESTClient() *http.Client {
	return client
}
