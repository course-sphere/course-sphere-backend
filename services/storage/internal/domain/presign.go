package domain

type CreatePresignedRequestData struct {
	FileName    string
	ContentType string
}

type PresignedRequest struct {
	URL    string
	Values map[string]string
}
