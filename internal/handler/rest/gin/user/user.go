package user

import (
	"context"
	"net/http"
	"time"

	"pharma-backend/internal/constants"
	"pharma-backend/internal/constants/errors"
	"pharma-backend/internal/constants/model/dto"
	"pharma-backend/internal/handler/rest"
	"pharma-backend/internal/module"
	"pharma-backend/platform/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type user struct {
	logger         logger.Logger
	userModule     module.User
	contextTimeout time.Duration
}

func Init(log logger.Logger, userModule module.User, contextTimeout time.Duration) rest.User {

	return &user{
		logger:         log,
		userModule:     userModule,
		contextTimeout: contextTimeout,
	}

}

func (u *user) CreateUser(ctx *gin.Context) {

	cntx, cancel := context.WithTimeout(ctx.Request.Context(), u.contextTimeout)
	defer cancel()

	user := dto.CreateUserRequest{}

	err := ctx.ShouldBind(&user)
	if err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		u.logger.Error(ctx, "unable to bind user data", zap.Error(err))
		_ = ctx.Error(err)
		return
	}

	newUser, err := u.userModule.CreateUser(cntx, user)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	constants.SuccessResponse(ctx, http.StatusCreated, newUser, nil)

}

func (u *user) LoginUser(ctx *gin.Context) {

	cntx, cancel := context.WithTimeout(ctx.Request.Context(), u.contextTimeout)
	defer cancel()

	userInfo := dto.LoginUserRequest{}

	err := ctx.ShouldBind(&userInfo)
	if err != nil {
		err = errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		u.logger.Error(ctx, "unable to bind user data", zap.Error(err))
		_ = ctx.Error(err)
		return
	}

	user, err := u.userModule.GetUser(cntx, userInfo)

	if err != nil {
		_ = ctx.Error(err)
		return
	}

	constants.SuccessResponse(ctx, http.StatusOK, user, nil)
}
