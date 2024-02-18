package main

import (
	"Hrms/pkg/configs"
	"Hrms/pkg/logging"
	"Hrms/pkg/middleware"
	"Hrms/pkg/routes/api"
	"Hrms/pkg/routes/web"
	"Hrms/pkg/utils"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {

	config := configs.FiberConfig()

	app := fiber.New(config)
	middleware.FiberMiddleware(app)

	api.PrivateRoutes(app)
	api.PublicRoutes(app)
	web.PrivateRoutes(app)
	web.PublicRoutes(app)

	logging.Config(app)

	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}

}
