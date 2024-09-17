package initiator

import (
	"pharma-backend/internal/module"
	"pharma-backend/internal/module/drug"
	"pharma-backend/internal/module/pharmacy"
	"pharma-backend/internal/module/user"
	"pharma-backend/platform/logger"
)

type Module struct {
	user     module.User
	pharmacy module.Pharmacy
	drug     module.Drug
}

func InitModule(persistence Persistence, platformLayer PlatformLayer, log logger.Logger) Module {
	return Module{
		user:     user.Init(log.Named("user-module"), persistence.user, platformLayer.Token, persistence.userSession),
		pharmacy: pharmacy.Init(log.Named("pharmacy-module"), persistence.pharmacy),
		drug:     drug.Init(log.Named("drug-module"), persistence.drug, persistence.pharmacy),
	}
}
