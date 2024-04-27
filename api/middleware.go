package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/anewgd/pharma_backend/util/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "no authorization header found",
			})
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "invalid authorization format",
			})
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": fmt.Errorf("no authorization type %s", authorizationType).Error(),
			})
			return

		}

		accessToken := fields[1]
		authPayload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": fmt.Errorf("invalid token: %s", err.Error()).Error(),
			})
			return
		}
		ctx.Set(authorizationHeaderKey, authPayload)
		ctx.Next()
	}
}
