package domain

import "io"

type UploadFileData struct {
	Filename    string
	ContentType string
	File        io.Reader
}
