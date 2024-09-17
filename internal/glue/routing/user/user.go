package user

import (
	"net/http"
	"pharma-backend/internal/glue/routing"
	"pharma-backend/internal/handler/rest"

	"github.com/gin-gonic/gin"
)

func InitRoute(grp *gin.RouterGroup, user rest.User) {

	users := grp.Group("users")
	userRoutes := []routing.Router{
		{
			Method:  http.MethodPost,
			Handler: user.CreateUser,
		},
		{
			Method:  http.MethodPost,
			Handler: user.LoginUser,
			Path:    "/login",
		},
	}
	routing.RegisterRoutes(users, userRoutes)

}
