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
func configDbPostgres() error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("PGUSER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))
	dsn := psqlInfo
	db1, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error : Fail to connect to Database")
	}
	db = db1
	// g := gen.NewGenerator(gen.Config{
	// 	OutPath: "./query",
	// 	Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	// })

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	//g.UseDB(db) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions

	// g.ApplyBasic(
	// 	// Generate struct `User` based on table `users`
	// 	g.GenerateModel("users"),

	// 	// Generate struct `Employee` based on table `users`
	// 	g.GenerateModelAs("users", "Employee"),

	// 	// Generate struct `User` based on table `users` and generating options
	// 	g.GenerateModel("users", gen.FieldIgnore("address"), gen.FieldType("id", "int64")),
	// )
	// g.ApplyBasic(
	// 	// Generate structs from all tables of current database
	// 	g.GenerateAllTable()...,
	// )
	// Generate the code
	// g.Execute()
	return nil
}

func InitDB() {
	configDbPostgres()
}
