package db

import (
	"fmt"

	"github.com/markovidakovic/surl/server/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg config.Config) *gorm.DB {
	conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort, cfg.DbSslMode)
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic("cannot connect the db")
	}
	fmt.Println("Database connected")
	return db
}
