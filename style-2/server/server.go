package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type QuoteHTTPResponse struct {
	USDBRL Quote `json:"USDBRL"`
}

type QuoteResponse struct {
	Bid string `json:"bid"`
}

type Quote struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

const (
	SERVER_PORT = "8080"
	BASE_URL    = "https://economia.awesomeapi.com.br/json/last/USD-BRL"
)

func main() {
	database := getDatabaseConnection()

	mux := http.NewServeMux()

	mux.HandleFunc("/cotacao", getQuoteHttp(database))

	log.Println(fmt.Sprintf("Rodando o servidor na porta: %s", SERVER_PORT))
	http.ListenAndServe(fmt.Sprintf(":%s", SERVER_PORT), mux)
}

func getDatabaseConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./desafio-client-server.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(Quote{})

	return db
}

func getQuoteHttp(database *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		quoteFromURL, err := getQuoteFromURL(ctx)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		quoteSaved, err := saveQuoteIntoDatabase(ctx, database, quoteFromURL)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(&QuoteResponse{
			Bid: quoteSaved.Bid,
		})
	}
}

func getQuoteFromURL(ctx context.Context) (*Quote, error) {
	ctxRequest, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	response := QuoteHTTPResponse{}
	req, err := http.NewRequestWithContext(ctxRequest, http.MethodGet, BASE_URL, nil)

	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)

	if errors.Is(err, context.DeadlineExceeded) {
		log.Println("Context Error:", err)
	}

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		return nil, err
	}

	return &response.USDBRL, nil
}

func saveQuoteIntoDatabase(ctx context.Context, database *gorm.DB, quote *Quote) (*Quote, error) {
	ctxDb, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	err := database.WithContext(ctxDb).Create(quote).Error

	if errors.Is(err, context.DeadlineExceeded) {
		log.Println("Context Error:", err)
	}

	if err != nil {
		return nil, err
	}

	return quote, nil
}
