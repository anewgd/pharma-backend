package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	db "github.com/anewgd/pharma_backend/data/sqlc"
	"github.com/anewgd/pharma_backend/service"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
)

// type testStore struct {
// 	store map[string]db.Drug
// }

// func newTestStore(store map[string]db.Drug) *testStore {
// 	return &testStore{
// 		store: store,
// 	}
// }

// func (ts *testStore) CreateDrug(ctx context.Context, arg db.CreateDrugParams) (db.Drug, error) {
// 	newDrug := db.Drug{}

// 	newDrug.DrugID = uuid.Max
// 	newDrug.PharmacyID = arg.PharmacyID
// 	newDrug.BrandName = arg.BrandName
// 	newDrug.GenericName = arg.GenericName
// 	newDrug.Quantity = arg.Quantity
// 	newDrug.ExpirationDate = arg.ExpirationDate
// 	newDrug.ManufacturingDate = arg.ManufacturingDate
// 	newDrug.UserID = arg.UserID
// 	newDrug.AddedAt = time.Now()

// 	ts.store["testDrug"] = newDrug
// 	return newDrug, nil
// }

func TestAddDrug(t *testing.T) {

	connPool, err := pgxpool.New(context.Background(), "postgresql://pharma_dev_user:pharmadevpass@localhost:5432/pharma-dev?sslmode=disable")
	assert.NoError(t, err)
	testStore := db.NewStore(connPool)
	drugService := service.NewDrugService(testStore)
	drugHandler := NewDrugHandler(drugService)
	testServer := NewServer(drugHandler, nil, nil)

	pID, err := uuid.NewV7()
	assert.NoError(t, err)
	uID, err := uuid.NewV7()
	assert.NoError(t, err)

	args := db.CreateDrugParams{

		PharmacyBranchID:  pID,
		BrandName:         "test brand",
		GenericName:       "test Generic name",
		Quantity:          1200,
		ExpirationDate:    time.Now().AddDate(2, -1, 0),
		ManufacturingDate: time.Now().AddDate(0, -1, 0),
		PharmacistID:      uID,
	}
	reqBody, err := json.Marshal(&args)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/drugs", strings.NewReader(string(reqBody)))
	assert.NoError(t, err)
	res := httptest.NewRecorder()

	testServer.router.ServeHTTP(res, req)

	d := db.Drug{}
	err = json.Unmarshal(res.Body.Bytes(), &d)
	assert.NoError(t, err)

	assert.Equal(t, d.DrugID, uuid.Max)
	assert.Equal(t, d.PharmacyBranchID, pID)
	assert.Equal(t, d.BrandName, args.BrandName)
	assert.Equal(t, d.GenericName, args.GenericName)
	assert.Equal(t, d.Quantity, args.Quantity)
	assert.Equal(t, d.ExpirationDate, args.ExpirationDate)
	assert.Equal(t, d.ManufacturingDate, args.ManufacturingDate)
	assert.InEpsilon(t, d.AddedAt.Nanosecond(), time.Now().Nanosecond(), 0.01)
}
