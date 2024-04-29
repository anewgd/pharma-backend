package api

import (
	"github.com/anewgd/pharma_backend/service"
	"github.com/anewgd/pharma_backend/util"
	"github.com/gin-gonic/gin"
)

type PharmacyHandler struct {
	pharmacyService service.PharmacyService
}

func NewPharmacyHandler(pharmacyService service.PharmacyService) *PharmacyHandler {
	return &PharmacyHandler{
		pharmacyService: pharmacyService,
	}
}

func (p *PharmacyHandler) createPharmacy(ctx *gin.Context) {

	c, err := util.GetContextWithValues(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
	}

	req := service.CreatePharmacyRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := p.pharmacyService.CreatePharmacy(c, req)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, res)

}

func (p *PharmacyHandler) createPharmacyBranch(ctx *gin.Context) {
	c, err := util.GetContextWithValues(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	req := service.CreatePharmacyBranchRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := p.pharmacyService.CreatePharmacyBranch(c, req)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, res)

}

func (p *PharmacyHandler) createManager(ctx *gin.Context) {

	c, err := util.GetContextWithValues(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	req := service.CreatePharmacyManagerRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := p.pharmacyService.CreateBranchManager(c, req)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, res)

}

func (p *PharmacyHandler) pharmacyLogin(ctx *gin.Context) {
	c := ctx.Request.Context()
	req := service.LoginUserRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := p.pharmacyService.PharmacyLogin(c, req)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, res)
}
