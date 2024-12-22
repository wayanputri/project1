package main

import (
	"project1/app/config"
	"project1/app/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	// mysql := database.InitMysql(cfg)
	database.InitMysql(cfg)
	// database.InitialMigration(mysql)
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	// routes.InitRouter(e, mysql)
	e.Logger.Fatal(e.Start(":80"))

}
