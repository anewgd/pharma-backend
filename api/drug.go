package api

import (
	"net/http"

	"github.com/anewgd/pharma_backend/service"
	"github.com/anewgd/pharma_backend/util"
	"github.com/gin-gonic/gin"
)

type DrugHandler struct {
	drugService service.DrugService
}

func NewDrugHandler(drugService service.DrugService) *DrugHandler {
	return &DrugHandler{
		drugService: drugService,
	}
}

func (d *DrugHandler) addDrugHandler(ctx *gin.Context) {

	c, err := getContextWithValues(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}
	req := service.CreateDrugRequest{}
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(util.NewErrorResponse(util.RequestError.New("malformed request body"), http.StatusBadRequest, "invalid request"))
		return
	}

	drug, err := d.drugService.AddDrug(c, req)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusCreated, drug)
}
