package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/booscaaa/desafio-client-server-api-go-expert-pos/style-1/di"
	"gorm.io/gorm"
)

func InitializeRest(serverPort string, database *gorm.DB, client *http.Client) {
	quoteHTTPService := di.ConfigQuoteDIService(database, client)

	mux := http.NewServeMux()

	mux.HandleFunc("/cotacao", quoteHTTPService.Get)

	log.Println(fmt.Sprintf("Rodando o servidor na porta: %s", serverPort))
	http.ListenAndServe(fmt.Sprintf(":%s", serverPort), mux)
}
