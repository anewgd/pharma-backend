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

	c, err := util.GetContextWithValues(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	req := service.CreateDrugRequest{}
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	drug, err := d.drugService.AddDrug(c, req)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, drug)
}
