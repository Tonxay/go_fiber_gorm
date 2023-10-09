package logger

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetLoggerMiddleware(app *fiber.App) {
   
	// My preferrenced Logger config
	myLogger := logger.Config{
		Output:     os.Stdout,
		Format:     `[API] ${cyan}${time} |${yellow}${status}${reset}|${blue}${method}${reset}| IP: ${green}${ip}${reset} | ${magenta}${path} ${reset} ${locals:requestid} | ${error}` + "\n",
		TimeFormat: "2006/01/02 - 15:04:05",
		// TimeZone:   "Asia/Bangkok",
	}
	app.Use(logger.New(myLogger))
}
