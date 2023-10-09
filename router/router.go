package router

import (
	"fiber_gorm/controllers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func setUpRouter(){
    app := fiber.New()
	
	postRouter := app.Group("/get")
	postRouter.Get("",controllers.PostGetAll())

	log.Fatal(app.Listen(":8082"))
   
}
func Init() {
	setUpRouter();
}