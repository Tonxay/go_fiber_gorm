package controllers

import (
	"encoding/json"
	"fiber_gorm/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type Abser interface {
}

func PostGetAll() fiber.Handler {

	return func(c *fiber.Ctx) error {
		name := c.Params("name")
		user, err := services.PostCreate(c)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": user, "name": name})
	}
}
func GetAll() fiber.Handler {

	return func(c *fiber.Ctx) error {
		user, err := services.GetRead(c)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": user})
	}
}
func Update() fiber.Handler {

	return func(c *fiber.Ctx) error {
		user, err := services.UpDate(c)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": user})
	}
}
func Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {

		intVar, err := services.Delete(c)
		if err != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(err.Error())
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": intVar})
	}
}

func GetOtherAPI() fiber.Handler {

	return func(c *fiber.Ctx) error {

		a := fiber.AcquireAgent()
		req := a.Request()
		req.Header.SetMethod("GET")
		req.SetRequestURI("https://api.publicapis.org/entries")

		if err := a.Parse(); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		code, body, errs := a.Bytes()

		if errs != nil {
			return c.Status(code).JSON(errs)
		}
		var something fiber.Map
		err := json.Unmarshal(body, &something)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"err": err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"test": "test", "data": something})
	}
}

func PostBodyData() fiber.Handler {
	var data fiber.Map
	return func(c *fiber.Ctx) error {
		if err := c.BodyParser(&data); err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": data})
	}
}
func PostBodyFromFile() fiber.Handler {

	return func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"err": err,
			})
		}
		err1 := c.SaveFile(file, "./"+file.Filename)
		if err1 != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"err": err1,
			})
		}

		return c.SendFile("./" + file.Filename)
	}
}

func Getimage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		name := c.Params("name")
		return c.SendFile("./" + name)
	}
}

func Login() fiber.Handler {

	return func(c *fiber.Ctx) error {

		user := c.FormValue("user")
		pass := c.FormValue("pass")
		// Throws Unauthorized error
		if user != "john" || pass != "doe" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		// Create the Claims
		claims := jwt.MapClaims{
			"name":  "John Doe",
			"admin": true,
			"exp":   time.Now().Add(time.Hour * 72).Unix(),
		}

		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiQWRtaW5zZGZuc25kbW5tZ2xrbmZham9lcm9waWphW2hqaSIsIklzc3VlciI6Iklzc3VlcmZrbWdbb2RmaFtqYXNmZGpbaG9lckgiLCJVc2VybmFtZSI6IlNGREdKYXZhSW5Vc2VzZHNkZmtrc2Rma24iLCJleHAiOjE2OTcwODYzNDIsImlhdCI6MTY5NzA4NjM0Mn0.CFW865hEWeKy7VFM2PQxMREMuiX6X2_4-JCgQK92KNw"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"token": t})
	}

}
