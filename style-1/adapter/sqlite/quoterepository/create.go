package quoterepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/booscaaa/desafio-client-server-api-go-expert-pos/style-1/core/domain"
)

func (repository repository) Create(ctx context.Context, quote *domain.Quote) (*domain.Quote, error) {
	ctxDb, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	err := repository.database.WithContext(ctxDb).Create(quote).Error

	if errors.Is(err, context.DeadlineExceeded) {
		log.Println("Context Error:", err)
	}

	if err != nil {
		return nil, err
	}

	return quote, nil
}
