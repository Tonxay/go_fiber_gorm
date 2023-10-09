package controllers

import (
	"fiber_gorm/services"

	"github.com/gofiber/fiber/v2"
)
func PostGetAll()fiber.Handler {

   return func(c *fiber.Ctx) error {
	      name := c.Params("name")
	       user,err := services.PostCreate(c)
	    if err != nil{
			 	return c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON( fiber.Map{"data": user ,"name":name})
	}
}
func GetAll()fiber.Handler {

   return func(c *fiber.Ctx) error {
	       user,err := services.GetRead(c)
	    if err != nil{
		return c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON( fiber.Map{"data": user })
	}
}
func Update()fiber.Handler {
   return func(c *fiber.Ctx) error {
	       user,err := services.UpDate(c)
	    if err != nil{
		return c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON( fiber.Map{"data": user })
	}
}
func Delete()fiber.Handler {
   return func(c *fiber.Ctx) error {
	       intVar,err := services.Delete(c)
	    if err != nil{
		return c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON( fiber.Map{"data": intVar })
	}
}