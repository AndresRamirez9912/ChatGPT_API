package main

import (
	"GoGPT/src/constants"
	"GoGPT/src/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env variables
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")

	// Create the requests' body
	body := models.RequestBody{
		Model: "gpt-3.5-turbo",
		Message: []models.Message{
			{
				Role:    "user",
				Content: "Hola, como te llamas y de que te encargas?",
			},
		},
		Temperature: 0.7,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Println("Error serialization the request ", err)
	}

	// Create request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Println("Error creating the request ", err)
	}

	// Set the heaers
	req.Header.Add(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	req.Header.Add(constants.AUTHORIZATION, "Bearer "+apiKey)

	// Send Request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error sending the request ", err)
	}
	defer res.Body.Close()

	bodyResponse, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading the body response ", err)
	}

	// Unmarshall the body
	response := &models.GPTResponse{}
	err = json.Unmarshal(bodyResponse, response)
	if err != nil {
		log.Println("Error unmarshalling the body response ", err)
	}

	// Print the response
	fmt.Printf("%+v", response)
}
