package api

import (
	"net/http"
	"strings"

	"github.com/anewgd/pharma_backend/token"
	"github.com/anewgd/pharma_backend/util"
	"github.com/gin-gonic/gin"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(util.AuthorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			ctx.Error(util.NewErrorResponse(util.AuthenticationError.New("no authorization header found in request"), http.StatusBadRequest, "authentication error"))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			ctx.Error(util.NewErrorResponse(util.RequestError.New("invalid authorization format"), http.StatusBadRequest, "authentication error"))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != util.AuthorizationTypeBearer {
			ctx.Error(util.NewErrorResponse(util.RequestError.New("no authorization type %s", authorizationType), http.StatusBadRequest, "authentication error"))
			return

		}

		accessToken := fields[1]
		authPayload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.Error(util.NewErrorResponse(util.AuthenticationError.Wrap(err, "invalid token"), http.StatusBadRequest, "invalid request"))
			return
		}
		ctx.Set(util.AuthorizationPayloadKey, authPayload)
		ctx.Next()
	}
}

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		err := ctx.Errors.Last()
		if err != nil {
			//Extract the error from gin.Error and cast it as util.ErrorResponse
			e, ok := err.Err.(util.ErrorResponse)
			if !ok {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": e.Error()})
				return
			}
			ctx.AbortWithStatusJSON(e.StatusCode, gin.H{"error": e.Error()})
		}
	}
}
