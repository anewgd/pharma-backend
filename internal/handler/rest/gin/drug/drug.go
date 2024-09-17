package drug

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

type drug struct {
	logger         logger.Logger
	drugModule     module.Drug
	contextTimeout time.Duration
}

func Init(log logger.Logger, drugModule module.Drug, contextTimeout time.Duration) rest.Drug {

	return &drug{
		logger:         log,
		drugModule:     drugModule,
		contextTimeout: contextTimeout,
	}
}

func (d *drug) CreateDrug(ctx *gin.Context) {
	cntx, cancel := context.WithTimeout(ctx.Request.Context(), d.contextTimeout)
	defer cancel()

	drugInfo := dto.CreateDrugReq{}

	err := ctx.ShouldBind(&drugInfo)
	if err != nil {
		err = errors.ErrInvalidUserInput.Wrap(err, "invalid input")
		d.logger.Error(ctx, "unable to bind drug data", zap.Error(err))
		_ = ctx.Error(err)
		return
	}

	newDrug, err := d.drugModule.CreateDrug(cntx, drugInfo)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	constants.SuccessResponse(ctx, http.StatusCreated, newDrug, nil)
}
