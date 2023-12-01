package di

import (
	"net/http"

	"github.com/desafio-client-server-api-go-expert-pos/style-1/adapter/http/rest/quotehttpservice"
	"github.com/desafio-client-server-api-go-expert-pos/style-1/adapter/httpclient/quotehttpclient"
	"github.com/desafio-client-server-api-go-expert-pos/style-1/adapter/sqlite/quoterepository"
	"github.com/desafio-client-server-api-go-expert-pos/style-1/core/domain"
	"github.com/desafio-client-server-api-go-expert-pos/style-1/core/usecase/quoteusecase"
	"gorm.io/gorm"
)

func ConfigQuoteDIService(database *gorm.DB, client *http.Client) domain.QuoteHTTPService {
	quoteRepository := quoterepository.New(database)
	quoteHTTPClient := quotehttpclient.New(client)
	quoteUseCase := quoteusecase.New(quoteRepository, quoteHTTPClient)
	quoteHTTPService := quotehttpservice.New(quoteUseCase)

	return quoteHTTPService
}
