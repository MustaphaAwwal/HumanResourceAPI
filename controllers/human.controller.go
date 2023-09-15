package controllers

import (
	"hng-stage2/resource"
	"hng-stage2/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type HumanController struct {
	humanService services.HumanService
}

func NewHumanController(humanService services.HumanService) HumanController{
	return HumanController{humanService}
}

func (hc *HumanController) CreateHuman(ctx *gin.Context){

	var human *resource.Human

	if err := ctx.ShouldBindJSON(&human); err != nil {
		ctx.JSON(http.StatusBadRequest , err.Error())
		return
	}


	newHuman, err := hc.humanService.CreateHuman(human)
	if err != nil{
		if strings.Contains(err.Error(), "name already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newHuman})
}

func (h *HumanController) UpdateHuman(ctx *gin.Context){
	humanId := ctx.Param("user_id")

	var human *resource.Human
	if err := ctx.ShouldBindJSON(&human); err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"status":"fail", "message": err})
		return
	}

	updated, err := h.humanService.UpdateHuman(humanId, human)

	if err != nil {
		if strings.Contains(err.Error(), "Id exists"){
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
			return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updated})
}

func (h *HumanController) DeleteHuman(ctx *gin.Context){
	humanID := ctx.Param("user_id")

	err := h.humanService.DeleteHuman(humanID)

	if err != nil {
		if strings.Contains(err.Error(), "Id exist") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (h *HumanController) GetHumanbyID(ctx *gin.Context){
	humanID := ctx.Param("user_id")

	human, err := h.humanService.GetHumanbyID(humanID)
	if err != nil{
		if strings.Contains(err.Error(), "Id exists") {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
			return
	}
	ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": human})
}

func (h *HumanController) GetAllHuman(ctx *gin.Context)  {
	humans, err := h.humanService.GetAllHuman()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": humans})
}
