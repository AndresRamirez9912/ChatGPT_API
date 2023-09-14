package main

import (
	"GoGPT/src/models"
	"GoGPT/src/services"
	"fmt"
	"log"
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
	response, err := services.SendPromp("ft:gpt-3.5-turbo-0613:personal::7yjtezD4", "Hola")
	if err != nil {
		log.Fatal("An Error occours chatting with the IA")
	}
	fmt.Println(response)

	// Upload training file
	// services.UploadFile("train.jsonl")

	// List Files uploaded
	// files, err := services.ListFiles()
	// if err != nil {
	// 	log.Fatal("An Error occours listing the uplaoded files")
	// }
	// fmt.Println(files)

	// Create fine Tuning job
	// info, err := services.GetFineTuningInfo("ftjob-mgkhRzneSBjHnk2rIadcm7CR")
	// if err != nil {
	// 	log.Fatal("An Error occours getting the information of the Fine tunning")
	// }
	// fmt.Println(info)
}
