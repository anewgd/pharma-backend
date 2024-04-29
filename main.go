package main

import (
	"context"
	"log"

	"github.com/anewgd/pharma_backend/api"
	db "github.com/anewgd/pharma_backend/data/sqlc"
	"github.com/anewgd/pharma_backend/service"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	connPool, err := pgxpool.New(context.Background(), "postgresql://pharma_dev_user:pharmadevpass@localhost:5432/pharma-dev?sslmode=disable")

	if err != nil {
		log.Fatalf("cannot connect to database: %s", err.Error())
	}
	store := db.NewStore(connPool)
	drugService := service.NewDrugService(store)
	userService, err := service.NewUserService(store)
	if err != nil {
		log.Fatalf("cannot create service: %s", err.Error())
	}
	pharmacyService, err := service.NewPharmacyService(store)
	if err != nil {
		log.Fatalf("cannot create pharmacy service: %s", err.Error())
	}

	drugHandler := api.NewDrugHandler(drugService)
	userHandler := api.NewUserHandler(userService)
	pharmacyHandler := api.NewPharmacyHandler(pharmacyService)

	server := api.NewServer(drugHandler, userHandler, pharmacyHandler)

	server.Start(":8083")

}
