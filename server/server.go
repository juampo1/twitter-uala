package server

import (
	"fmt"
	"net/http"
	"twitter-uala/internal/domain"
	"twitter-uala/server/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type HTTPServer struct {
	engine   *gin.Engine
	validate *validator.Validate
	services *domain.Services
}

func (s *HTTPServer) Run(port string) {
	if err := s.engine.Run(port); err != nil {
		panic(err)
	}
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
	s.engine.POST(":userId/tweet", s.CreateTweet)
}

func (s *HTTPServer) CreateTweet(c *gin.Context) {
	var tweetRequest dto.UserTweetRequest
	userID, _ := c.Get("userId")
	id := userID.(string)

	if err := c.ShouldBind(&tweetRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	if err := s.validate.Struct(tweetRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Validation error: %s", err.Error())})
		return
	}

	err := s.services.UserService.CreateTweet(c.Request.Context(), tweetRequest.Content, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Tweet created successfully"})
}
