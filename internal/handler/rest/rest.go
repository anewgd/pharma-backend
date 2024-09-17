package rest

import "github.com/gin-gonic/gin"

type User interface {
	CreateUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
}

type Pharmacy interface {
	CreatePharmacy(ctx *gin.Context)
	CreatePharmacyBranch(ctx *gin.Context)
	CreatePharmacyBranchManager(ctx *gin.Context)
	CreatePharmacist(ctx *gin.Context)
	LoginPharmacist(ctx *gin.Context)
}

type Drug interface {
	CreateDrug(ctx *gin.Context)
}
