package api

import (
	"net/http"

	"github.com/anewgd/pharma_backend/service"
	"github.com/anewgd/pharma_backend/util"
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
	c := ctx.Request.Context()
	req := service.CreateUserRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(util.NewErrorResponse(util.RequestError.New("malformed request body"), http.StatusBadRequest, "invalid request"))
		return
	}

	res, err := usrHandler.userService.CreateUser(c, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, res)

}

func (UserHandler *UserHandler) loginUser(ctx *gin.Context) {
	req := service.LoginUserRequest{}

	c := ctx.Request.Context()
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(util.NewErrorResponse(util.RequestError.New("malformed request body"), http.StatusBadRequest, "invalid request"))
		return
	}

	res, err := UserHandler.userService.LoginUser(c, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)

}
