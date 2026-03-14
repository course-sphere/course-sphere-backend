package http

import (
	"fmt"
	"net/http"

	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func (s *Server) RegisterRoutes(f *fuego.Server) {
	fuego.Post(f, "/upload", s.Upload, option.RequestContentType("multipart/form-data"))
	fuego.Post(f, "/upload/{id}", s.Upload, option.RequestContentType("multipart/form-data"))
	fuego.Get(f, "/{key}", s.Get)

	presign := fuego.Group(f, "/presign")
	fuego.Post(presign, "/", s.CreatePresignedRequest)
}

func (s *Server) OpenAPI(specURL string) http.Handler {
	return httpSwagger.Handler(
		httpSwagger.Layout(httpSwagger.StandaloneLayout),
		httpSwagger.PersistAuthorization(true),
		httpSwagger.URL(fmt.Sprintf("/storage%s", specURL)),
	)
}
