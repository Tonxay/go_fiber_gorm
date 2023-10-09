package controllers

import (
	"fiber_gorm/services"

	"github.com/gofiber/fiber/v2"
)
func PostGetAll()fiber.Handler {

   return func(c *fiber.Ctx) error {
	       user,err := services.PostCreate()
	    if err != nil{
			 		return c.Status(fiber.ErrBadRequest.Code).JSON(err)
		}
		return c.Status(fiber.StatusOK).JSON(user)
	}
}