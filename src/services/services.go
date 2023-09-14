package services

import (
	"GoGPT/src/constants"
	"GoGPT/src/models"
	"GoGPT/src/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
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

	bodyResponse, err := utils.CreateSendRequest("POST", constants.URL_COMPLETION, jsonBody)
	if err != nil {
		log.Println("Error creating / sending the request", err)
	}

	// Unmarshall the body
	response := &models.GPTResponse{}
	err = json.Unmarshal(bodyResponse, response)
	if err != nil {
		log.Println("Error unmarshalling the body response ", err)
	}

	// Check the response
	for _, v := range response.Choices {
		return v.Message.Content
	}

	return constants.NO_RESPONSE
}

func UploadFile(fileName string) {
	// Open the file to send
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Error opening the file ", err)
	}
	defer file.Close()

	// Create Multipart writer
	buffer := &bytes.Buffer{} // Buffer to store the request body
	writer := multipart.NewWriter(buffer)

	// Create the Fields
	purposeField, err := writer.CreateFormField("purpose")
	if err != nil {
		log.Println("Error creating the purpose field ", err)
	}
	_, err = io.WriteString(purposeField, constants.FINE_TUNING_PURPOSE)
	if err != nil {
		log.Println("Error writing in the purpose field ", err)
	}

	fileBody, err := writer.CreateFormField("file")
	if err != nil {
		log.Println("Error creating the file field ", err)
	}
	_, err = io.WriteString(fileBody, fileName)
	if err != nil {
		log.Println("Error writing in the purpose field ", err)
	}

	// Create the form field
	fileField, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		log.Println("Error creating the fileField ", err)
	}

	// Copy the file content to the form file field
	_, err = io.Copy(fileField, file)
	if err != nil {
		log.Println("Error copying file content ", err)
	}

	// Close the multipar writer
	writer.Close()

	// Create request
	req, err := http.NewRequest("POST", constants.URL_UPLOAD_FILE, buffer)
	if err != nil {
		log.Println("Error creating the request ", err)
	}

	// Set the heaers
	req.Header.Add(constants.CONTENT_TYPE, writer.FormDataContentType())
	req.Header.Add(constants.AUTHORIZATION, "Bearer "+models.GptClient.Api_Key)

	// Send Request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error sending the request ", err)
	}
	defer res.Body.Close()

	response := &models.FileUploadResponse{}
	err = json.NewDecoder(res.Body).Decode(response)
	if err != nil {
		log.Println("Error decoding the response ", err)
		return
	}

	fmt.Println(res.Body)
}

func ListFiles() {
	res, err := utils.CreateSendRequest("GET", constants.URL_LIST_UPLOADED_FILES, nil)
	if err != nil {
		log.Println("Error creating / sending the request ", err)
		return
	}

	response := &models.ListFilesResponse{}
	err = json.Unmarshal(res, response)
	if err != nil {
		log.Println("Error parsing the request request ", err)
		return
	}

	fmt.Printf("%+v\n", response)
}

func CreateFineTuning(tuningFile string) {
	// Create Body request
	body := &models.CreateFineTuningBody{
		Training_file: tuningFile,
		Model:         constants.MODEL_GPT_3_5_TURBO,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Println("Error serialization the request ", err)
	}

	bodyResponse, err := utils.CreateSendRequest("POST", constants.URL_CREATE_FINE_TUNING, jsonBody)
	if err != nil {
		log.Println("Error creating / sending the request", err)
	}

	// Unmarshall the body
	response := &models.CreateFineTuningResponse{}
	err = json.Unmarshal(bodyResponse, response)
	if err != nil {
		log.Println("Error unmarshalling the body response ", err)
	}

	fmt.Println(string(bodyResponse))
	fmt.Printf("%+v", response)
}
