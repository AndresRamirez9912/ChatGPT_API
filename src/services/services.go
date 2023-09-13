package services

import (
	"GoGPT/src/constants"
	"GoGPT/src/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func SendPromp(promp string) string {
	// Create the requests' body
	body := models.RequestBody{
		Model: "gpt-3.5-turbo",
		Message: []models.Message{
			{
				Role:    "user",
				Content: promp,
			},
		},
		Temperature: 0.7,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Println("Error serialization the request ", err)
	}

	// Create request
	req, err := http.NewRequest("POST", constants.COMPLETION_URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Println("Error creating the request ", err)
	}

	// Set the heaers
	req.Header.Add(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	req.Header.Add(constants.AUTHORIZATION, "Bearer "+models.GptClient.Api_Key)

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

	for _, v := range response.Choices {
		return v.Message.Content
	}

	return constants.NO_RESPONSE
}

func TrainModel() {

}
