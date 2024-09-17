package initiator

import (
	"pharma-backend/internal/glue/routing/drug"
	"pharma-backend/internal/glue/routing/pharmacy"
	"pharma-backend/internal/glue/routing/user"
	"pharma-backend/internal/handler/middleware"
	"pharma-backend/platform/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/swaggo/swag/example/basic/docs"
)

func InitRouter(group *gin.RouterGroup, handler Handler, logger logger.Logger, platformLayer PlatformLayer) {

	authMiddleware := middleware.InitAuthMiddleware(logger.Named("auth-middleware"), platformLayer.Token)
	docs.SwaggerInfo.BasePath = "/v1"
	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	user.InitRoute(group, handler.user)
	pharmacy.InitRoute(group, handler.pharmacy, authMiddleware)
	drug.InitRoute(group, handler.drug)
}
