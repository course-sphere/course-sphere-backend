package http

import (
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/course-sphere/course-sphere-backend/pkg/middleware"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/config"
	"github.com/course-sphere/course-sphere-backend/services/general/internal/usecase"
	"github.com/course-sphere/course-sphere-backend/shared/ports"
)

type Handler struct {
	Config     *config.Config
	Course     *usecase.Course
	Presigner  *s3.PresignClient
	AuthClient ports.AuthClient
	UserClient ports.UserClient
}

func (h *Handler) RegisterRoutes(s *fuego.Server) {
	fuego.Get(s, "/", h.Ping)

	course := fuego.Group(s, "/course",
		option.Middleware(middleware.MustToken(h.AuthClient)),
		option.Security(openapi3.SecurityRequirement{"bearerAuth": []string{}}),
	)
	fuego.Post(course, "/", h.CreateCourse)
	fuego.Get(course, "/{id}", h.GetCourse)
	fuego.Patch(course, "/{id}", h.UpdateCourse)
	fuego.Delete(course, "/{id}", h.UpdateCourse)
	fuego.Post(course, "/{id}/material", h.CreateMaterial)
	fuego.Get(course, "/{id}/material", h.GetMaterials)

	material := fuego.Group(s, "/material",
		option.Middleware(middleware.MustToken(h.AuthClient)),
		option.Security(openapi3.SecurityRequirement{"bearerAuth": []string{}}),
	)
	fuego.Patch(material, "/{id}", h.UpdateMaterial)
	fuego.Delete(material, "/{id}", h.DeleteMaterial)

	storage := fuego.Group(s, "/storage")
	fuego.Post(storage, "/create-presigned-url", h.CreateS3PresignedURL)
}

func (h *Handler) OpenAPI(specURL string) http.Handler {
	return httpSwagger.Handler(
		httpSwagger.Layout(httpSwagger.StandaloneLayout),
		httpSwagger.PersistAuthorization(true),
		httpSwagger.URL(specURL),
	)
}
