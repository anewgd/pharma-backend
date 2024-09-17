package pharmacy

import (
	"context"
	"net/http"
	"pharma-backend/internal/constants"
	"pharma-backend/internal/constants/errors"
	"pharma-backend/internal/constants/model/dto"
	"pharma-backend/internal/handler/rest"
	"pharma-backend/internal/module"
	"pharma-backend/platform/logger"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type pharmacy struct {
	logger         logger.Logger
	pharmacyModule module.Pharmacy
	contextTimeout time.Duration
}

func Init(log logger.Logger, pharmacyModule module.Pharmacy, contextTimeout time.Duration) rest.Pharmacy {

	return &pharmacy{
		logger:         log,
		pharmacyModule: pharmacyModule,
		contextTimeout: contextTimeout,
	}
}

func (p *pharmacy) CreatePharmacy(ctx *gin.Context) {

	cntx, cancel := context.WithTimeout(ctx.Request.Context(), p.contextTimeout)
	defer cancel()

	newPharmacyInfo := dto.CreatePharmacyReq{}

	err := ctx.ShouldBind(&newPharmacyInfo)
	if err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		p.logger.Error(ctx, "unable to bind pharmacy data", zap.Error(err))
		_ = ctx.Error(err)
		return
	}

	newPharmacy, err := p.pharmacyModule.CreatePharmacy(cntx, newPharmacyInfo.PharmacyName)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	constants.SuccessResponse(ctx, http.StatusCreated, newPharmacy, nil)
}
func (p *pharmacy) CreatePharmacyBranch(ctx *gin.Context) {
	cntx, cancel := context.WithTimeout(ctx.Request.Context(), p.contextTimeout)
	defer cancel()

	newBranchInfo := dto.CreatePharmacyBranchReq{}

	err := ctx.ShouldBind(&newBranchInfo)
	if err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		p.logger.Error(ctx, "unable to bind pharmacy data", zap.Error(err))
		_ = ctx.Error(err)
		return
	}

	newBranch, err := p.pharmacyModule.CreatePharmacyBranch(cntx, newBranchInfo)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	constants.SuccessResponse(ctx, http.StatusCreated, newBranch, nil)
}
func (p *pharmacy) CreatePharmacyBranchManager(ctx *gin.Context) {
	cntx, cancel := context.WithTimeout(ctx.Request.Context(), p.contextTimeout)
	defer cancel()

	managerInfo := dto.CreatePharmacistReq{}

	err := ctx.ShouldBind(&managerInfo)
	if err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		p.logger.Error(ctx, "unable to bind pharmacist data", zap.Error(err))
		_ = ctx.Error(err)
		return
	}

	newManager, err := p.pharmacyModule.CreatePharmacyBranchManager(cntx, managerInfo)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	constants.SuccessResponse(ctx, http.StatusCreated, newManager, nil)

}
func (p *pharmacy) CreatePharmacist(ctx *gin.Context) {
	cntx, cancel := context.WithTimeout(ctx.Request.Context(), p.contextTimeout)
	defer cancel()

	managerInfo := dto.CreatePharmacistReq{}

	err := ctx.ShouldBind(&managerInfo)
	if err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		p.logger.Error(ctx, "unable to bind pharmacist data", zap.Error(err))
		_ = ctx.Error(err)
		return
	}

	newManager, err := p.pharmacyModule.CreatePharmacist(cntx, managerInfo)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	constants.SuccessResponse(ctx, http.StatusCreated, newManager, nil)
}
func (p *pharmacy) LoginPharmacist(ctx *gin.Context) {
	cntx, cancel := context.WithTimeout(ctx.Request.Context(), p.contextTimeout)
	defer cancel()

	userInfo := dto.LoginUserRequest{}

	err := ctx.ShouldBind(&userInfo)
	if err != nil {
		err := errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		p.logger.Error(ctx, "unable to bind pharmacy login data", zap.Error(err))
		_ = ctx.Error(err)
		return
	}

	pharmacist, err := p.pharmacyModule.GetPharmacist(cntx, userInfo)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	constants.SuccessResponse(ctx, http.StatusOK, pharmacist, nil)

}
