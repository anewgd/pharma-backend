package middleware

import (
	"context"
	"net/http"
	"pharma-backend/internal/constants"
	"pharma-backend/internal/constants/errors"
	"pharma-backend/platform"
	"pharma-backend/platform/logger"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

type AuthenticationMiddleware interface {
	Authentication() gin.HandlerFunc
}

type authMiddleware struct {
	token  platform.Token
	logger logger.Logger
}

func InitAuthMiddleware(logger logger.Logger, token platform.Token) AuthenticationMiddleware {
	return &authMiddleware{
		token:  token,
		logger: logger,
	}
}

func (a *authMiddleware) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearer := "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			err := errors.ErrInvalidToken.New("Unauthorized")
			a.logger.Error(ctx, "no token found in request", zap.Error(err))
			ctx.Error(err)
			ctx.Abort()
			return
		}

		tokenString := authHeader[len(bearer):]
		valid, claims, err := a.token.VerifyToken(jwt.SigningMethodPS512, tokenString)
		if !valid {
			err := errors.ErrInvalidToken.Wrap(err, "Unauthorized")
			a.logger.Error(ctx, "invalid token", zap.Error(err))
			ctx.Error(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// fmt.Println(claims.UserID)

		ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), constants.Context("x-user-role"), claims.Role))
		ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), constants.Context("x-user-id"), claims.Subject))
		ctx.Next()
	}
}
