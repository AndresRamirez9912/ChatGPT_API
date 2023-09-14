package utils

import (
	"GoGPT/src/constants"
	"GoGPT/src/models"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateSendRequest(method string, url string, jsonBody []byte) ([]byte, error) {
	// Create request
	req, err := http.NewRequest(method, url, bytes.NewReader(jsonBody))
	if err != nil {
		log.Println("Error creating the request ", err)
		return nil, err
	}

	// Set the heaers
	req.Header.Add(constants.CONTENT_TYPE, constants.APPLICATION_JSON)
	req.Header.Add(constants.AUTHORIZATION, "Bearer "+models.GptClient.Api_Key)

	// Send Request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error sending the request ", err)
		return nil, err
	}
	defer res.Body.Close()

	bodyResponse, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading the body response ", err)
		return nil, err
	}
	return bodyResponse, nil
}
