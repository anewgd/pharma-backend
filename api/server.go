package api

import (
	"context"
	"net/http"

	"github.com/anewgd/pharma_backend/token"
	"github.com/anewgd/pharma_backend/util"
	"github.com/gin-gonic/gin"
	"github.com/joomcode/errorx"
)

type Server struct {
	router          *gin.Engine
	drugHandler     *DrugHandler
	userHandler     *UserHandler
	pharmacyHandler *PharmacyHandler
}

func NewServer(drugHandler *DrugHandler, userHandler *UserHandler, pharmacyHandler *PharmacyHandler) *Server {

	server := &Server{
		drugHandler:     drugHandler,
		userHandler:     userHandler,
		pharmacyHandler: pharmacyHandler,
	}
	server.router = server.setupRoutes()
	return server
}

func (s *Server) setupRoutes() *gin.Engine {

	router := gin.Default()
	router.Use(ErrorHandlerMiddleware())
	router.POST("/users", s.userHandler.createUser)
	router.POST("/login", s.userHandler.loginUser)
	router.POST("/pharmacy/login", s.pharmacyHandler.pharmacyLogin)

	adminRoute := router.Group("/admin").Use(authMiddleware(*s.userHandler.userService.GetTokenMaker()))

	adminRoute.POST("/pharmacies", s.pharmacyHandler.createPharmacy)
	adminRoute.POST("/branches", s.pharmacyHandler.createPharmacyBranch)
	adminRoute.POST("/managers", s.pharmacyHandler.createManager)

	managerRoute := router.Group("/manager").Use(authMiddleware(*s.userHandler.userService.GetTokenMaker()))
	managerRoute.POST("/drugs", s.drugHandler.addDrugHandler).Use(authMiddleware(*s.userHandler.userService.GetTokenMaker()))

	// pharmacistRoute := router.Group("/pharmacy").Use(authMiddleware(*s.userHandler.userService.GetTokenMaker()))

	return router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func getContextWithValues(ctx *gin.Context) (context.Context, error) {

	c := ctx.Request.Context()
	payload, ok := ctx.Get(util.AuthorizationPayloadKey)
	if !ok {
		return nil, util.NewErrorResponse(util.AuthorizationError.New("no authorization header found"), http.StatusUnauthorized, "no authorization header found")
	}

	usrPayload, ok := (payload).(*token.Payload)
	if !ok {
		return nil, util.NewErrorResponse(errorx.InternalError.New("failed to extract user info from context"), http.StatusInternalServerError, "internal error")
	}

	c = context.WithValue(c, util.UserID, usrPayload.UserID)
	c = context.WithValue(c, util.Role, usrPayload.Role)
	return c, nil
}
