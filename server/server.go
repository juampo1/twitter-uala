package server

import (
	"fmt"
	"net/http"
	"twitter-uala/internal/domain"
	"twitter-uala/server/dto"
	"twitter-uala/server/middleware"

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
	auth := s.engine.Group("/").Use(middleware.AuthMiddleware())
	{
		auth.POST(":userId/tweet", s.CreateTweet)
		auth.POST(":userId/follow", s.FollowUser)
		auth.GET(":userId/timeline", s.GetTimeline)
	}
}

func (s *HTTPServer) CreateTweet(c *gin.Context) {
	var tweetRequest dto.UserTweetRequest
	id := c.Param("userId")

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

func (s *HTTPServer) FollowUser(c *gin.Context) {
	var followRequest dto.UserFollowRequest
	followerID := c.Param("userId")

	if err := c.ShouldBind(&followRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if followRequest.FollowedID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User to follow ID is required"})
		return
	}

	err := s.services.UserService.FollowUser(c.Request.Context(), followerID, followRequest.FollowedID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User followed successfully"})
}

func (s *HTTPServer) GetTimeline(c *gin.Context) {
	userId := c.Param("userId")

	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	tweets, err := s.services.UserService.GetTimeline(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tweets": tweets})
}
