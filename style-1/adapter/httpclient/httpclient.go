package httpclient

import "net/http"

func InitializeHttpClient() *http.Client {
	return http.DefaultClient
}
