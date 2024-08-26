package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/lfcifuentes/online-payment-platform/api/app/cmd"
	"github.com/spf13/viper"
)

// @title			Transaction API
// @version		1.0
// @description	This is a simple Bank Simulator API documentation.
// @termsOfService	http://swagger.io/terms/
// @contact.name	Luis Cifuentes
// @contact.url	https://lfcifuentes.netlify.app
// @contact.email	lfcifuentes28@gmail.com
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host
// @Accept		json
// @Produce	json
// @BasePath	/
func main() {
	// Create a new logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	// Show the config file path
	slog.Info("Starting application", "event", "app_starting")

	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	// Configuring the viper to read the environment variables
	viper.AutomaticEnv()
	// Show the config file path
	slog.Info("Using config file", "event", "app_env_file_load", "file", ".env")

	cmd.Execute()
}
