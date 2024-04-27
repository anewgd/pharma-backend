package api

import (
	"net/http"

	"github.com/anewgd/pharma_backend/service"
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

	req := service.CreateDrugRequest{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "bad request")
		return
	}

	//You are sending the gin context here. it should be the context of the request
	drug, err := d.drugService.AddDrug(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(200, drug)
}
