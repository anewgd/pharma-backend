package initiator

import (
	"pharma-backend/internal/constants/dbinstance"
	"pharma-backend/internal/storage"
	"pharma-backend/internal/storage/sqlc/drug"
	"pharma-backend/internal/storage/sqlc/pharmacy"
	"pharma-backend/internal/storage/sqlc/pharmacySession"
	"pharma-backend/internal/storage/sqlc/user"
	"pharma-backend/internal/storage/sqlc/userSession"
	"pharma-backend/platform/logger"
)

type Persistence struct {
	user            storage.User
	pharmacy        storage.Pharmacy
	drug            storage.Drug
	userSession     storage.UserSession
	pharamcySession storage.PharmacySession
}

func InitPersistence(db dbinstance.DBInstance, log logger.Logger) Persistence {
	return Persistence{
		user:            user.Init(db, log.Named("user-persistence")),
		pharmacy:        pharmacy.Init(db, log.Named("pharmacy-persistence")),
		drug:            drug.Init(db, log.Named("drug-persistence")),
		userSession:     userSession.Init(db, log.Named("userSession-persistence")),
		pharamcySession: pharmacySession.Init(db, log.Named("pharamcySession-persistence")),
	}
}
