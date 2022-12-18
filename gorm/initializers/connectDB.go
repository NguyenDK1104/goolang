package initializers

import (
	"fmt"
	"go-jwt/gorm/utils"
	"log"
	// "github.com/emirpasic/gods/utils"
	oracle "github.com/wdrabbit/gorm-oracle"
	"gorm.io/gorm"
)

func ConnectDb() (db *gorm.DB) {
	config, errors := utils.LoadConfig("./")
	if errors != nil {
		log.Fatal("cannot load config:", errors)
	}
	dsn := "oracle://" + config.DBUserName + ":" + config.DBPassword + "@" + config.DBAddress + ":" + config.DBPort + "/" + config.DBServiceName
	db, err := gorm.Open(oracle.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("connect err")
		return
	}
	return db
}
