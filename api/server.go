package api

import "github.com/gin-gonic/gin"

type Server struct {
	router  *gin.Engine
	handler *DrugHandler
}

func NewServer(drugHandler DrugHandler) *Server {
	server := Server{}
	server.router = server.setupRoutes()
	server.handler = &drugHandler
	return &server
}

func (s *Server) setupRoutes() *gin.Engine {

	router := gin.Default()
	router.POST("/drugs", s.handler.addDrug)

	return router
}
