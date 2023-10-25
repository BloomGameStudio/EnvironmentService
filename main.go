package main

import (
	"github.com/BloomGameStudio/EnvironmentService/controllers"
	"github.com/BloomGameStudio/EnvironmentService/controllers/ws/scripts"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config")
	viper.AddConfigPath("config/")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	// Enable Echo Debug mode if we are In DEBUG mode. Debug mode sets the log level to DEBUG.
	e.Debug = viper.GetBool("DEBUG")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// HTTP Testing routes
	e.GET("ping", controllers.Ping)
	// End of HTTP testing routes

	// WebSocket Routes
	ws := e.Group("/ws/")

	// Web Socket Testing routes
	ws.File("", "public/index.html") // http://127.0.0.1:1323/ws/
	// ws://localhost:1323/ws
	ws.GET("hello", controllers.Hello)
	ws.GET("ping", controllers.PingWS)
	// End of Web Socket testing routes

	ws.GET("mesh", controllers.MeshWS)
	ws.GET("scripts", scripts.Scripts)

	port := viper.GetString("PORT")
	e.Logger.Fatal(e.Start(":" + port))

}
