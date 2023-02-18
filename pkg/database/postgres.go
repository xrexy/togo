package database

import (
	"log"

	"github.com/xrexy/togo/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgesClient *gorm.DB

func StartPostgresDB(env *config.EnvVars) error {
	log.Default().Println("Starting Postgres DB")

	var err error
	PostgesClient, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  env.DB_DSN,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Default().Println("Postgres DB started")

	log.Default().Println("Migrating Postgres DB")
	err = PostgesClient.AutoMigrate(&User{}, &Task{})
	if err != nil {
		return err
	}

	log.Default().Println("Postgres DB migrated")

	return nil
}

func ClosePostgresDB() error {
	log.Default().Println("Closing Postgres DB")
	db, err := PostgesClient.DB()
	if err != nil {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}

	log.Default().Println("Postgres DB closed")
	return nil
}
