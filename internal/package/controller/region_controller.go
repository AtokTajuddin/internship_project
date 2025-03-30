package controller

import (
	"net/http"

	"project_virtual_internship_evermos/internal/package/usecase"

	"github.com/gin-gonic/gin"
)

type RegionController struct {
	RegionUsecase usecase.RegionUsecase
}

func NewRegionController(regionUsecase usecase.RegionUsecase) *RegionController {
	return &RegionController{
		RegionUsecase: regionUsecase,
	}
}

func (c *RegionController) GetProvinces(ctx *gin.Context) {
	provinces, err := c.RegionUsecase.GetProvinces()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": provinces})
}

func (c *RegionController) GetRegencies(ctx *gin.Context) {
	provinceID := ctx.Param("provinceId")
	regencies, err := c.RegionUsecase.GetRegencies(provinceID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": regencies})
}

func (c *RegionController) GetDistricts(ctx *gin.Context) {
	regencyID := ctx.Param("regencyId")
	districts, err := c.RegionUsecase.GetDistricts(regencyID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": districts})
}

func (c *RegionController) GetVillages(ctx *gin.Context) {
	districtID := ctx.Param("districtId")
	villages, err := c.RegionUsecase.GetVillages(districtID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": villages})
}
