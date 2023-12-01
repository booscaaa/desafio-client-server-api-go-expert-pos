package sqlite

import (
	"github.com/desafio-client-server-api-go-expert-pos/style-1/core/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data/desafio-client-server.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&domain.Quote{})

	return db
}
