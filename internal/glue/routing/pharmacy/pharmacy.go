package pharmacy

import (
	"net/http"
	"pharma-backend/internal/glue/routing"
	"pharma-backend/internal/handler/middleware"
	"pharma-backend/internal/handler/rest"

	"github.com/gin-gonic/gin"
)

func InitRoute(grp *gin.RouterGroup, pharmacy rest.Pharmacy, authMiddleware middleware.AuthenticationMiddleware) {

	pharmacies := grp.Group("pharmacies")
	pharmacyRoutes := []routing.Router{
		{
			Method:  http.MethodPost,
			Handler: pharmacy.CreatePharmacy,
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(),
			},
		},
		{
			Method:  http.MethodPost,
			Handler: pharmacy.CreatePharmacist,
			Path:    "/pharmacist",
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(),
			},
		},
		{
			Method:  http.MethodPost,
			Handler: pharmacy.CreatePharmacyBranch,
			Path:    "/branch",
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(),
			},
		},
		{
			Method:  http.MethodPost,
			Handler: pharmacy.CreatePharmacyBranchManager,
			Path:    "/manager",
			Middlewares: []gin.HandlerFunc{
				authMiddleware.Authentication(),
			},
		},
		{
			Method:      http.MethodPost,
			Handler:     pharmacy.LoginPharmacist,
			Path:        "/login",
			Middlewares: []gin.HandlerFunc{
				// authMiddleware.Authentication(),
			},
		},
	}
	routing.RegisterRoutes(pharmacies, pharmacyRoutes)

}
