package main

import (
	"GoGPT/src/models"
	"GoGPT/src/services"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env variables
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")

	// Create the client
	models.GetGPTClient(apiKey, "Personal")

	// Send a promp
	// response := services.SendPromp("Hola")
	// fmt.Println(response)

	// Upload training file
	// services.UploadFile("train.jsonl")

	// List Files uploaded
	services.ListFiles()
}
