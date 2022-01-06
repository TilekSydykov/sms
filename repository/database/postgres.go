package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"solar-faza/entity/util"
	"solar-faza/env"
	"time"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func Connect() *gorm.DB {
	var config = &gorm.Config{}
	if env.GetEnvVars().IsDBDebugEnabled {
		config.Logger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{SlowThreshold: time.Second, IgnoreRecordNotFoundError: true, LogLevel: logger.Info},
		)
	}
	db, err := gorm.Open(postgres.Open(env.GetEnvVars().DbUri), config)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func SetupDB(db *gorm.DB) {
	err := db.AutoMigrate(util.Entities()...)
	if err != nil {
		fmt.Println(err)
		panic("Can't migrate database")
	}
}
