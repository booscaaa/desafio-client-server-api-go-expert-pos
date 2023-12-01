package quotehttpclient

import (
	"net/http"

	"github.com/desafio-client-server-api-go-expert-pos/style-1/core/domain"
)

type httpclient struct {
	client *http.Client
}

func New(client *http.Client) domain.QuoteHTTPClient {
	return &httpclient{
		client,
	}
}
