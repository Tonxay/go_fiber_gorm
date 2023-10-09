package main

import (
	"fiber_gorm/db"
	"fiber_gorm/db/dotenv"
	"fiber_gorm/router"
)
  
func main() {
   dotenv.SetDotenv()
   db.InitDB()
   router.Init() 
}   