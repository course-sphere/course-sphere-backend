package http

import (
	"context"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	"github.com/course-sphere/course-sphere-backend/shared/transports/http/middleware"
)

func (s *Server) RegisterRoutes(f *fuego.Server) {
	ctx := context.Background()

	fuego.Get(f, "/", s.Ping)

	jwks := s.AuthClient.MustGetJwks(ctx)
	tokenMiddleware := middleware.RequireToken(jwks)

	course := fuego.Group(f, "/course",
		option.Middleware(tokenMiddleware),
		option.Security(openapi3.SecurityRequirement{"bearerAuth": []string{}}),
	)
	fuego.Post(course, "/", s.CreateCourse)
	fuego.Get(course, "/{id}", s.GetCourse)
	fuego.Patch(course, "/{id}", s.UpdateCourse)
	fuego.Delete(course, "/{id}", s.UpdateCourse)
	fuego.Post(course, "/{id}/material", s.CreateMaterial)
	fuego.Get(course, "/{id}/material", s.GetMaterials)

	material := fuego.Group(f, "/material",
		option.Middleware(tokenMiddleware),
		option.Security(openapi3.SecurityRequirement{"bearerAuth": []string{}}),
	)
	fuego.Patch(material, "/{id}", s.UpdateMaterial)
	fuego.Delete(material, "/{id}", s.DeleteMaterial)
}

func (s *Server) OpenAPI(specURL string) http.Handler {
	return httpSwagger.Handler(
		httpSwagger.Layout(httpSwagger.StandaloneLayout),
		httpSwagger.PersistAuthorization(true),
		httpSwagger.URL(specURL),
	)
}

func (s *Server) Ping(c fuego.ContextNoBody) (string, error) {
	return "pong", nil
}
