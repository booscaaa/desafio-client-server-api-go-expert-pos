package quoteusecase

import (
	"context"

	"github.com/desafio-client-server-api-go-expert-pos/style-1/core/domain"
	"github.com/desafio-client-server-api-go-expert-pos/style-1/core/dto"
)

func (usecase usecase) Get(ctx context.Context) (*dto.QuoteOutput, error) {
	quote, err := usecase.quoteHTTPClient.GetQuote(ctx, "USD-BRL")
	if err != nil {
		return nil, err
	}

	quoteMapped, err := dto.Mapper(quote, domain.Quote{})
	if err != nil {
		return nil, err
	}

	quoteCreated, err := usecase.repository.Create(ctx, quoteMapped)
	if err != nil {
		return nil, err
	}

	quoteMappedOutput, err := dto.Mapper(quoteCreated, dto.QuoteOutput{})
	if err != nil {
		return nil, err
	}
	return quoteMappedOutput, nil
}
