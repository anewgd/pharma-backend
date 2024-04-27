package api

import (
	"github.com/anewgd/pharma_backend/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (usrHandler *UserHandler) createUser(ctx *gin.Context) {
	req := service.CreateUserRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := usrHandler.userService.CreateUser(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, res)

}

func (UserHandler *UserHandler) loginUser(ctx *gin.Context) {
	req := service.LoginUserRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := UserHandler.userService.LoginUser(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, res)

}
