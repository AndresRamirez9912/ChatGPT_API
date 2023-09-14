package models

type RequestBody struct {
	Model       string    `json:"model"`
	Message     []Message `json:"messages"`
	Temperature float32   `json:"temperature"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type GPTResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int32    `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

type FileUploadBody struct {
	Purpose string `json:"purpose"`
	File    string `json:"file"`
}

type FileUploadResponse struct {
	Id        string `json:"id"`
	Object    string `json:"object"`
	Bytes     int    `json:"bytes"`
	CreatedAt int32  `json:"created_at"`
	FileName  string `json:"fileName"`
	Purpose   string `json:"purpose"`
	Status    string `json:"status"`
}

type ListFilesResponse struct {
	Data   []FileUploadResponse `json:"data"`
	Object string               `json:"object"`
}

type CreateFineTuningBody struct {
	Training_file string `json:"training_file"`
	Model         string `json:"model"`
}

type CreateFineTuningResponse struct {
	Object         string   `json:"object"`
	ID             string   `json:"id"`
	Model          string   `json:"model"`
	CreatedAt      int64    `json:"created_at"`
	FineTunedModel string   `json:"fine_tuned_model"`
	OrganizationID string   `json:"organization_id"`
	ResultFiles    []string `json:"result_files"`
	Status         string   `json:"status"`
	ValidationFile string   `json:"validation_file"`
	TrainingFile   string   `json:"training_file"`
}
