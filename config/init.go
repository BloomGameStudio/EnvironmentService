package config

import (
	"github.com/BloomGameStudio/EnvironmentService/database"
	models "github.com/BloomGameStudio/EnvironmentService/models/private"
	"github.com/spf13/viper"
)

func Init() {
	ViperInit()

	// Initialize Database
	database.Init()

	// TODO: Cleanup Migration
	db := database.GetDB()

	db.AutoMigrate(&models.Scripts{})
}

func ViperInit() {
	// Initialize Viper with project configuration.

	// Setting Default Values
	viper.SetDefault("DEBUG", true)
	viper.SetDefault("PORT", "1323") // Has to be a string as Echo expects a string
	viper.SetDefault("WS_TIMEOUT_SECONDS", 10)
}
