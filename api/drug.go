package api

import "github.com/gin-gonic/gin"

type DrugHandler struct{}

func NewDrugHandler() *DrugHandler {
	return &DrugHandler{}
}

func (d *DrugHandler) addDrug(ctx *gin.Context) {
	ctx.String(200, "test drug")
}
