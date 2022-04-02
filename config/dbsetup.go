package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
	"time"
)

type GormClient struct {
	DB *gorm.DB
}

var gormClient *GormClient

func ProvideDB() (*GormClient, error) {
	db, err := gorm.Open("mysql", "root:Pa55word@tcp(localhost:3306)/kumparan_test_db")
	if err != nil {
		log.Panic().Err(err).Msg("fail to initialize database")
		return nil, err
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)
	db.LogMode(true)

	log.Info().Msg("database connected")
	gormClient = &GormClient{DB: db}
	return gormClient, nil
}

func GetGormClient() *GormClient {
	return gormClient
}
