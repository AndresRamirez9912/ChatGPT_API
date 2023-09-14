package services

import (
	"GoGPT/src/constants"
	"GoGPT/src/models"
	"GoGPT/src/utils"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func SendPromp(model string, promp string) (string, error) {
	// Create the requests' body
	body := models.RequestBody{
		Model: model,
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
		return constants.NO_RESPONSE, err
	}

	bodyResponse, err := utils.CreateSendRequest("POST", constants.URL_COMPLETION, jsonBody)
	if err != nil {
		log.Println("Error creating / sending the request", err)
		return constants.NO_RESPONSE, err
	}

	// Unmarshall the body
	response := &models.GPTResponse{}
	err = json.Unmarshal(bodyResponse, response)
	if err != nil {
		log.Println("Error unmarshalling the body response ", err)
		return constants.NO_RESPONSE, err
	}

	// Check the response
	for _, v := range response.Choices {
		return v.Message.Content, nil
	}

	return constants.NO_RESPONSE, errors.New("There are not responses")
}

func UploadFile(fileName string) (*models.FileUploadResponse, error) {
	// Open the file to send
	file, err := os.Open(fileName)
	if err != nil {
		log.Println("Error opening the file ", err)
		return nil, err
	}
	defer file.Close()

	// Create Multipart writer
	buffer := &bytes.Buffer{} // Buffer to store the request body
	writer := multipart.NewWriter(buffer)

	// Create the Fields
	purposeField, err := writer.CreateFormField("purpose")
	if err != nil {
		log.Println("Error creating the purpose field ", err)
		return nil, err
	}
	_, err = io.WriteString(purposeField, constants.FINE_TUNING_PURPOSE)
	if err != nil {
		log.Println("Error writing in the purpose field ", err)
		return nil, err
	}

	fileBody, err := writer.CreateFormField("file")
	if err != nil {
		log.Println("Error creating the file field ", err)
		return nil, err
	}
	_, err = io.WriteString(fileBody, fileName)
	if err != nil {
		log.Println("Error writing in the purpose field ", err)
		return nil, err
	}

	// Create the form field
	fileField, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		log.Println("Error creating the fileField ", err)
		return nil, err
	}

	// Copy the file content to the form file field
	_, err = io.Copy(fileField, file)
	if err != nil {
		log.Println("Error copying file content ", err)
		return nil, err
	}

	// Close the multipar writer
	writer.Close()

	// Create request
	req, err := http.NewRequest("POST", constants.URL_UPLOAD_FILE, buffer)
	if err != nil {
		log.Println("Error creating the request ", err)
		return nil, err
	}

	// Set the heaers
	req.Header.Add(constants.CONTENT_TYPE, writer.FormDataContentType())
	req.Header.Add(constants.AUTHORIZATION, "Bearer "+models.GptClient.Api_Key)

	// Send Request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error sending the request ", err)
		return nil, err
	}
	defer res.Body.Close()

	response := &models.FileUploadResponse{}
	err = json.NewDecoder(res.Body).Decode(response)
	if err != nil {
		log.Println("Error decoding the response ", err)
		return nil, err
	}

	return response, nil

}

func ListFiles() (*models.ListFilesResponse, error) {
	res, err := utils.CreateSendRequest("GET", constants.URL_LIST_UPLOADED_FILES, nil)
	if err != nil {
		log.Println("Error creating / sending the request ", err)
		return nil, err
	}

	response := &models.ListFilesResponse{}
	err = json.Unmarshal(res, response)
	if err != nil {
		log.Println("Error parsing the request request ", err)
		return nil, err
	}

	return response, nil
}

func CreateFineTuning(tuningFile string) (*models.CreateFineTuningResponse, error) {
	// Create Body request
	body := &models.CreateFineTuningBody{
		Training_file: tuningFile,
		Model:         constants.MODEL_GPT_3_5_TURBO,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Println("Error serialization the request ", err)
		return nil, err
	}

	bodyResponse, err := utils.CreateSendRequest("POST", constants.URL_CREATE_FINE_TUNING, jsonBody)
	if err != nil {
		log.Println("Error creating / sending the request", err)
		return nil, err
	}

	// Unmarshall the body
	response := &models.CreateFineTuningResponse{}
	err = json.Unmarshal(bodyResponse, response)
	if err != nil {
		log.Println("Error unmarshalling the body response ", err)
		return nil, err
	}

	return response, nil
}

func GetFineTuningInfo(id string) (*models.FineTuningInfoResponse, error) {
	// Create the request
	bodyResponse, err := utils.CreateSendRequest("GET", constants.URL_GET_INFO_FINE_TUNING+id, nil)
	if err != nil {
		log.Println("Error creating / sending the request", err)
		return nil, err
	}

	// Unmarshall the body
	response := &models.FineTuningInfoResponse{}
	err = json.Unmarshal(bodyResponse, response)
	if err != nil {
		log.Println("Error unmarshalling the body response ", err)
		return nil, err
	}

	return response, nil
}
