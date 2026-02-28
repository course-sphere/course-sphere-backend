package http

type CreatePresignedURLRequest struct {
	FileName    string `json:"fileName"`
	FileType    string `json:"fileType"`
	FileSize    int    `json:"fileSize"`
	ContentType string `json:"contentType"`
}

type PresignedURL struct {
	// Represents the Base URL to make a request to
	URL string `json:"url"`
	// Values is a key-value map of values to be sent as FormData
	// these values are not encoded
	Values map[string]string `json:"values"`
}
