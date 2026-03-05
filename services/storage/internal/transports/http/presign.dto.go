package http

type CreatePresignedRequestData struct {
	FileName    string `json:"fileName"`
	ContentType string `json:"contentType"`
}

type PresignedRequest struct {
	URL    string            `json:"url"`
	Values map[string]string `json:"values"`
}
