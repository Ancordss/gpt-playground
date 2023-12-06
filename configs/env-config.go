package configs

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Cargar variables de entorno desde el archivo .env
	if err := godotenv.Load("configs/.env"); err != nil {
		fmt.Println("could not load config from env file")
	}
}

func SetupConfig() {
	// Cargar variables de entorno desde el archivo .env
	LoadEnv()

	// Obtener el valor de una variable de entorno
	mode := os.Getenv("GIN_MODE")

	// Verificar si la variable de entorno está definida

	if mode == "debug" {
		fmt.Println("set debug mode")
		gin.SetMode(gin.DebugMode)
	}
}
