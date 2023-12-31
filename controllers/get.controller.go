package controllers

import (
	"encoding/json"
	"fiber_gorm/services"
	"strings"
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
		req.Header.Add("", "sd")
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
		t, err := token.SignedString([]byte("a"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"token": t})
	}

}

func GetAuthentication() fiber.Handler {

	return func(c *fiber.Ctx) error {
		username := c.Locals("username").(string)

		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"data": "Authorization accept", "name": username})

	}
}

func InsertOnetoOne() fiber.Handler {

	return func(c *fiber.Ctx) error {

		err := services.InsertOnetoOne(c)
		if err != nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"data": "Insert Doneß accept"})

	}
}
func QueryOnetoOne() fiber.Handler {

	return func(c *fiber.Ctx) error {

		data, err := services.GetQueryOnetoOne(c)
		if err != nil {
			return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"data": data})

	}
}

func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing Authorization Header",
			})
		}

		splitToken := strings.Split(authHeader, "Bearer ")

		if len(splitToken) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid Token Format",
			})
		}
		tokenString := splitToken[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("a"), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid Token",
			})
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to parse claims",
			})
		}

		name := claims["name"].(string)

		c.Locals("username", name)

		return c.Next()
	}
}
