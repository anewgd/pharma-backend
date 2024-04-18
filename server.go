package main

import "github.com/gin-gonic/gin"

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := Server{}
	server.router = server.setupRoutes()
	return &server
}
func addDrug(ctx *gin.Context) {
	ctx.String(200, "test drug")
}

func (s *Server) setupRoutes() *gin.Engine {

	router := gin.Default()
	router.POST("/drugs", addDrug)

	return router
}
