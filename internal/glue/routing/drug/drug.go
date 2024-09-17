package drug

import (
	"net/http"
	"pharma-backend/internal/glue/routing"
	"pharma-backend/internal/handler/rest"

	"github.com/gin-gonic/gin"
)

func InitRoute(grp *gin.RouterGroup, drug rest.Drug) {

	drugs := grp.Group("drugs")
	drugRoutes := []routing.Router{
		{
			Method:  http.MethodPost,
			Handler: drug.CreateDrug,
		},
	}

	routing.RegisterRoutes(drugs, drugRoutes)

}
