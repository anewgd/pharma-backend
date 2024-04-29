package api

import "github.com/gin-gonic/gin"

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
