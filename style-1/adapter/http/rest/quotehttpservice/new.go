package quotehttpservice

import "github.com/desafio-client-server-api-go-expert-pos/style-1/core/domain"

type service struct {
	usecase domain.QuoteUseCase
}

func New(usecase domain.QuoteUseCase) domain.QuoteHTTPService {
	return &service{
		usecase: usecase,
	}
}
