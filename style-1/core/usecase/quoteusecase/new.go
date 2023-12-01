package quoteusecase

import "github.com/booscaaa/desafio-client-server-api-go-expert-pos/style-1/core/domain"

type usecase struct {
	repository      domain.QuoteRepository
	quoteHTTPClient domain.QuoteHTTPClient
}

func New(
	repository domain.QuoteRepository,
	quoteHTTPClient domain.QuoteHTTPClient,
) domain.QuoteUseCase {
	return &usecase{
		repository:      repository,
		quoteHTTPClient: quoteHTTPClient,
	}
}
