package routes

import (
	"hng-stage2/controllers"

	"github.com/gin-gonic/gin"
)

type HumanRouteController struct {
	humanController controllers.HumanController
}

func NewHumanControllerRoute(humanController controllers.HumanController)  HumanRouteController{
	return HumanRouteController{humanController}
	
}

func (r *HumanRouteController) HumanRoute(rg *gin.RouterGroup) {
	router := rg.Group("/api")

	router.GET("/", r.humanController.GetAllHuman)
	router.GET("/:user_id", r.humanController.GetHumanbyID)
	router.POST("/", r.humanController.CreateHuman)
	router.PATCH("/:user_id", r.humanController.UpdateHuman)
	router.DELETE("/:user_id", r.humanController.DeleteHuman)
}
