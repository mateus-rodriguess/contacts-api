package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/mateus-rodriguess/contacts-api/pkg/config"
	"github.com/mateus-rodriguess/contacts-api/pkg/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func Connet() {
	p := config.Config("DATABASE_PORT")
	port, err := strconv.ParseInt(p, 10, 32)

	if err != nil {
		log.Fatalln("Error parsing str to int")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.Config("DB_HOST"), config.Config("POSTGRES_USER"), config.Config("POSTGRES_PASSWORD"), config.Config("POSTGRES_DB"), port)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&model.User{})
	DB = Dbinstance{
		Db: db,
	}

}
