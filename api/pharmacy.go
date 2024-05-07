package api

import (
	"net/http"

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

	c, err := getContextWithValues(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	req := service.CreatePharmacyRequest{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(util.NewErrorResponse(util.RequestError.New("malformed request body"), http.StatusBadRequest, "invalid request"))
		return
	}

	res, err := p.pharmacyService.CreatePharmacy(c, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, res)

}

func (p *PharmacyHandler) createPharmacyBranch(ctx *gin.Context) {
	c, err := getContextWithValues(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	req := service.CreatePharmacyBranchRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(util.NewErrorResponse(util.RequestError.New("malformed request body"), http.StatusBadRequest, "invalid request"))
		return
	}
	res, err := p.pharmacyService.CreatePharmacyBranch(c, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, res)

}

func (p *PharmacyHandler) createManager(ctx *gin.Context) {

	c, err := getContextWithValues(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	req := service.CreatePharmacyManagerRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(util.NewErrorResponse(util.RequestError.New("malformed request body"), http.StatusBadRequest, "invalid request"))
		return
	}

	res, err := p.pharmacyService.CreateBranchManager(c, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, res)

}

func (p *PharmacyHandler) pharmacyLogin(ctx *gin.Context) {
	c := ctx.Request.Context()
	req := service.LoginUserRequest{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.Error(util.NewErrorResponse(util.RequestError.New("malformed request body"), http.StatusBadRequest, "invalid request"))
		return
	}

	res, err := p.pharmacyService.PharmacyLogin(c, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, res)
}
