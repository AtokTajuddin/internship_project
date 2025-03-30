package handler

import (
	"project_virtual_internship_evermos/internal/package/controller"

	"github.com/gin-gonic/gin"
)

func RegionRoutes(r *gin.Engine, regionController *controller.RegionController) {
	regionGroup := r.Group("/regions")
	{
		regionGroup.GET("/provinces", regionController.GetProvinces)
		regionGroup.GET("/provinces/:provinceId/regencies", regionController.GetRegencies)
		regionGroup.GET("/regencies/:regencyId/districts", regionController.GetDistricts)
		regionGroup.GET("/districts/:districtId/villages", regionController.GetVillages)
	}
}
