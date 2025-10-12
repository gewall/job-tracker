package config

import (
	models "job-tracker/internal/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnect struct {
	Config *Config
}

func NewConnect(c *Config) *DBConnect{
	return &DBConnect{
		Config: c,
	}
}

func (c DBConnect) Connect() *gorm.DB {
	dsn := "host=localhost user=dev password=dev dbname=dev_job port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("fail to connect database: %s", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("fail to connect database: %s", err)
	}

	if err := sqlDb.Ping(); err != nil{
		log.Fatalf("database ping failed: %s",err)
	}
	// db.Migrator().DropTable(&models.User{}, &models.EmailVerification{}, &models.RefreshToken{})


	if err := db.AutoMigrate(&models.User{},&models.EmailVerification{},&models.RefreshToken{},&models.Application{}); err != nil {
		log.Fatalf("fail to migrate databse models: %s", err)
	}

	return db
}