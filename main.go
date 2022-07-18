package main

import (
	"flag"
	"fmt"
	"github.com/bpatyi/gopoc/app"
	"github.com/bpatyi/gopoc/config"
	"github.com/bpatyi/gopoc/controllers"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/labstack/echo/v4/middleware"

	fiber "github.com/gofiber/fiber/v2"
	echo "github.com/labstack/echo/v4"
)

var configPath = flag.String("config", "config.json", "Path to configuration file")
var configEnvPrefix = flag.String("config_env_prefix", "gopoc", "Environment prefix of configs")

func main() {
	flag.Parse()

	application, err := app.NewApp(configPath, configEnvPrefix)
	if err != nil {
		panic(fmt.Errorf("Failed to create application: %w", err))
	}

	if application.Config.Framework == config.FrameworkGin {
		app := gin.New()
		app.Use(gin.Logger())
		app.Use(gin.Recovery())
		controllers.GinRoute(app)
		app.Run(":9000")
	} else if application.Config.Framework == config.FrameworkEcho {
		app := echo.New()
		app.Use(middleware.Logger())
		app.Use(middleware.Recover())
		controllers.RouteEcho(app)
		app.Logger.Fatal(app.Start(":9000"))
	} else if application.Config.Framework == config.FrameworkFiber {
		app := fiber.New()
		app.Use(recover.New())
		app.Use(logger.New())
		controllers.RouteFiber(app)
		app.Listen(":9000")
	}

}
