package httpclient

import "net/http"

var client *http.Client

func Init() {
	client = &http.Client{}
}

func GetClient() *http.Client {
	return client
}
