package api

import (
	"context"
	"fmt"

	"github.com/anewgd/pharma_backend/service"
	"github.com/anewgd/pharma_backend/util"
	"github.com/anewgd/pharma_backend/util/token"
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

	c := ctx.Request.Context()

	payload, ok := ctx.Get(authorizationHeaderKey)
	if !ok {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": "cannot find authorization payload",
		})
		return
	}
	usrPayload, ok := (payload).(*token.Payload)
	if !ok {
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": "can't get user id",
		})
		return
	}
	c = context.WithValue(c, util.UserID, usrPayload.UserID)

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
	c, err := getContext(ctx)
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

	c, err := getContext(ctx)
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

func getContext(ctx *gin.Context) (context.Context, error) {

	c := ctx.Request.Context()
	payload, ok := ctx.Get(authorizationHeaderKey)
	if !ok {
		return nil, fmt.Errorf("cannot find authorization payload")
	}

	usrPayload, ok := (payload).(*token.Payload)
	if !ok {
		return nil, fmt.Errorf("can't get user id")
	}

	c = context.WithValue(c, util.UserID, usrPayload.UserID)
	c = context.WithValue(c, util.Role, usrPayload.Role)

	return c, nil
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
