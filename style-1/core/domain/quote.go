package domain

import (
	"context"
	"net/http"

	"github.com/desafio-client-server-api-go-expert-pos/style-1/core/dto"
)

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

type QuoteRepository interface {
	Create(context.Context, *Quote) (*Quote, error)
}

type QuoteUseCase interface {
	Get(context.Context) (*dto.QuoteOutput, error)
}

type QuoteHTTPService interface {
	Get(http.ResponseWriter, *http.Request)
}

type QuoteHTTPClient interface {
	GetQuote(context.Context, string) (*dto.QuoteClientResponse, error)
}
