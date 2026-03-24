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
	authOptions := []fuego.RouteOption{
		option.Middleware(tokenMiddleware),
		option.Security(openapi3.SecurityRequirement{"bearerAuth": []string{}}),
	}

	course := fuego.Group(f, "/course")
	fuego.Get(course, "/", s.GetAllCourses)
	fuego.Get(course, "/my", s.GetEnrolledCourses, authOptions...)
	fuego.Post(course, "/", s.CreateCourse, authOptions...)
	fuego.Get(course, "/{id}", s.GetCourse)
	fuego.Patch(course, "/{id}", s.UpdateCourse, authOptions...)
	fuego.Post(course, "/{id}/material", s.CreateMaterial, authOptions...)
	fuego.Get(course, "/{id}/material", s.GetMaterialsByCourse)
	fuego.Post(course, "/{id}/enroll", s.EnrollCourse, authOptions...)
	fuego.Get(course, "/{id}/progress", s.GetCourseProgress, authOptions...)

	material := fuego.Group(f, "/material")
	fuego.Get(material, "/{id}", s.GetMaterial, option.Hide())
	fuego.Post(material, "/{id}/move", s.MoveMaterial, authOptions...)
	fuego.Patch(material, "/{id}", s.UpdateMaterial, authOptions...)
	fuego.Post(material, "/{id}/question", s.CreateQuestion, authOptions...)
	fuego.Get(material, "/{id}/question", s.GetQuestionsByCourse)
	fuego.Post(material, "/{id}/attempt", s.CreateAttempt, authOptions...)
	fuego.Get(material, "/{id}/attempt", s.GetAttempts, authOptions...)

	question := fuego.Group(f, "/question")
	fuego.Post(question, "/{id}/move", s.MoveQuestion, authOptions...)
	fuego.Patch(question, "/{id}", s.UpdateQuestion, authOptions...)

	attempt := fuego.Group(f, "/attempt")
	fuego.Post(attempt, "/{id}", s.CreateAttemptDetails, authOptions...)
	fuego.Get(attempt, "/{id}", s.GetAttemptDetails, authOptions...)
	fuego.Patch(attempt, "/{id}", s.UpdateAttempt, authOptions...)

	roadmap := fuego.Group(f, "/roadmap")
	fuego.Post(roadmap, "/", s.CreateRoadmap, authOptions...)
	fuego.Post(roadmap, "/{id}", s.AddRoadmapCourse, authOptions...)
	fuego.Post(roadmap, "/{id}/apply", s.ApplyRoadmap, authOptions...)
	fuego.Get(roadmap, "/", s.GetAllRoadmap)
	fuego.Get(roadmap, "/my", s.GetRoadmapsByStudent, authOptions...)
	fuego.Get(roadmap, "/{id}", s.GetRoadmap)
	fuego.Post(roadmap, "/{id}/move-course", s.MoveRoadmapCourse, authOptions...)
	fuego.Patch(roadmap, "/{id}", s.UpdateRoadmap, authOptions...)
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
