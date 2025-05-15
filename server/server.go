package server

import (
	"twitter-uala/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type HTTPServer struct {
	engine   *gin.Engine
	validate *validator.Validate
	services *domain.Services
}

func NewHTTPServer(engine *gin.Engine, services *domain.Services, validate *validator.Validate) *HTTPServer {
	server := &HTTPServer{
		engine:   engine,
		validate: validate,
		services: services,
	}
	server.registerRoutes()
	return server
}

func (s *HTTPServer) registerRoutes() {
	s.engine.POST(":userId/tweet", s.)
}

func (s *HTTPServer) Start() error {
	return s.engine.Run(":8080")
}
