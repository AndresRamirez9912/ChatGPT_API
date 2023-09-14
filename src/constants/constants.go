package constants

const (
	CONTENT_TYPE        = "Content-Type"
	APPLICATION_JSON    = "application/json"
	MULTIPART_FORM_DATA = "multipart/form-data"
	AUTHORIZATION       = "Authorization"
	BEARER              = "Bearer"
)

const (
	NO_RESPONSE = "ChatGPT didn't response"
)

const (
	URL_COMPLETION          = "https://api.openai.com/v1/chat/completions"
	URL_FINE_TUNING         = "https://api.openai.com/v1/fine_tuning/jobs"
	URL_UPLOAD_FILE         = "https://api.openai.com/v1/files"
	URL_LIST_UPLOADED_FILES = "https://api.openai.com/v1/files"
)

const (
	FINE_TUNING_PURPOSE = "fine-tune"
)
