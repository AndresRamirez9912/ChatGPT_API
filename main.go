package main

import (
	"GoGPT/src/models"
	"GoGPT/src/services"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env variables
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")

	models.GetGPTClient(apiKey, "Personal")

	response := services.SendPromp("Hola")
	fmt.Println(response)
}
