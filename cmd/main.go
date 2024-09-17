package main

import (
	"context"

	"pharma-backend/initiator"
)

func main() {

	initiator.Initiator(context.Background())
	// connPool, err := pgxpool.New(context.Background(), "postgresql://pharma_dev_user:pharmadevpass@localhost:5432/pharma-dev?sslmode=disable")

	// if err != nil {
	// 	log.Fatalf("cannot connect to database: %s", err.Error())
	// }
	// store := db.NewStore(connPool)
	// drugService := service.NewDrugService(store)
	// userService, err := service.NewUserService(store)
	// if err != nil {
	// 	log.Fatalf("cannot create service: %s", err.Error())
	// }
	// pharmacyService, err := service.NewPharmacyService(store)
	// if err != nil {
	// 	log.Fatalf("cannot create pharmacy service: %s", err.Error())
	// }

	// drugHandler := api.NewDrugHandler(drugService)
	// userHandler := api.NewUserHandler(userService)
	// pharmacyHandler := api.NewPharmacyHandler(pharmacyService)

	// server := api.NewServer(drugHandler, userHandler, pharmacyHandler)

	// server.Start(":8083")

}
