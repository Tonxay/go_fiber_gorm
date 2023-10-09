package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

 // Get the Database instance via this function
func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Error : Database instance is not instanciated yet")
	}
	return db
}
// connect db
func configDbPostgres()error{
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("PORT"),os.Getenv("PGUSER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
	dsn := psqlInfo
    db1 , err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
	  log.Fatal("Error : Fail to connect to Database")
	}
	db=db1
	return nil
}

func InitDB(){
	configDbPostgres()
}