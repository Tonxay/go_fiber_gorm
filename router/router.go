package router

import (
	"fiber_gorm/controllers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func setUpRouter() {
	app := fiber.New()

	postRouter := app.Group("/get")
	postRouter.Post("create/:name/:id", controllers.PostGetAll())
	postRouter.Get("user", controllers.GetAll())
	postRouter.Post("update/:name/:id", controllers.Update())
	postRouter.Delete("delete/:id", controllers.Delete())
	postRouter.Get("getotherPAI", controllers.GetOtherAPI())
	log.Fatal(app.Listen(":8082"))

}
func Init() {
	setUpRouter()
}
