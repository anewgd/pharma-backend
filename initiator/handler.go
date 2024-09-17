package initiator

import (
	"pharma-backend/internal/handler/rest"
	"pharma-backend/internal/handler/rest/gin/drug"
	"pharma-backend/internal/handler/rest/gin/pharmacy"
	"pharma-backend/internal/handler/rest/gin/user"
	"pharma-backend/platform/logger"
	"time"
)

type Handler struct {
	user     rest.User
	pharmacy rest.Pharmacy
	drug     rest.Drug
}

func InitHandler(module Module, log logger.Logger, timeout time.Duration) Handler {
	return Handler{
		user:     user.Init(log.Named("user-handler"), module.user, timeout),
		pharmacy: pharmacy.Init(log.Named("pharmacy-handler"), module.pharmacy, timeout),
		drug:     drug.Init(log.Named("drug-handler"), module.drug, timeout),
	}
}
