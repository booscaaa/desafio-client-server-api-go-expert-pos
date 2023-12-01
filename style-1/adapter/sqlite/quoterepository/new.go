package quoterepository

import (
	"github.com/desafio-client-server-api-go-expert-pos/style-1/core/domain"
	"gorm.io/gorm"
)

type repository struct {
	database *gorm.DB
}

func New(database *gorm.DB) domain.QuoteRepository {
	return &repository{
		database: database,
	}
}
